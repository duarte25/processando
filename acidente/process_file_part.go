package accident

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Mapa para armazenar as contagens e somas por UF
type UFData struct {
	Count         int `json:"count"`
	TotalDeath    int `json:"total_death"`
	TotalInvolved int `json:"total_involved"`
}

// Define a struct Year, que inclui UFData
type YearData struct {
	mu  sync.Mutex
	UFs map[string]*UFData
}

func processFilePart(filePath string, startOffset, endOffset int64, idxColumn, dateColumnIndex, amountDeathColumn, amountInvolvedColumn int, wg *sync.WaitGroup, counts *sync.Map) {
	defer wg.Done()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	file.Seek(startOffset, 0)
	reader := bufio.NewReader(file)

	// Ajustar o início para o início da próxima linha
	if startOffset > 0 {
		_, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("Erro ao ajustar o início da linha:", err)
			return
		}
	}

	currentPos := startOffset
	for currentPos < endOffset {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Erro ao ler a linha:", err)
			break
		}

		idx := 0
		var date string
		for i := 0; i <= dateColumnIndex; i++ {
			date = nextColumn(line, &idx, ";")
		}

		if !strings.HasPrefix(date, "20") {
			currentPos += int64(len(line))
			continue
		}

		idx = 0
		var uf string
		for i := 0; i <= idxColumn; i++ {
			uf = nextColumn(line, &idx, ";")
		}

		// Ler e somar `amountDeath`
		idx = 0
		var amountDeathStr string
		for i := 0; i <= amountDeathColumn; i++ {
			amountDeathStr = nextColumn(line, &idx, ";")
		}
		amountDeathStr = strings.TrimSpace(amountDeathStr)
		amountDeath, err := strconv.Atoi(amountDeathStr)
		if err != nil {
			fmt.Println("Erro ao converter amountDeath:", err)
			amountDeath = 0
		}

		// Ler e somar `amountInvolved`
		idx = 0
		var amountInvolvedStr string
		for i := 0; i <= amountInvolvedColumn; i++ {
			amountInvolvedStr = nextColumn(line, &idx, ";")
		}
		amountInvolved, _ := strconv.Atoi(amountInvolvedStr)

		// fmt.Println(date)
		// year := date[:4] // Assume que o ano está nos primeiros 4 caracteres da data

		// Carregar ou inicializar o YearData para o ano atual
		yearDataInterface, _ := counts.LoadOrStore(date, &YearData{
			UFs: make(map[string]*UFData),
		})
		yearData := yearDataInterface.(*YearData)

		// Usar mutex para sincronizar o acesso ao mapa yearData
		yearData.mu.Lock()
		if _, exists := yearData.UFs[uf]; !exists {
			yearData.UFs[uf] = &UFData{}
		}
		yearData.UFs[uf].Count++
		yearData.UFs[uf].TotalDeath += amountDeath
		yearData.UFs[uf].TotalInvolved += amountInvolved
		yearData.mu.Unlock()

		currentPos += int64(len(line))
	}
}
