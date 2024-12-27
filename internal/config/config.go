package config

import (
	"github.com/spf13/viper"
)

// Config is the configuration for the application
type Config struct {
	PhoneNumber string `mapstructure:"phone_number"`
	Password    string `mapstructure:"password"`
	GroupID     string `mapstructure:"group_id"`
	RootPath    string `mapstructure:"root_path"`
}

// LoadConfig loads the configuration from a file
func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var cfg Config
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
