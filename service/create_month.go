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

func createDataMonth(rdb *redis.Client, ctx context.Context) {

	file := os.Getenv("ACIDENTE_FILE")

	result := accident.AnalyzeAccidentData(file, "mes_acidente", "ano_acidente", "", "")

	// Itera sobre os dados e insere no Redis
	for year, yearData := range result {
		redisKey := fmt.Sprintf("data_month_%s", year)

		for month, count := range yearData.TotalAcciden {

			data, err := json.Marshal(count)
			if err != nil {
				log.Fatalf("Erro ao converter dados para JSON: %v", err)
			}

			err = rdb.HSet(ctx, redisKey, month, data).Err()
			if err != nil {
				log.Fatalf("Erro ao inserir dados no Redis: %v", err)
			}
		}
	}
}
