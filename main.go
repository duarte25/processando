package main

import (
	"log"
	"net/http"
	"processando/src/configs"
	"processando/src/handlers"

	"github.com/go-chi/chi"
)

func main() {
	// Carregar as configurações
	err := configs.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	r := chi.NewRouter()
	// ctx := context.Background()

	// Obter o cliente Redis do pacote configs
	rdb := configs.GetRedisClient()
	defer rdb.Close()

	// Chama a função para processar os acidentes
	// result := acidente.Acidente("./Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "data_acidente", "2022")

	// Itera sobre os dados e insere no Redis
	// for state, ufData := range result {
	// 	// Convertendo UFData para JSON
	// 	data, err := json.Marshal(ufData)
	// 	if err != nil {
	// 		log.Fatalf("Erro ao converter dados para JSON: %v", err)
	// 	}

	// 	err = rdb.HSet(ctx, "dados_acidentes", state, string(data)).Err()
	// 	if err != nil {
	// 		log.Fatalf("Erro ao inserir dados no Redis: %v", err)
	// 	}
	// }

	// Definindo rotas
	r.Get("/", handlers.List)

	// Configurando servidor HTTP
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
