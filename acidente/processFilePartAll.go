package accident

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func ProcessFilePartAll(
	filePath string,
	startOffset, endOffset int64,
	columns []string,
	idxUF, idxDate, idxClimate, idxSpeed, idxTrack, idxPhaseDay, idxMonth, idxDayWeek, idxShoulder, idxGuardrail, idxMedian, idxDeaths, idxInjured, idxInvolved int,
	wg *sync.WaitGroup,
	results *sync.Map,
) {
	defer wg.Done()
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Erro ao abrir arquivo: %v\n", err)
		return
	}
	defer file.Close()

	// Pular até startOffset
	_, err = file.Seek(startOffset, 0)
	if err != nil {
		log.Printf("Erro ao posicionar no offset: %v\n", err)
		return
	}

	scanner := bufio.NewScanner(file)

	// Buffer para ler linha por linha
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		// Garantir que estamos dentro do intervalo desejado
		currentPos, _ := file.Seek(0, 1)
		if currentPos > endOffset {
			break
		}

		row := strings.Split(line, ";")

		year := safeGetValue(row, idxDate)
		deaths := parseSafeInt(safeGetValue(row, idxDeaths))
		injured := parseSafeInt(safeGetValue(row, idxInjured))
		involved := parseSafeInt(safeGetValue(row, idxInvolved))

		// Função auxiliar para atualizar o sync.Map com dados de uma categoria
		updateMap := func(keyPrefix, category string, dataKey string) {
			mapKey := keyPrefix + "_" + year
			val, _ := results.LoadOrStore(mapKey, &YearData{
				TotalAcciden: make(map[string]*AccidentData),
				mu:           sync.Mutex{},
			})
			yearData := val.(*YearData)

			yearData.mu.Lock()
			item, exists := yearData.TotalAcciden[dataKey]
			if !exists {
				item = &AccidentData{}
			}

			item.TotalAccident++
			item.TotalDeath += deaths
			item.TotalInjured += injured
			item.TotalInvolved += involved

			yearData.TotalAcciden[dataKey] = item
			yearData.mu.Unlock()
		}

		// Processar cada categoria
		if idxUF >= 0 {
			updateMap("uf", "uf_acidente", safeGetValue(row, idxUF))
		}
		if idxClimate >= 0 {
			updateMap("climate", "clima", safeGetValue(row, idxClimate))
		}
		if idxSpeed >= 0 {
			updateMap("speed", "velocidade_maxima", safeGetValue(row, idxSpeed))
		}
		if idxTrack >= 0 {
			updateMap("track", "pavimento", safeGetValue(row, idxTrack))
		}
		if idxPhaseDay >= 0 {
			updateMap("phase_day", "fase_dia", safeGetValue(row, idxPhaseDay))
		}
		if idxMonth >= 0 {
			updateMap("month", "mes_acidente", safeGetValue(row, idxMonth))
		}
		if idxDayWeek >= 0 {
			updateMap("day_week", "dia_semana", safeGetValue(row, idxDayWeek))
		}
		if idxShoulder >= 0 {
			updateMap("shoulder", "acostamento", safeGetValue(row, idxShoulder))
		}
		if idxGuardrail >= 0 {
			updateMap("guardrail", "guard_rail", safeGetValue(row, idxGuardrail))
		}
		if idxMedian >= 0 {
			updateMap("median", "canteiro_central", safeGetValue(row, idxMedian))
		}
	}
}

func parseSafeInt(s string) int {
	if s == "" {
		return 0
	}
	if i, err := strconv.Atoi(s); err == nil {
		return i
	}
	return 0
}

func safeGetValue(row []string, idx int) string {
	if idx >= 0 && idx < len(row) {
		return row[idx]
	}
	return ""
}
