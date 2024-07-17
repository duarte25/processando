package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	accident "processando/acidente"
	"processando/src/configs"
	"time"

	"github.com/go-redis/redis/v8"
)

// type YearData struct {
// 	UFData map[string]acidente.UFData `json:"uf_data"`
// }

func createDataUF(rdb *redis.Client, ctx context.Context) {
	// Carregar as configurações
	err := configs.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	start := time.Now()
	result := accident.AnalyzeAccidentData("./Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "ano_acidente")

	elapsed := time.Since(start)
	fmt.Println(elapsed, "/")

	// Itera sobre os dados e insere no Redis
	for year, yearData := range result {
		redisKey := fmt.Sprintf("dados_acidentes_%s", year)

		// Iterar sobre os dados de cada estado e inseri-los como hash
		for state, ufData := range yearData.UFs {
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

	fmt.Println("Dados inseridos em redis!")
}
