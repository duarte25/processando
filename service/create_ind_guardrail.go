package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	accident "processando/acidente"

	"github.com/go-redis/redis/v8"
)

func createDataGuardrail(rdb *redis.Client, ctx context.Context) {

	result := accident.AnalyzeAccidentData("./Acidentes_DadosAbertos_20230412.csv", "ind_guardrail", "ano_acidente")

	// Mapeamento de nomes originais para novos nomes
	nameMapping := map[string]string{
		"NAO INFORMADO": "not_informed",
		"DESCONHECIDO":  "unknown",
		"NAO":           "not",
		"SIM":           "yes",
	}

	// Itera sobre os dados e insere no Redis
	for year, yearData := range result {
		redisKey := fmt.Sprintf("data_guardrail_%s", year)

		for guardrail, count := range yearData.TotalAcciden {

			newGuardrailName, exists := nameMapping[guardrail]
			if !exists {
				newGuardrailName = guardrail
			}

			data, err := json.Marshal(count)
			if err != nil {
				log.Fatalf("Erro ao converter dados para JSON: %v", err)
			}

			err = rdb.HSet(ctx, redisKey, newGuardrailName, data).Err()
			if err != nil {
				log.Fatalf("Erro ao inserir dados no Redis: %v", err)
			}
		}
	}
}
