package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server Server `mapstructure:",squash"`
	Redis  Redis  `mapstructure:",squash"`
}

type Server struct {
	Address string `mapstructure:"server_address"`
}

type Redis struct {
	Address  string `mapstructure:"redis_address"`
	Password string `mapstructure:"redis_password"`
	Database int    `mapstructure:"redis_database"`
}

func New() (*Config, error) {
	viper.SetDefault("SERVER_ADDRESS", ":8080")
	viper.SetDefault("REDIS_ADDRESS", "redis:6379")
	viper.SetDefault("REDIS_PASSWORD", "")
	viper.SetDefault("REDIS_DATABASE", 0)

	viper.SetConfigType("env")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("config file not found")
	}

	cfg := &Config{}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
