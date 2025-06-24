package main

import (
	"log"
	"net/http"
	"os" // Pacote para ler variáveis de ambiente
	"os/signal"
	"processando/service"
	"processando/src/configs"
	"processando/src/middleware"
	"processando/src/routes"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	// Inicializar os serviços
	err := configs.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar .env")
	}

	// f, err := os.Create("trace.out")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// if err := trace.Start(f); err != nil {
	// 	log.Fatal(err)
	// }
	// defer trace.Stop()
	// defer f.Close()

	service.Controller()

	// start := time.Now()

	// sequencial.Acidente("acidentes.csv")

	// duration := time.Since(start) // calcula a duração
	// fmt.Printf("Tempo de execução: %s\n", duration)

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

	// Configurar um canal para capturar sinais do sistema
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Iniciar o servidor HTTP em uma goroutine
	go func() {
		log.Printf("Iniciando servidor na porta %s", port)
		if err := http.ListenAndServe(":"+port, r); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Erro ao iniciar o servidor: %v", err)
		}
	}()

	// Obter o cliente Redis do pacote configs
	rdb := configs.GetRedisClient()

	// Aguardar por um sinal de encerramento
	<-c
	log.Println("Desligando o servidor e fechando a conexão com o Redis...")

	// Fechar a conexão com o Redis antes de encerrar
	if err := rdb.Close(); err != nil {
		log.Printf("Erro ao fechar a conexão com o Redis: %v", err)
	} else {
		log.Println("Conexão com o Redis fechada com sucesso.")
	}

	log.Println("Servidor encerrado.")
}
