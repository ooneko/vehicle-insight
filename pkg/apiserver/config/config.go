package config

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	DefaultCfgName = "car-pacemaker"
	DefaultCfgPath = "/etc/car-pacemaker"
)

type Config struct {
}

func new() *Config {
	return &Config{}
}

func Load() *Config {
	config := new()
	viper.SetConfigName(DefaultCfgName)
	viper.AddConfigPath(DefaultCfgPath)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// log.Print("configration file not found")
			return config
		}
		msg := fmt.Sprintf("error parsing configuration %v", err)
		panic(msg)
	}

	if err := viper.Unmarshal(config); err != nil {
		msg := fmt.Sprintf("error unmarshal configration %v", err)
		panic(msg)
	}
	return config
}
