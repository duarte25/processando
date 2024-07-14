package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"processando/acidente"
	"processando/src/configs"
	"processando/src/routes"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type YearData struct {
	UFData map[string]acidente.UFData `json:"uf_data"`
}

func main() {
	// Carregar as configurações
	err := configs.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	ctx := context.Background()

	// Obter o cliente Redis do pacote configs
	rdb := configs.GetRedisClient()
	defer rdb.Close()

	// Chama a função para processar os acidentes
	result2021 := acidente.Acidente("./Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "data_acidente", "2021")
	result2022 := acidente.Acidente("./Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "data_acidente", "2022")

	// Criar um mapa para armazenar os dados por ano
	result := make(map[string]YearData)

	// Adicionar os resultados de 2021 e 2022 ao mapa result
	result["2021"] = YearData{UFData: result2021}
	result["2022"] = YearData{UFData: result2022}

	// Exemplo de acesso aos dados
	fmt.Println("Dados de result:", result)

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

	// Definindo rotas
	r := chi.NewRouter()
	// Configurando CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Permita apenas o frontend local
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Cache pré-checada por 5 minutos
	}))

	routes.RegisterRoutes(r)

	// Capturando a porta do arquivo de configuração
	port := configs.GetServerPort()

	// Configurando servidor HTTP
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
