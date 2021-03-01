package options

import (
	"time"

	"ooneko.github.com/vehicle-insight/pkg/apiserver"
	"ooneko.github.com/vehicle-insight/pkg/apiserver/config"

	"go.etcd.io/etcd/clientv3"
	cliflag "k8s.io/component-base/cli/flag"
)

type ServerRunOptions struct {
	*config.Config
}

func NewServerRunOptions() *ServerRunOptions {
	return &ServerRunOptions{
		Config: config.Load(),
	}
}

func (s *ServerRunOptions) Flags() (fss cliflag.NamedFlagSets) {
	s.EtcdOptions.AddFlags(fss.FlagSet("etcd"), s.EtcdOptions)
	return
}

func (s *ServerRunOptions) NewAPIServer() *apiserver.APIServer {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   s.EtcdOptions.Endpoint,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	return &apiserver.APIServer{
		EtcdClinet: client,
	}
}
