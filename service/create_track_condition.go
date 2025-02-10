package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	accident "processando/acidente"

	"github.com/go-redis/redis/v8"
)

func createDataTrackCondition(rdb *redis.Client, ctx context.Context) {

	result := accident.AnalyzeAccidentData("./Acidentes_DadosAbertos_20230412.csv", "cond_pista", "ano_acidente", "", "")

	// Mapeamento de nomes originais para novos nomes
	nameMapping := map[string]string{
		"COM BURACO":             "hole",
		"COM LAMA":               "mud",
		"COM MATERIAL GRANULADO": "granular_material",
		"DESCONHECIDO":           "unknown",
		"ESCORREGADIA":           "slippery",
		"MOLHADA":                "wet",
		"NAO INFORMADO":          "not_informed",
		"OBSTRUIDA":              "obstructed",
		"SECA":                   "dry",
	}

	// Itera sobre os dados e insere no Redis
	for year, yearData := range result {
		redisKey := fmt.Sprintf("data_track_condition_%s", year)

		for day, count := range yearData.TotalAcciden {

			newTrackCondition, exists := nameMapping[day]
			if !exists {
				newTrackCondition = day
			}

			data, err := json.Marshal(count)
			if err != nil {
				log.Fatalf("Erro ao converter dados para JSON: %v", err)
			}

			err = rdb.HSet(ctx, redisKey, newTrackCondition, data).Err()
			if err != nil {
				log.Fatalf("Erro ao inserir dados no Redis: %v", err)
			}
		}
	}
}
