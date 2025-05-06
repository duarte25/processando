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

func createDataMedian(rdb *redis.Client, ctx context.Context) {

	file := os.Getenv("ACIDENTE_FILE")

	result := accident.AnalyzeAccidentData(file, "ind_cantcentral", "ano_acidente", "", "")

	// Mapeamento de nomes originais para novos nomes
	nameMapping := map[string]string{
		"NAO INFORMADO": "not_informed",
		"DESCONHECIDO":  "unknown",
		"NAO":           "not",
		"SIM":           "yes",
	}

	// Itera sobre os dados e insere no Redis
	for year, yearData := range result {
		redisKey := fmt.Sprintf("data_median_%s", year)

		for median, count := range yearData.TotalAcciden {

			newMedianName, exists := nameMapping[median]
			if !exists {
				newMedianName = median
			}

			data, err := json.Marshal(count)
			if err != nil {
				log.Fatalf("Erro ao converter dados para JSON: %v", err)
			}

			err = rdb.HSet(ctx, redisKey, newMedianName, data).Err()
			if err != nil {
				log.Fatalf("Erro ao inserir dados no Redis: %v", err)
			}
		}
	}
}
