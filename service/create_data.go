package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"processando/acidente"
	"processando/src/configs"
)

type YearData struct {
	UFData map[string]acidente.UFData `json:"uf_data"`
}

func CreateData() {
	// Carregar as configurações
	err := configs.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	ctx := context.Background()

	// Obter o cliente Redis do pacote configs
	rdb := configs.GetRedisClient()
	defer rdb.Close()

	key := "dados_acidentes_2021"

	// Verificar se a chave existe
	exists, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		log.Fatalf("Erro ao verificar existência da chave: %v", err)
	}

	if exists == 1 {
		// Chave existe, verificar integridade dos dados (exemplo com hash)
		data, err := rdb.HGetAll(ctx, key).Result()
		if err != nil {
			log.Fatalf("Erro ao obter dados da chave: %v", err)
		}
		fmt.Println(data)
		// Verificar se há dados válidos
		if len(data) > 0 {
			// Os dados estão presentes e podem ser processados
			fmt.Println("Dados válidos encontrados para", key)
		} else {
			// A chave existe, mas não há dados válidos associados
			fmt.Println("Chave existe, mas não há dados válidos para", key)
		}
	} else {
		// Chave não existe, nenhum dado foi armazenado ou todos foram removidos
		fmt.Println("Chave não encontrada no Redis:", key)
	}

	// Chama a função para processar os acidentes
	result2021 := acidente.Acidente("./Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "data_acidente", "2021")
	result2022 := acidente.Acidente("./Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "data_acidente", "2022")

	// Criar um mapa para armazenar os dados por ano
	result := make(map[string]YearData)

	// Adicionar os resultados de 2021 e 2022 ao mapa result
	result["2021"] = YearData{UFData: result2021}
	result["2022"] = YearData{UFData: result2022}

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
}
