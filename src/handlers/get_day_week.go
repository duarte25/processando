package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"processando/src/configs"
	"processando/src/entities"

	"github.com/go-redis/redis/v8"
)

// @Summary Listar tabelas de dia da semana (Day Week)
// @Description Retorna uma lista com os valores de dia da semana e na base do parâmetro e com os valores de total de acidentes, total de óbitos, total de envolvidos e feridos.
// @Tags day_week
// @Accept json
// @Produce json
// @Param data query string true "Identificador dos dados a serem recuperados (ex.: data_day_week_2022)" sendo os 4 últimos dígitos o ano a ser pesquisado, estando presente na base de dados os anos de 2018 a 2023.
// @Success 200 {array} string "Lista de guarda corpo"
// @Router /day_week [get]
func ListDayWeek(w http.ResponseWriter, r *http.Request) {
	// Obter o parâmetro da URL "dados"
	redisKey := r.URL.Query().Get("data")

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

	dayWeekDatas := make(map[string]entities.Accident)

	for dayWeek, jsonData := range data {
		var datas entities.Accident
		err := json.Unmarshal([]byte(jsonData), &datas)

		if err != nil {
			http.Error(w, "Error ao desserializar dados do day week "+dayWeek, http.StatusInternalServerError)
			return
		}

		dayWeekDatas[dayWeek] = datas
	}

	jsonData, err := json.Marshal(dayWeekDatas)
	if err != nil {
		http.Error(w, "Erro ao converter dados para JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
