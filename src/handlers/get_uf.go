package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"processando/src/configs"
	"processando/src/entities"

	"github.com/go-redis/redis/v8"
)

func ListUF(w http.ResponseWriter, r *http.Request) {

	// Obter o parâmetro da URL "dados"
	redisKey := r.URL.Query().Get("dados")

	// Definindo um contexto para o Redis
	ctx := context.Background()

	// Obter o cliente Redis do pacote configs
	rdb := configs.GetRedisClient()

	// Tentando buscar os dados no Redis como um hash
	data, err := rdb.HGetAll(ctx, redisKey).Result()

	if err == redis.Nil {
		// Se os dados não estiverem no Redis, retorne um erro ou uma mensagem apropriada
		http.Error(w, "Nenhum registro encontrado no cache", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Erro ao buscar registros no cache", http.StatusInternalServerError)
		return
	}

	stateDatas := make(map[string]entities.UF)

	for state, jsonData := range data {
		var datas entities.UF
		err := json.Unmarshal([]byte(jsonData), &datas)

		if err != nil {
			http.Error(w, "Erro ao desserializar dados do estado "+state, http.StatusInternalServerError)
			return
		}

		stateDatas[state] = datas
	}

	// Se os dados estiverem no Redis, converta para JSON e retorne-os
	jsonData, err := json.Marshal(stateDatas)
	if err != nil {
		http.Error(w, "Erro ao converter dados para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
