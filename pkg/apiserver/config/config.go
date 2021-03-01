package config

import (
	"log"

	"github.com/spf13/viper"
	"ooneko.github.com/vehicle-insight/pkg/simple/client/etcd"
)

const (
	DefaultCfgName = "car-pacemaker"
	DefaultCfgPath = "/etc/car-pacemaker"
)

type Config struct {
	EtcdOptions *etcd.Options
}

func new() *Config {
	return &Config{
		EtcdOptions: etcd.New(),
	}
}

func Load() *Config {
	config := new()
	viper.SetConfigName(DefaultCfgName)
	viper.AddConfigPath(DefaultCfgPath)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Print("configration file not found")
			return config
		}
		log.Panic("error parsing configuration", err)
	}

	if err := viper.Unmarshal(config); err != nil {
		log.Panic("error unmarshal configration", err)
	}
	return config
}
