package configs

import (
	"os"
	"strconv" // Para converter string para int

	"github.com/go-redis/redis/v8"
)

var (
	cfg *Config
	rdb *redis.Client
)

// Config estrutura para as configurações da aplicação
type Config struct {
	Redis RedisConfig
}

// RedisConfig estrutura para as configurações do Redis
type RedisConfig struct {
	Addr     string
	Password string
	DB       int
}

// Load carrega as configurações das variáveis de ambiente e inicializa o cliente Redis
func Load() error {
	// Converte a variável REDIS_DB para int
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		db = 0 // Valor padrão caso REDIS_DB não esteja definido corretamente
	}

	// Configurações do Redis obtidas diretamente das variáveis de ambiente
	cfg = &Config{
		Redis: RedisConfig{
			Addr:     os.Getenv("REDIS_ADDR"),     // Endereço do Redis
			Password: os.Getenv("REDIS_PASSWORD"), // Senha do Redis
			DB:       db,                          // Banco de dados do Redis (inteiro)
		},
	}

	// Inicializar o cliente Redis com as configurações carregadas
	rdb = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	return nil
}

// GetRedisClient retorna o cliente Redis
func GetRedisClient() *redis.Client {
	return rdb
}
