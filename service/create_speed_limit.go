package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	accident "processando/acidente"

	"github.com/go-redis/redis/v8"
)

func createDataSpeed(rdb *redis.Client, ctx context.Context) {

	result := accident.AnalyzeAccidentData("./Acidentes_DadosAbertos_20230412.csv", "lim_velocidade", "ano_acidente")

	// Mapeamento de nomes originais para novos nomes
	nameMapping := map[string]string{
		"60 KMH":        "60_kmh",
		"40 KMH":        "40_kmh",
		"80 KMH":        "80_kmh",
		"30 KMH":        "30_kmh",
		"110 KMH":       "110_kmh",
		"NAO INFORMADO": "not_informed",
	}

	// Itera sobre os dados e insere no Redis
	for year, yearData := range result {
		redisKey := fmt.Sprintf("data_speed_%s", year)

		for speed, count := range yearData.TotalAcciden {

			newSpeedLimit, exists := nameMapping[speed]

			if !exists {
				newSpeedLimit = speed
			}

			data, err := json.Marshal(count)
			if err != nil {
				log.Fatalf("Erro ao converter dados para JSON: %v", err)
			}

			err = rdb.HSet(ctx, redisKey, newSpeedLimit, data).Err()
			if err != nil {
				log.Fatalf("Erro ao inserir dados no Redis: %v", err)
			}
		}
	}
}
