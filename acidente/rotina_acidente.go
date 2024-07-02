package acidente

import (
	"fmt"
	"io/ioutil"
	"log"
	"runtime"
	"strings"
	"sync"
)

const numWorkers = 16  // Número de trabalhadores igual ao número de núcleos
const batchSize = 5000 // Tamanho do lote de linhas para cada trabalho

func extrairRow(text string, columnIndex int) string {
	startPos := 0
	for i := 0; i < columnIndex; i++ {
		pos := strings.IndexByte(text[startPos:], ';')
		if pos == -1 {
			return ""
		}
		startPos += pos + 1
	}
	endPos := strings.IndexByte(text[startPos:], ';')
	if endPos == -1 {
		return text[startPos:]
	}
	return text[startPos : startPos+endPos]
}

func worker(id int, jobs <-chan []string, results chan<- map[string]int, dateColumnIndex, ufAcidenteColumn int, wg *sync.WaitGroup) {
	defer wg.Done()
	localCounts := make(map[string]int)
	for batch := range jobs {
		for _, row := range batch {
			date := extrairRow(row, dateColumnIndex)
			if strings.HasPrefix(date, "2022") {
				ufAcidente := extrairRow(row, ufAcidenteColumn)
				localCounts[ufAcidente]++
			}
		}
	}
	results <- localCounts
}

func mergeResults(results <-chan map[string]int, finalCounts *sync.Map, wg *sync.WaitGroup) {
	defer wg.Done()
	for counts := range results {
		for state, count := range counts {
			if existing, ok := finalCounts.Load(state); ok {
				finalCounts.Store(state, existing.(int)+count)
			} else {
				finalCounts.Store(state, count)
			}
		}
	}
}

func ProcessoAcidente(file string) {

	// Configure GOMAXPROCS para usar todos os núcleos
	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	stateCounts := &sync.Map{}
	dateColumnIndex := 4
	ufAcidenteColumn := 3

	var wg sync.WaitGroup
	var mergeWG sync.WaitGroup
	jobs := make(chan []string, numWorkers)
	results := make(chan map[string]int, numWorkers)

	// Iniciar trabalhadores
	for w := 0; w < numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, dateColumnIndex, ufAcidenteColumn, &wg)
	}

	// Iniciar merge de resultados
	mergeWG.Add(1)
	go mergeResults(results, stateCounts, &mergeWG)

	// Enviar linhas para os trabalhadores em lotes
	var batch []string
	for _, line := range lines {
		batch = append(batch, line)
		if len(batch) >= batchSize {
			jobs <- batch
			batch = nil // Iniciar um novo lote
		}
	}
	if len(batch) > 0 {
		jobs <- batch
	}

	close(jobs)

	// Esperar todos os trabalhadores terminarem
	wg.Wait()
	close(results)

	// Esperar o merge dos resultados
	mergeWG.Wait()

	// Exibir resultados finais
	stateCounts.Range(func(key, value interface{}) bool {
		fmt.Printf("%s: %d\n", key, value.(int))
		return true
	})
}
