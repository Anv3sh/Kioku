package config

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerHost      string        `mapstructure:"host"`
	ServerPort      string        `mapstructure:"port"`
	MaxClients      int           `mapstructure:"maxclients"`
	TotalTimetoLive time.Duration `mapstructure:"ttl"`
	Eviction        bool          `mapstructure:"eviction"`
	// CacheSize  int64	`mapstructure:"cache_size"`
	MaxMem float64 `mapstructure:"maxmemory"`
}

func SetConfig(config *Config) {
	viper.SetConfigFile("./kioku.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config: ", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error unmarshaling config: ", err)
	}
	log.Println("Kioku configurations set.")
}
