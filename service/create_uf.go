package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"processando/acidente"
	"processando/src/configs"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

type YearData struct {
	UFData map[string]acidente.UFData `json:"uf_data"`
}

func createDataUF(rdb *redis.Client, ctx context.Context) {
	// Carregar as configurações
	err := configs.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Chama a função para processar os acidentes
	var wg sync.WaitGroup
	start := time.Now()
	wg.Add(5)

	results := make(chan interface{}, 5)

	processYear := func(year string) {
		defer wg.Done()
		result := acidente.Acidente("../acidentes_202304/Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "data_acidente", year)
		results <- result
	}

	go processYear("2018")
	go processYear("2019")
	go processYear("2020")
	go processYear("2021")
	go processYear("2022")

	go func() {
		wg.Wait()
		close(results)
	}()

	// result2018 := acidente.Acidente("./Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "data_acidente", "2018")
	// result2019 := acidente.Acidente("./Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "data_acidente", "2019")
	// result2020 := acidente.Acidente("./Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "data_acidente", "2020")
	// result2021 := acidente.Acidente("./Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "data_acidente", "2021")
	// result2022 := acidente.Acidente("./Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "data_acidente", "2022")

	// Ler os resultados do canal
	for result := range results {
		fmt.Println(result) // Aqui você pode processar os resultados conforme necessário
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed, "/")

	// Criar um mapa para armazenar os dados por ano
	result := make(map[string]YearData)

	// Adicionando os result ao map
	// result["2018"] = YearData{UFData: result2018}
	// result["2019"] = YearData{UFData: result2019}
	// result["2020"] = YearData{UFData: result2020}
	// result["2021"] = YearData{UFData: result2021}
	// result["2022"] = YearData{UFData: result2022}

	// Itera sobre os dados e insere no Redis
	for year, yearData := range result {
		redisKey := fmt.Sprintf("dados_acidentes_%s", year)

		// Iterar sobre os dados de cada estado e inseri-los como hash
		for state, ufData := range yearData.UFData {
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

	fmt.Println("Dados inseridos em redis!")
}
