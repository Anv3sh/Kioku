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
	// CacheSize  int64	`mapstructure:"cache_size"`
	MaxMem float64 `mapstructure:"maxmemory"`
	LFUEviction     bool `mapstructure:"lfu-eviction"`
	LRUEviction 	bool `mapstructure:"lru-eviction"`
	AOFPolicy       string   `mapstructure:"appendfsync"`
	AOFPath         string  `mapstructure:"aoffilepath"`
}

func (config *Config) CreateConfig(){
	config.ServerHost="localhost"
	config.ServerPort="6379"
	config.MaxClients=100
	config.TotalTimetoLive=300
	config.MaxMem=0.00005
	config.LFUEviction=true
	config.LRUEviction=false
	config.AOFPolicy="NO"
	config.AOFPath="./appendonly.aof"
}

func (config *Config) SetConfig() {
	viper.SetConfigFile("./kioku.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config: ", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error unmarshaling config: ", err)
	}
	log.Println("Kioku configurations set.")
}
