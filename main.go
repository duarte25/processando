package main

import (
	"log"
	"net/http"
	"processando/service"
	"processando/src/configs"
	"processando/src/middleware"
	"processando/src/routes"

	"github.com/go-chi/chi"
)

func main() {

	service.CreateData()
	// Carregar as configurações
	err := configs.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Definindo rotas
	r := chi.NewRouter()
	// Usar middleware CORS
	r.Use(middleware.CORS)

	routes.RegisterRoutes(r)

	// Capturando a porta do arquivo de configuração
	port := configs.GetServerPort()

	// Configurando servidor HTTP
	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
