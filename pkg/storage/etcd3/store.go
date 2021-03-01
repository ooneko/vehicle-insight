package etcd3

import (
	"encoding/json"
	"errors"
	"fmt"
	"path"
	"strconv"

	"go.etcd.io/etcd/clientv3"
	"golang.org/x/net/context"

	metav1 "ooneko.github.com/vehicle-insight/pkg/apis/meta/v1"
	"ooneko.github.com/vehicle-insight/pkg/storage"
)

type store struct {
	client     *clientv3.Client
	getOps     []clientv3.OpOption
	versioner  *APIObjectVersioner
	pathPrefix string
}

func New(c *clientv3.Client, prefix string) storage.Interface {
	return newStore(c, prefix)
}

func newStore(c *clientv3.Client, prefix string) *store {
	return &store{
		client:     c,
		pathPrefix: path.Join("/", prefix),
	}
}

func (s *store) Create(ctx context.Context, key string, obj metav1.Object, ttl uint64) error {
	if version, err := s.versioner.ObjectResourceVersion(obj); err == nil && version != 0 {
		return errors.New("resourceVersion should not be set on objects to be created")
	}
	if err := s.versioner.PrepareObjectForStorage(obj); err != nil {
		return fmt.Errorf("PrepareObjectForStorage: %v", err)
	}

	data, err := json.Marshal(obj)
	if err != nil {
		return err
	}
	key = path.Join(s.pathPrefix, key)
	txnResp, err := s.client.KV.Txn(ctx).
		If(notFound(key)).Then(
		clientv3.OpPut(key, string(data))).Commit()
	if err != nil {
		return err
	}
	if !txnResp.Succeeded {
		return storage.NewKeyExistError(key, 0)
	}
	return nil
}

func (s *store) Get(ctx context.Context, key string, resourceVersion string, out metav1.Object, ignoreNotFound bool) error {
	key = path.Join(s.pathPrefix, key)
	getResp, err := s.client.KV.Get(ctx, key, s.getOps...)
	if err != nil {
		return err
	}
	if err = s.ensureMinimumResourceVersion(resourceVersion, uint64(getResp.Header.Revision)); err != nil {
		return err
	}
	if len(getResp.Kvs) == 0 {
		if ignoreNotFound {
			return nil
		}
		return storage.NewKeyNotFoundError(key, 0)
	}

	kv := getResp.Kvs[0]
	return json.Unmarshal(kv.Value, out)
}

func (s *store) Delete(ctx context.Context, key string, out interface{}) error {

}

func (s *store) List(ctx context.Context, key string, resourceVersion string, listObj interface{}) error {
}

func (s *store) Update(ctx context.Context, key string, obj, out interface{}) error {}

func (s *store) ensureMinimumResourceVersion(minimumResourceVersion string, actualRevision uint64) error {
	if minimumResourceVersion == "" || minimumResourceVersion == "0" {
		return nil
	}
	minimumRV, err := strconv.ParseUint(minimumResourceVersion, 10, 64)
	if err != nil {
		return fmt.Errorf("parseUint %s: %v", minimumResourceVersion, err)
	}
	if minimumRV > actualRevision {
		return fmt.Errorf("Too Large ResourceVersion Error")
	}
	return nil
}

func notFound(key string) clientv3.Cmp {
	return clientv3.Compare(clientv3.ModRevision(key), "=", 0)
}
