package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	accident "processando/acidente"

	"github.com/go-redis/redis/v8"
)

func createDataUF(rdb *redis.Client, ctx context.Context) {

	result := accident.AnalyzeAccidentData("./Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "ano_acidente", "", "")

	// Itera sobre os dados e insere no Redis
	// SE PÀ CONSIGO MUDAR ISSO PARA ACONTECER EM OUTRO ARQUIVO AI ACONTECERA TUDO DE UMA VEZ COM TODOS OS DADOS
	for year, yearData := range result {
		redisKey := fmt.Sprintf("data_uf_%s", year)

		// Iterar sobre os dados de cada estado e inseri-los como hash
		for state, ufData := range yearData.TotalAcciden {
			data, err := json.Marshal(ufData)
			if err != nil {
				log.Fatalf("Erro ao converter dados para JSON: %v", err)
			}

			err = rdb.HSet(ctx, redisKey, state, data).Err()
			if err != nil {
				log.Fatalf("Erro ao inserir dados no Redis: %v", err)
			}
		}
	}

}
