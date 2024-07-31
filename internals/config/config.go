package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	ServerHost string `mapstructure:"host"`
	ServerPort string `mapstructure:"port"`
	MaxClients int    `mapstructure:"maxclients"`
}

func SetConfig(config *Config) {
	viper.SetConfigFile("./kioku.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config: ", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error unmarshaling config: ", err)
	}
}
