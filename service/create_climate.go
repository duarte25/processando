package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	accident "processando/acidente"
	"processando/src/configs"

	"github.com/go-redis/redis/v8"
)

func createDataClimate(rdb *redis.Client, ctx context.Context) {
	// Carregar as configurações
	err := configs.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	result := accident.AnalyzeAccidentData("./Acidentes_DadosAbertos_20230412.csv", "cond_meteorologica", "ano_acidente")

	// Itera sobre os dados e insere no Redis
	for year, yearData := range result {
		redisKey := fmt.Sprintf("dados_climate_%s", year)

		// Iterar sobre os dados de cada climate e inseri-los como hash
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
