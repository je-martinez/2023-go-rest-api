package config

import (
	"errors"
	"log"
	"os"
	"time"

	constants "github.com/je-martinez/2023-go-rest-api/pkg/constants"

	"github.com/spf13/viper"
)

// Config Variables
var AppConfig *Config

// App config struct
type Config struct {
	Server   ServerConfig
	Database DatabaseConfig
	Redis    RedisConfig
	Metrics  Metrics
	Logger   Logger
	AWS      AWS
}

// Server config struct
type ServerConfig struct {
	AppVersion        string
	Address           string
	Port              string
	PprofPort         string
	Mode              string
	JwtSecretKey      string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Postgresql config
type DatabaseConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  bool
	Insecure           bool
}

// Redis config
type RedisConfig struct {
	RedisAddr      string
	RedisPassword  string
	RedisDB        string
	RedisDefaultdb string
	MinIdleConns   int
	PoolSize       int
	PoolTimeout    int
	Password       string
	DB             int
}

// Metrics config
type Metrics struct {
	URL         string
	ServiceName string
}

// AWS S3
type AWS struct {
	Endpoint       string
	MinioAccessKey string
	MinioSecretKey string
	UseSSL         bool
	MinioEndpoint  string
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.SetConfigType("yml")
	//For Local Environment
	v.AddConfigPath(".")
	//For Docker Images
	v.AddConfigPath("/app/config")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New(constants.CONFIG_NOT_FOUND_ERROR)
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf(constants.CONFIG_DECODE_ERROR, err)
		return nil, err
	}

	return &c, nil
}

func getConfigPath(configPath string) string {
	if configPath == "" {
		return "./config/config"
	}
	return configPath
}

func InitConfig() *Config {
	configPath := getConfigPath(os.Getenv("config"))
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		log.Fatalf(constants.LOADING_CONFIG_ERROR, err)
	}

	cfg, err := ParseConfig(cfgFile)
	AppConfig = cfg
	if err != nil {
		log.Fatalf(constants.PARSING_CONFIG_ERROR, err)
	}
	return AppConfig
}
