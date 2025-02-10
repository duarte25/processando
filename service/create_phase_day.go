package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	accident "processando/acidente"

	"github.com/go-redis/redis/v8"
)

func createDataPhaseDay(rdb *redis.Client, ctx context.Context) {

	result := accident.AnalyzeAccidentData("./Acidentes_DadosAbertos_20230412.csv", "fase_dia", "ano_acidente", "", "")

	// Mapeamento de nomes originais para novos nomes
	nameMapping := map[string]string{
		"DESCONHECIDO":  "unknown",
		"MADRUGADA":     "dawn",
		"MANHA":         "morning",
		"NAO INFORMADO": "not_informed",
		"NOITE":         "night",
		"TARDE":         "afternoon",
	}

	// Itera sobre os dados e insere no Redis
	for year, yearData := range result {
		redisKey := fmt.Sprintf("data_phase_day_%s", year)

		for day, count := range yearData.TotalAcciden {

			newPhaseDayName, exists := nameMapping[day]
			if !exists {
				newPhaseDayName = day
			}

			data, err := json.Marshal(count)
			if err != nil {
				log.Fatalf("Erro ao converter dados para JSON: %v", err)
			}

			err = rdb.HSet(ctx, redisKey, newPhaseDayName, data).Err()
			if err != nil {
				log.Fatalf("Erro ao inserir dados no Redis: %v", err)
			}
		}
	}
}
