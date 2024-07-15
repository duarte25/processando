package routes

import (
	"processando/src/handlers"

	"github.com/go-chi/chi"
)

func RegisterRoutes(router *chi.Mux) {
	router.Get("/uf", handlers.ListUF) // Chama Get diretamente
	// Adicione outras rotas aqui, usando handlers do pacote "handlers"
}
