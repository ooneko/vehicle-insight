package options

import (
	"github.com/spf13/pflag"
	"go.uber.org/zap"

	"ooneko.github.com/vehicle-insight/pkg/log"
)

func NewAPIServerFlags() *APIServerFlags {
	return &APIServerFlags{
		Etcd:     NewEtcdOptions(),
		LogLevel: log.DefaultLogLevel,
	}
}

type APIServerFlags struct {
	Etcd     *EtcdOptions
	LogLevel string
}

func (f *APIServerFlags) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&f.LogLevel, "log-level", "info", "log level")
	f.Etcd.AddFlags(fs)
}

type APIServer struct {
	APIServerFlags

	Logger *zap.Logger
}

func ValidateAPIServerFlags(f *APIServerFlags) error {
	return nil
}
