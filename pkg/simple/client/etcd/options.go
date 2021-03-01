package etcd

import "github.com/spf13/pflag"

type Options struct {
	Endpoint []string `json:"endpoint,omitempty" yaml:"endpoint,omitempty" mapstructure:"endpoint"`
	Username string   `json:"username,omitempty" yaml:"username,omitempty" mapstructure:"username"`
	Password string   `json:"password,omitempty" yaml:"password,omitempty" mapstructure:"password"`
}

func New() *Options {
	return &Options{}
}

func (o *Options) AddFlags(fs *pflag.FlagSet, s *Options) {
	fs.StringSliceVar(&o.Endpoint, "etcd-endpoints", s.Endpoint, "etcd endpoints. e.g: --ectd-endpoint=\"localhost:2379,localhost12379\"")
	fs.StringVar(&o.Username, "etcd-username", s.Username, "etcd username.")
	fs.StringVar(&o.Password, "etcd-password", s.Password, "etcd password")
}
