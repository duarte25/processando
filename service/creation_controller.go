package service

import (
	"context"
	"fmt"
	"log"
	"processando/src/configs"
	"time"

	"github.com/go-redis/redis/v8"
)

func Controller() {

	ctx := context.Background()

	// Obter o cliente Redis do pacote configs
	rdb := configs.GetRedisClient()

	start := time.Now()

	if !validation(rdb, ctx) {
		createDataUF(rdb, ctx)
		createDataClimate(rdb, ctx)
		createDataHighway(rdb, ctx)
		createDataSpeed(rdb, ctx)
		createDataShoulder(rdb, ctx)
		createDataGuardrail(rdb, ctx)
		createDataMedian(rdb, ctx)
	}

	elapsed := time.Since(start)
	fmt.Println(elapsed, "/")

	fmt.Println("Dados inseridos em redis!")
}

func validation(rdb *redis.Client, ctx context.Context) bool {
	key := "data_uf_2021"

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

		// Verificar se há dados válidos
		if len(data) > 0 {
			// Os dados estão presentes e podem ser processados
			fmt.Println("Dados válidos encontrados para", key)
			return true
		} else {
			// A chave existe, mas não há dados válidos associados
			fmt.Println("Chave existe, mas não há dados válidos para", key)
			return false
		}
	}

	fmt.Println("Chave não encontrada no Redis:", key)
	return false
}
