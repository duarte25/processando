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

func createDataSusAlcool(rdb *redis.Client, ctx context.Context) {

	file := os.Getenv("VITIMA_FILE")

	result := accident.AnalyzeAccidentData(file, "susp_alcool", "ano_acidente", "MOTORISTA", "tp_envolvido")

	nameMapping := map[string]string{
		"NAO INFORMADO": "note_informed",
		"DESCONHECIDO":  "unknown",
		"SIM":           "yes",
		"NAO":           "not",
		"NAO APLICAVEL": "not_applicable",
	}

	for year, yearData := range result {
		redisKey := fmt.Sprintf("data_susp_alcohol_%s", year)

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
