package service

import (
	"context"
	"fmt"
	"log"
	"testing"

	"processando/src/configs"
)

func TestService(t *testing.T) {
	err := configs.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	ctx := context.Background()
	rdb := configs.GetRedisClient()

	err = rdb.Del(ctx, "data_uf_2021").Err()
	if err != nil {
		t.Fatalf("Erro ao deletar a chave 'data_uf_2021': %v", err)
	}

	fmt.Printf("Teste 1 - Verificando chave inexistente...\n")
	exists := Validation(rdb, ctx)
	if exists {
		t.Errorf("Esperava false para chave inexistente, mas obteve true")
	}

	// Testar novamente para garantir consistência se a função for chamada múltiplas vezes
	exists = Validation(rdb, ctx)
	if exists {
		t.Errorf("Esperava false para chave inexistente (segunda verificação), mas obteve true")
	}

	fmt.Printf("Teste 2 - Populando Redis diretamente para simular chave existente...\n")
	err = rdb.HSet(ctx, "data_uf_2021", "SP", "some_data_json", "RJ", "other_data_json").Err()
	if err != nil {
		t.Fatalf("Erro ao inserir dados diretamente no Redis para teste: %v", err)
	}

	fmt.Printf("Teste 2 - Verificando chave existente...\n")
	exists = Validation(rdb, ctx)
	if !exists {
		t.Errorf("Esperava true para chave existente, mas obteve false")
	}

	// --- Limpeza Final ---
	// Remover a chave "data_uf_2021" após o teste para não afetar outros testes
	err = rdb.Del(ctx, "data_uf_2021").Err()
	if err != nil {
		t.Fatalf("Erro ao deletar a chave 'data_uf_2021' após o teste: %v", err)
	}

	fmt.Printf("Teste de validação concluído.\n")
}
