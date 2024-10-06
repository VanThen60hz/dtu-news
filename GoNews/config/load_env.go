package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	// Mysql setup
	DbHost     string `mapstructure:"SQL_HOST"`
	DbUsername string `mapstructure:"SQL_USER"`
	DbPassword string `mapstructure:"SQL_PASSWORD"`
	DbName     string `mapstructure:"SQL_DB"`
	DbPort     string `mapstructure:"SQL_PORT"`

	// Redis setup
	RedisUrl string `mapstructure:"REDIS_URL"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("app")

	viper.AutomaticEnv()

	// handle null
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	return
}
