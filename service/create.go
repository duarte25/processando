package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	accident "processando/acidente"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type DataConfig struct {
	File           string // Caminho do arquivo CSV
	IndexColumn    string // Coluna usada como chave (ex: uf_acidente)
	DateColumn     string // Coluna da data (ex: ano_acidente)
	FilterColumn   string // Coluna opcional para filtro (ex: clima)
	FilterValue    string // Valor opcional de filtro
	RedisKeyPrefix string // Prefixo da chave Redis (ex: data_uf_)
}

func CreateData(rdb *redis.Client, ctx context.Context) {
	file := os.Getenv("ACIDENTE_FILE")
	if file == "" {
		log.Fatal("Variável ACIDENTE_FILE não definida")
	}

	result := accident.AnalyzeAll(file)

	insertToRedis(rdb, ctx, result.YearDataByUF, "data_uf_")
	insertToRedis(rdb, ctx, result.YearDataByClimate, "data_climate_")
	insertToRedis(rdb, ctx, result.YearDataBySpeed, "data_speed_")
	insertToRedis(rdb, ctx, result.YearDataByTrack, "data_track_condition_")
	insertToRedis(rdb, ctx, result.YearDataByPhaseDay, "data_phase_day_")
	insertToRedis(rdb, ctx, result.YearDataByMonth, "data_month_")
	insertToRedis(rdb, ctx, result.YearDataByDayWeek, "data_day_week_")
	insertToRedis(rdb, ctx, result.YearDataByShoulder, "data_shoulder_")
	insertToRedis(rdb, ctx, result.YearDataByGuardrail, "data_guardrail_")
	insertToRedis(rdb, ctx, result.YearDataByMedian, "data_median_")
}

func insertToRedis(rdb *redis.Client, ctx context.Context, data map[string]*accident.YearData, prefix string) {
	for year, yearData := range data {
		// Verifica se o "year" é um número de 4 dígitos
		if _, err := strconv.Atoi(year); err != nil || len(year) != 4 {
			// log.Printf("Ignorando ano inválido ou não numérico: %q", year)
			continue
		}

		redisKey := fmt.Sprintf("%s%s", prefix, year)

		for key, item := range yearData.TotalAcciden {
			jsonData, err := json.Marshal(item)
			if err != nil {
				log.Printf("Erro ao serializar JSON (%s): %v", redisKey, err)
				continue
			}

			err = rdb.HSet(ctx, redisKey, key, jsonData).Err()
			if err != nil {
				log.Printf("Erro ao inserir no Redis (%s): %v", redisKey, err)
			}
		}
	}
}
