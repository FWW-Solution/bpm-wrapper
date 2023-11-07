package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Cache       *RedisConfig
	ServiceName string `json:"service_name"`
	IsVerbose   bool   `json:"is_verbose"`
	Queue       *MessageQueueConfig
	Bonita      *BonitaConfig
}

type BonitaConfig struct {
	Host               string `json:"host"`
	Port               string `json:"port"`
	Username           string `json:"username"`
	Password           string `json:"password"`
	Timeout            int    `json:"timeout"`
	LoginCacheDuration int    `json:"login_cache_duration"`
}

type MessageQueueConfig struct {
	Host                string `json:"host"`
	Port                string `json:"port"`
	Username            string `json:"username"`
	Password            string `json:"password"`
	ExchangeName        string `json:"exchange_name"`
	PublishTopic        string `json:"publish_topic"`
	DeadLetterNameQueue string `json:"dead_letter_name_queue"`
	SubscribeTopic      string `json:"subscribe_topic"`
}

type RedisConfig struct {
	Host            string        `json:"host"`
	Port            string        `json:"port"`
	Username        string        `json:"username"`
	Password        string        `json:"password"`
	DB              int           `json:"db"`
	MaxRetries      int           `json:"max_retries"`
	PoolFIFO        bool          `json:"pool_fifo"`
	PoolSize        int           `json:"pool_size"`
	PoolTimeout     time.Duration `json:"pool_timeout"`
	MinIdleConns    int           `json:"min_idle_conns"`
	MaxIdleConns    int           `json:"max_idle_conns"`
	ConnMaxIdleTime time.Duration `json:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`
}

func InitConfig() *Config {
	// Viper add remote provider
	viper.AddRemoteProvider("consul", "localhost:8500", "config/hello-service.json")
	viper.SetConfigType("json")
	err := viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}

	// Viper read file from path
	viper.SetConfigName("hello-service")
	viper.AddConfigPath("./internal/config")
	viper.SetConfigType("json")
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var Cfg Config

	// Viper Populates the struct
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		panic(err)
	}
	return &Cfg
}
