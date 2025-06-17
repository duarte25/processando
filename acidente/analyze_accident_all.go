package accident

import (
	"bufio"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	numParts = 8 // Número de partes para dividir o arquivo
)

type AnalysisResult struct {
	YearDataByUF        map[string]*YearData
	YearDataByClimate   map[string]*YearData
	YearDataBySpeed     map[string]*YearData
	YearDataByTrack     map[string]*YearData
	YearDataByPhaseDay  map[string]*YearData
	YearDataByMonth     map[string]*YearData
	YearDataByDayWeek   map[string]*YearData
	YearDataByShoulder  map[string]*YearData
	YearDataByGuardrail map[string]*YearData
	YearDataByMedian    map[string]*YearData
	YearDataByHighway   map[string]*YearData
}

func AnalyzeAll(filePath string) *AnalysisResult {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Erro ao abrir o arquivo:", err)
	}
	defer file.Close()

	// Ler cabeçalho
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		log.Fatal("Erro ao ler cabeçalho")
	}
	header := scanner.Text()
	columns := strings.Split(header, ";")

	// Mapear índices de todas as colunas que queremos
	idxUF := getColumnIndex(columns, "uf_acidente")
	idxDate := getColumnIndex(columns, "ano_acidente")
	idxClimate := getColumnIndex(columns, "cond_meteorologica")
	idxSpeed := getColumnIndex(columns, "lim_velocidade")
	idxTrack := getColumnIndex(columns, "cond_pista")
	idxPhaseDay := getColumnIndex(columns, "fase_dia")
	idxMonth := getColumnIndex(columns, "mes_acidente")
	idxDayWeek := getColumnIndex(columns, "dia_semana")
	idxShoulder := getColumnIndex(columns, "ind_acostamento")
	idxGuardrail := getColumnIndex(columns, "ind_guardrail")
	idxMedian := getColumnIndex(columns, "ind_cantcentral")
	idxHighway := getColumnIndex(columns, "tp_pavimento")

	idxDeaths := getColumnIndex(columns, "qtde_obitos")
	idxInjured := getColumnIndex(columns, "qtde_feridosilesos")
	idxInvolved := getColumnIndex(columns, "qtde_envolvidos")

	if idxUF == -1 || idxDate == -1 || idxDeaths == -1 || idxInjured == -1 || idxInvolved == -1 {
		log.Fatal("Coluna crítica não encontrada no cabeçalho")
	}

	// Dividir o arquivo para paralelismo
	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	partSize := fileSize / int64(numParts)

	var wg sync.WaitGroup
	var partialResults sync.Map // chave: tipo de análise (ex: "uf", "climate") -> valor: map[string]*YearData

	// Processar cada parte do arquivo
	for i := 0; i < numParts; i++ {
		startOffset := int64(i) * partSize
		endOffset := startOffset + partSize
		if i == numParts-1 {
			endOffset = fileSize
		}

		wg.Add(1)
		go ProcessFilePartAll(
			filePath,
			startOffset,
			endOffset,
			columns,
			idxUF,
			idxDate,
			idxClimate,
			idxSpeed,
			idxTrack,
			idxPhaseDay,
			idxMonth,
			idxDayWeek,
			idxShoulder,
			idxGuardrail,
			idxMedian,
			idxHighway,
			idxDeaths,
			idxInjured,
			idxInvolved,
			&wg,
			&partialResults,
		)
	}

	wg.Wait()

	// Agregar resultados
	result := &AnalysisResult{
		YearDataByUF:        mergeYearData(partialResults, "uf"),
		YearDataByClimate:   mergeYearData(partialResults, "climate"),
		YearDataBySpeed:     mergeYearData(partialResults, "speed"),
		YearDataByTrack:     mergeYearData(partialResults, "track"),
		YearDataByPhaseDay:  mergeYearData(partialResults, "phase_day"),
		YearDataByMonth:     mergeYearData(partialResults, "month"),
		YearDataByDayWeek:   mergeYearData(partialResults, "day_week"),
		YearDataByShoulder:  mergeYearData(partialResults, "shoulder"),
		YearDataByGuardrail: mergeYearData(partialResults, "guardrail"),
		YearDataByMedian:    mergeYearData(partialResults, "median"),
		YearDataByHighway:   mergeYearData(partialResults, "highway"),
	}

	return result
}

func mergeYearData(partialResults sync.Map, key string) map[string]*YearData {
	result := make(map[string]*YearData)

	partialResults.Range(func(k, v interface{}) bool {
		if kStr, ok := k.(string); ok && strings.HasPrefix(kStr, key+"_") {
			yearKey := strings.TrimPrefix(kStr, key+"_")
			yearData := v.(*YearData)

			// Garantir que o mapa já exista
			if existing, ok := result[yearKey]; ok {
				existing.mu.Lock()
				for k, v := range yearData.TotalAcciden {
					existingData, exists := existing.TotalAcciden[k]
					if !exists {
						existing.TotalAcciden[k] = &AccidentData{}
					}

					existingData.TotalAccident += v.TotalAccident
					existingData.TotalDeath += v.TotalDeath
					existingData.TotalInvolved += v.TotalInvolved
					existingData.TotalInjured += v.TotalInjured
				}
				existing.mu.Unlock()
			} else {
				result[yearKey] = &YearData{
					TotalAcciden: deepCopy(yearData.TotalAcciden),
					mu:           sync.Mutex{},
				}
			}
		}
		return true
	})

	return result
}

func deepCopy(data map[string]*AccidentData) map[string]*AccidentData {
	copyMap := make(map[string]*AccidentData)
	for k, v := range data {
		copyMap[k] = &AccidentData{
			TotalAccident: v.TotalAccident,
			TotalDeath:    v.TotalDeath,
			TotalInvolved: v.TotalInvolved,
			TotalInjured:  v.TotalInjured,
		}
	}
	return copyMap
}

func getColumnIndex(columns []string, columnName string) int {
	for i, col := range columns {
		if strings.TrimSpace(col) == columnName {
			return i
		}
	}
	log.Fatalf("Coluna '%s' não encontrada no cabeçalho", columnName)
	return -1 // Isso só acontece se log.Fatal não parar a execução
}

func findColumnIndex(file *os.File, columnName string) int {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	if !scanner.Scan() {
		return -1 // Arquivo vazio ou erro ao ler
	}
	header := scanner.Text()
	columns := strings.Split(header, ";")

	for i, col := range columns {
		if strings.TrimSpace(col) == columnName {
			return i
		}
	}
	return -1 // Coluna não encontrada
}
