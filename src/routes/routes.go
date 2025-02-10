package routes

import (
	"processando/src/handlers"

	"github.com/go-chi/chi"
)

func RegisterRoutes(router *chi.Mux) {
	router.Get("/uf", handlers.ListUF) // Chama Get diretamente
	router.Get("/climate", handlers.ListClimate)
	router.Get("/guardrail", handlers.ListGuardrail)
	router.Get("/highway", handlers.ListHighway)
	router.Get("/median", handlers.ListMedian)
	router.Get("/shoulder", handlers.ListShoulder)
	router.Get("/speed", handlers.ListSpeed)
	router.Get("/susp_alcohol", handlers.ListSuspAlcohol)
	router.Get("/day_week", handlers.ListDayWeek)
	router.Get("/phase_day", handlers.ListPhaseDay)
	router.Get("/month", handlers.ListMonth)
	// Adicione outras rotas aqui, usando handlers do pacote "handlers"
}
