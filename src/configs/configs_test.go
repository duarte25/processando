package configs

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	// Limpar variáveis de ambiente antes do teste
	os.Clearenv()

	// Definir variáveis de ambiente simuladas
	os.Setenv("REDIS_ADDR", "localhost:6379")
	os.Setenv("REDIS_PASSWORD", "mysecretpassword")
	os.Setenv("REDIS_DB", "1")

	// Carregar as configurações
	err := Load()
	if err != nil {
		t.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Validar configurações do Redis
	if cfg.Redis.Addr != "localhost:6379" {
		t.Errorf("Endereço do Redis incorreto. Esperado: 'localhost:6379', Obtido: '%s'", cfg.Redis.Addr)
	}
	if cfg.Redis.Password != "mysecretpassword" {
		t.Errorf("Senha do Redis incorreta. Esperado: 'mysecretpassword', Obtido: '%s'", cfg.Redis.Password)
	}
	if cfg.Redis.DB != 1 {
		t.Errorf("Banco de dados do Redis incorreto. Esperado: 1, Obtido: %d", cfg.Redis.DB)
	}

	// Validar cliente Redis
	rdb := GetRedisClient()
	if rdb == nil {
		t.Fatal("Cliente Redis não foi inicializado")
	}
}

func TestLoadWithMissingEnvVars(t *testing.T) {
	// Limpar variáveis de ambiente antes do teste
	os.Clearenv()

	// Carregar as configurações sem definir variáveis de ambiente
	err := Load()
	if err != nil {
		t.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Validar valores padrão
	if cfg.Redis.Addr != "" {
		t.Errorf("Endereço do Redis deveria ser vazio. Obtido: '%s'", cfg.Redis.Addr)
	}
	if cfg.Redis.Password != "" {
		t.Errorf("Senha do Redis deveria ser vazia. Obtido: '%s'", cfg.Redis.Password)
	}
	if cfg.Redis.DB != 0 {
		t.Errorf("Banco de dados do Redis deveria ser 0. Obtido: %d", cfg.Redis.DB)
	}
}

func TestGetRedisClient(t *testing.T) {
	// Limpar variáveis de ambiente antes do teste
	os.Clearenv()

	// Definir variáveis de ambiente simuladas
	os.Setenv("REDIS_ADDR", "localhost:6379")
	os.Setenv("REDIS_PASSWORD", "mysecretpassword")
	os.Setenv("REDIS_DB", "1")

	// Carregar as configurações
	err := Load()
	if err != nil {
		t.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Obter o cliente Redis
	rdb := GetRedisClient()
	if rdb == nil {
		t.Fatal("Cliente Redis não foi inicializado")
	}

	// Verificar se o cliente Redis é consistente
	client := GetRedisClient()
	if rdb != client {
		t.Error("O cliente Redis retornado é inconsistente")
	}
}
