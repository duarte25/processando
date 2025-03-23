package routes

import (
	"net/http"
	"net/http/httptest"
	"os"
	"processando/service"
	"processando/src/configs"
	"processando/src/middleware"
	"testing"

	"github.com/go-chi/chi"
)

func TestRegisterRoutes(t *testing.T) {
	// Configura o ambiente para o teste (se necessário)
	os.Setenv("API_PORT", "8080") // Define a porta para o teste

	// Inicializa os serviços (igual ao main)
	err := configs.Load()
	if err != nil {
		t.Fatalf("Erro ao carregar configurações: %v", err)
	}

	service.Controller()

	// Cria o router e registra as rotas (igual ao main)
	r := chi.NewRouter()
	r.Use(middleware.CORS)
	RegisterRoutes(r)

	// Função de teste para as rotas
	tests := []struct {
		method         string
		route          string
		expectedStatus int
	}{
		{"GET", "/uf", http.StatusOK},
		{"GET", "/climate", http.StatusOK},
		{"GET", "/guardrail", http.StatusOK},
		{"GET", "/highway", http.StatusOK},
		{"GET", "/median", http.StatusOK},
		{"GET", "/shoulder", http.StatusOK},
		{"GET", "/speed", http.StatusOK},
		{"GET", "/susp_alcohol", http.StatusOK},
		{"GET", "/day_week", http.StatusOK},
		{"GET", "/phase_day", http.StatusOK},
		{"GET", "/month", http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.route, func(t *testing.T) {
			// Cria uma nova requisição
			req, err := http.NewRequest(tt.method, tt.route, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Cria um gravador de resposta
			rr := httptest.NewRecorder()

			// Chama o handler registrado no roteador
			r.ServeHTTP(rr, req)

			// Verifica se o status da resposta é o esperado
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("Esperado status %v para %v, mas obteve %v", tt.expectedStatus, tt.route, status)
			}
		})
	}
}
