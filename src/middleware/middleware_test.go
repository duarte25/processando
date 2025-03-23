package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCORS(t *testing.T) {
	// Cria um handler simples para testar o middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Aplica o middleware CORS ao handler
	corsHandler := CORS(handler)

	// Cria uma requisição HTTP OPTIONS simulada (pré-checagem CORS)
	req, err := http.NewRequest("OPTIONS", "http://example.com/", nil)
	if err != nil {
		t.Fatal("Erro ao criar a requisição:", err)
	}

	// Define o cabeçalho Origin para simular uma requisição de um frontend específico
	req.Header.Set("Origin", "http://localhost:3022")

	// Cria um ResponseRecorder para gravar a resposta
	rr := httptest.NewRecorder()

	// Executa o handler com o middleware CORS
	corsHandler.ServeHTTP(rr, req)

	// Verifica o status code da resposta
	if rr.Code != http.StatusOK {
		t.Errorf("Status code esperado: %d, obtido: %d", http.StatusOK, rr.Code)
	}

	// Log dos cabeçalhos retornados para depuração
	t.Logf("Cabeçalhos retornados: %v", rr.Header())

	// Verifica os cabeçalhos CORS na resposta
	expectedHeaders := map[string]string{
		"Access-Control-Allow-Origin":      "http://localhost:3022",
		"Access-Control-Allow-Credentials": "true",
		"Access-Control-Expose-Headers":    "Link",
		"Vary":                             "Origin",
	}

	for header, expectedValue := range expectedHeaders {
		actualValue := rr.Header().Get(header)
		if actualValue != expectedValue {
			t.Errorf("Cabeçalho %s esperado: %s, obtido: %s", header, expectedValue, actualValue)
		}
	}
}

// TestCORS_InvalidOrigin verifica se o middleware CORS bloqueia origens não permitidas
func TestCORS_InvalidOrigin(t *testing.T) {
	// Cria um handler simples para testar o middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Aplica o middleware CORS ao handler
	corsHandler := CORS(handler)

	// Cria uma requisição HTTP GET simulada com uma origem não permitida
	req, err := http.NewRequest("GET", "http://example.com/", nil)
	if err != nil {
		t.Fatal("Erro ao criar a requisição:", err)
	}

	// Define o cabeçalho Origin para uma origem não permitida
	req.Header.Set("Origin", "http://invalid-origin.com")

	// Cria um ResponseRecorder para gravar a resposta
	rr := httptest.NewRecorder()

	// Executa o handler com o middleware CORS
	corsHandler.ServeHTTP(rr, req)

	// Verifica o status code da resposta
	if rr.Code != http.StatusOK {
		t.Errorf("Status code esperado: %d, obtido: %d", http.StatusOK, rr.Code)
	}

	// Verifica se o cabeçalho Access-Control-Allow-Origin não está presente
	if rr.Header().Get("Access-Control-Allow-Origin") != "" {
		t.Error("Cabeçalho Access-Control-Allow-Origin presente para uma origem não permitida")
	}
}
