package configs

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

var (
	cfg *Config
	rdb *redis.Client
)

// Config estrutura para as configurações da aplicação
type Config struct {
	API   APIConfig
	Redis RedisConfig
}

// APIConfig estrutura para as configurações da API
type APIConfig struct {
	Port string
}

// RedisConfig estrutura para as configurações do Redis
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// Load carrega as configurações do arquivo e inicializa o cliente Redis
func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(Config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.Redis = RedisConfig{
		Addr:     viper.GetString("redis.addr"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	return nil
}

// GetServerPort retorna a porta do servidor
func GetServerPort() string {
	return cfg.API.Port
}

// GetRedisClient retorna o cliente Redis
func GetRedisClient() *redis.Client {
	return rdb
}
