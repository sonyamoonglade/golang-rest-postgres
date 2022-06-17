package config

import (
	"github.com/spf13/viper"
)

type DbConfig struct {
	Host         string
	Port         string
	User         string
	Driver       string
	DatabaseName string
	Password     string
}

func InitConfig() error {

	viper.AddConfigPath("internal/config")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
