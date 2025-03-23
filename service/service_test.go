package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	accident "processando/acidente"
	"processando/src/configs"
	"testing"
)

func TestValidation(t *testing.T) {

	// Carregar configurações
	err := configs.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	ctx := context.Background()
	rdb := configs.GetRedisClient()

	// Remover a chave específica "data_uf_2021" antes do teste
	err = rdb.Del(ctx, "data_uf_2021").Err()
	if err != nil {
		t.Fatalf("Erro ao deletar a chave 'data_uf_2021': %v", err)
	}

	// Teste 1: Chave não existe
	exists := Validation(rdb, ctx)
	fmt.Printf("Teste 1 - Chave existe? %v\n", exists)
	if exists {
		t.Errorf("Esperava false, mas obteve true")
	}

	exists = Validation(rdb, ctx)

	if exists {
		t.Errorf("Esperava false, mas obteve true")
	}

	// Teste 3: Chave existe com dados válidos
	err = rdb.HSet(ctx, "data_uf_2021", "key1", "value1", "key2", "value2").Err()
	if err != nil {
		t.Fatalf("Erro ao inserir dados no Redis: %v", err)
	}

	exists = Validation(rdb, ctx)
	fmt.Printf("Teste 2 - Chave existe? %v\n", exists)
	if !exists {
		t.Errorf("Esperava true, mas obteve false")
	}

	// Remover a chave novamente após o teste
	err = rdb.Del(ctx, "data_uf_2021").Err()
	if err != nil {
		t.Fatalf("Erro ao deletar a chave 'data_uf_2021' após o teste: %v", err)
	}

	// Populando de volta o banco após apagar tudo
	result := accident.AnalyzeAccidentData("../Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "ano_acidente", "", "")

	for year, yearData := range result {
		redisKey := fmt.Sprintf("data_uf_%s", year)

		// Iterar sobre os dados de cada estado e inseri-los como hash
		for state, ufData := range yearData.TotalAcciden {
			data, err := json.Marshal(ufData)
			if err != nil {
				log.Fatalf("Erro ao converter dados para JSON: %v", err)
			}

			err = rdb.HSet(ctx, redisKey, state, data).Err()
			if err != nil {
				log.Fatalf("Erro ao inserir dados no Redis: %v", err)
			}
		}
	}

}
