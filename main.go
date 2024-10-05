package main

import (
	"log"
	"net/http"
	"os" // Pacote para ler variáveis de ambiente
	"processando/service"
	"processando/src/configs"
	"processando/src/middleware"
	"processando/src/routes"

	"github.com/go-chi/chi"
)

func main() {
	// Inicializar os serviços
	err := configs.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	service.Controller()

	// Definindo rotas
	r := chi.NewRouter()

	// Usar middleware CORS
	r.Use(middleware.CORS)

	// Registrar as rotas
	routes.RegisterRoutes(r)

	// Pegar a porta diretamente da variável de ambiente API_PORT
	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080" // Definir porta padrão se a variável de ambiente não estiver definida
	}

	// Configurar e iniciar o servidor HTTP
	log.Printf("Iniciando servidor na porta %s", port)

	err = http.ListenAndServe(":"+port, r)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
