package options

import (
	"github.com/spf13/pflag"
	"ooneko.github.com/vehicle-insight/pkg/storage/storagebackend/config"
)

type EtcdOptions struct {
	config.Config
}

func NewEtcdOptions() *EtcdOptions {
	return &EtcdOptions{
		Config: config.NewDefaultConfig(),
	}
}

func (o *EtcdOptions) AddFlags(fs *pflag.FlagSet) {
	fs.StringSliceVar(&o.Transport.ServerList, "etcd-servers", []string{},
		"List of etcd servers to connect with (scheme://ip:port), comma separated.")

	fs.StringVar(&o.Transport.CertFile, "etcd-certfile", o.Transport.CertFile,
		"SSL certification file used to secure etcd communication")

	fs.StringVar(&o.Transport.KeyFile, "etcd-keyfile", o.Transport.KeyFile,
		"SSL key file used to secure etcd communication")

	fs.StringVar(&o.Transport.TrustedCAFile, "etcd-cafile", o.Transport.TrustedCAFile,
		"SSL Certificate Authority file used to secure etcd communication")

	fs.StringVar(&o.Prefix, "etcd-prefix", config.DefaultPrefix,
		"The prefix to prepend to all resource paths in etcd")
}
