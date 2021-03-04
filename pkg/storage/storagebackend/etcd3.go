package storagebackend

import (
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"

	"ooneko.github.com/vehicle-insight/pkg/storage"
	"ooneko.github.com/vehicle-insight/pkg/storage/etcd3"
	"ooneko.github.com/vehicle-insight/pkg/storage/storagebackend/config"
)

const dialTimeout = 20 * time.Second

func NewETCD3Storage(c config.Config, logger *zap.Logger) (storage.Interface, error) {
	client, err := newETCD3Client(c.Transport)
	if err != nil {
		return nil, err
	}
	return etcd3.New(client, c.Prefix, logger), nil
}

func newETCD3Client(c config.TransportConfig) (*clientv3.Client, error) {
	//TODO Add tls config
	cfg := clientv3.Config{
		Endpoints:   c.ServerList,
		DialTimeout: dialTimeout,
	}
	return clientv3.New(cfg)
}
