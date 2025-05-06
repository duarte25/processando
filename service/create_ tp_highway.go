package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	accident "processando/acidente"

	"github.com/go-redis/redis/v8"
)

func createDataHighway(rdb *redis.Client, ctx context.Context) {

	file := os.Getenv("ACIDENTE_FILE")

	result := accident.AnalyzeAccidentData(file, "tp_pavimento", "ano_acidente", "", "")

	// Mapeamento de nomes originais para novos nomes
	nameMapping := map[string]string{
		"CASCALHO":       "gravel",
		"NAO INFORMADO":  "not_informed",
		"DESCONHECIDO":   "unknown",
		"ASFALTO":        "asphalt",
		"CONCRETO":       "concrete",
		"PARALELEPIPEDO": "paving_stone",
		"TERRA":          "earth",
	}

	// Itera sobre os dados e insere no Redis
	for year, yearData := range result {
		redisKey := fmt.Sprintf("data_highway_%s", year)

		for highway, count := range yearData.TotalAcciden {

			newHighwayName, exists := nameMapping[highway]
			if !exists {
				newHighwayName = highway
			}

			data, err := json.Marshal(count)
			if err != nil {
				log.Fatalf("Erro ao converter dados para JSON: %v", err)
			}

			err = rdb.HSet(ctx, redisKey, newHighwayName, data).Err()
			if err != nil {
				log.Fatalf("Erro ao inserir dados no Redis: %v", err)
			}
		}
	}
}
