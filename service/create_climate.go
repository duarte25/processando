package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	accident "processando/acidente"

	"github.com/go-redis/redis/v8"
)

func createDataClimate(rdb *redis.Client, ctx context.Context) {

	result := accident.AnalyzeAccidentData("./Acidentes_DadosAbertos_20230412.csv", "cond_meteorologica", "ano_acidente", "", "")

	// Mapeamento de nomes originais para novos nomes
	nameMapping := map[string]string{
		"CHUVA":                     "rain",
		"CLARO":                     "clear",
		"DESCONHECIDAS":             "unknown",
		"GAROACHUVISCO":             "drizzle",
		"GRANIZO":                   "hail",
		"NAO INFORMADO":             "not_informed",
		"NEVE":                      "snow",
		"NEVOEIRO  NEVOA OU FUMACA": "fog",
		"NUBLADO":                   "cloudy",
		"OUTRAS CONDICOES":          "other_conditions",
		"VENTOS FORTES":             "strong_winds",
	}

	// Itera sobre os dados e insere no Redis
	for year, yearData := range result {
		redisKey := fmt.Sprintf("data_climate_%s", year)

		for climate, count := range yearData.TotalAcciden {
			newClimateName, exists := nameMapping[climate]
			if !exists {
				newClimateName = climate // Use o nome original se não houver mapeamento
			}

			data, err := json.Marshal(count)
			if err != nil {
				log.Fatalf("Erro ao converter dados para JSON: %v", err)
			}

			err = rdb.HSet(ctx, redisKey, newClimateName, data).Err()
			if err != nil {
				log.Fatalf("Erro ao inserir dados no Redis: %v", err)
			}
		}
	}
}
