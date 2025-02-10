package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	accident "processando/acidente"

	"github.com/go-redis/redis/v8"
)

func createDataDayWeek(rdb *redis.Client, ctx context.Context) {

	result := accident.AnalyzeAccidentData("./Acidentes_DadosAbertos_20230412.csv", "dia_semana", "ano_acidente", "", "")

	// Mapeamento de nomes originais para novos nomes
	nameMapping := map[string]string{
		"DOMINGO":       "sunday",
		"QUARTA-FEIRA":  "wednesday",
		"QUINTA-FEIRA":  "thursday",
		"SABADO":        "saturday",
		"SEGUNDA-FEIRA": "monday",
		"SEXTA-FEIRA":   "friday",
		"TERCA-FEIRA":   "tuesday",
	}

	// Itera sobre os dados e insere no Redis
	for year, yearData := range result {
		redisKey := fmt.Sprintf("data_day_week_%s", year)

		for day, count := range yearData.TotalAcciden {

			newDayName, exists := nameMapping[day]
			if !exists {
				newDayName = day
			}

			data, err := json.Marshal(count)
			if err != nil {
				log.Fatalf("Erro ao converter dados para JSON: %v", err)
			}

			err = rdb.HSet(ctx, redisKey, newDayName, data).Err()
			if err != nil {
				log.Fatalf("Erro ao inserir dados no Redis: %v", err)
			}
		}
	}
}
