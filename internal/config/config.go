package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Cache       RedisConfig
	ServiceName string              `mapstructure:"service_name"`
	IsVerbose   bool                `mapstructure:"is_verbose"`
	HttpServer  HttpServerConfig    `mapstructure:"http_server"`
	Queue       MessageStreamConfig `mapstructure:"queue"`
	Bonita      BonitaConfig        `mapstructure:"bonita"`
	Database    DatabaseConfig      `mapstructure:"database"`
	HttpClient  HttpClientConfig    `mapstructure:"http_client"`
}

type HttpClientConfig struct {
	Host                string  `mapstructure:"host"`
	Port                string  `mapstructure:"port"`
	Timeout             int     `mapstructure:"timeout"`
	ConsecutiveFailures int     `mapstructure:"consecutive_failures"`
	ErrorRate           float64 `mapstructure:"error_rate"` // 0.001 - 0.999
	Threshold           int     `mapstructure:"threshold"`
	Type                string  `mapstructure:"type"` // consecutive, error_rate
}
type DatabaseConfig struct {
	Host         string `mapstructure:"host"`
	Port         int    `mapstructure:"port"`
	Username     string `mapstructure:"username"`
	Password     string `mapstructure:"password"`
	DBName       string `mapstructure:"db_name"`
	SSL          string `mapstructure:"ssl"`
	SchemaName   string `mapstructure:"schema_name"`
	MaxIdleConns int    `mapstructure:"max_idle_conns"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
	Timeout      int    `mapstructure:"timeout"`
}

type HttpServerConfig struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type BonitaConfig struct {
	Host               string `mapstructure:"host"`
	Port               string `mapstructure:"port"`
	Username           string `mapstructure:"username"`
	Password           string `mapstructure:"password"`
	Timeout            int    `mapstructure:"timeout"`
	LoginCacheDuration int    `mapstructure:"login_cache_duration"`
}

type MessageStreamConfig struct {
	Host                string `mapstructure:"host"`
	Port                string `mapstructure:"port"`
	Username            string `mapstructure:"username"`
	Password            string `mapstructure:"password"`
	ExchangeName        string `mapstructure:"exchange_name"`
	PublishTopic        string `mapstructure:"publish_topic"`
	DeadLetterNameQueue string `mapstructure:"dead_letter_name_queue"`
	SubscribeTopic      string `mapstructure:"subscribe_topic"`
}

type RedisConfig struct {
	Host            string        `mapstructure:"host"`
	Port            string        `mapstructure:"port"`
	Username        string        `mapstructure:"username"`
	Password        string        `mapstructure:"password"`
	DB              int           `mapstructure:"db"`
	MaxRetries      int           `mapstructure:"max_retries"`
	PoolFIFO        bool          `mapstructure:"pool_fifo"`
	PoolSize        int           `mapstructure:"pool_size"`
	PoolTimeout     time.Duration `mapstructure:"pool_timeout"`
	MinIdleConns    int           `mapstructure:"min_idle_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

func InitConfig() *Config {
	// // Viper add remote provider
	// viper.AddRemoteProvider("consul", "localhost:8500", "config/hello-service.mapstructure")
	// viper.SetConfigType("mapstructure")
	// err := viper.ReadRemoteConfig()
	// if err != nil {
	// 	panic(err)
	// }

	// Viper read file from path
	viper.SetConfigName("config")
	viper.AddConfigPath("./internal/config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
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
