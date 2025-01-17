package accident

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

// Mapa para armazenar as contagens e somas
type AccidentData struct {
	TotalAccident int `json:"total_accident"`
	TotalDeath    int `json:"total_death"`
	TotalInvolved int `json:"total_involved"`
	TotalInjured  int `json:"total_injured"`
}

// Define a struct Year, que inclui a informação da coluna desejada
type YearData struct {
	mu           sync.Mutex
	TotalAcciden map[string]*AccidentData
}

func processFilePart(
	filePath string,
	startOffset, endOffset int64,
	idxColumn, dateColumnIndex, amountDeathColumn, amountInvolvedColumn, amountInjuredColumn, filterColumn int,
	filterValue string,
	wg *sync.WaitGroup,
	counts *sync.Map,
) {
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

		// Se o filtro estiver habilitado (não vazio), verificar a coluna
		if filterValue != "" {
			idx := 0
			var filterColumnValue string
			for i := 0; i <= filterColumn; i++ {
				filterColumnValue = nextColumn(line, &idx, ";")
			}

			if strings.TrimSpace(filterColumnValue) != filterValue {
				currentPos += int64(len(line))
				continue
			}
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
		var columnName string
		for i := 0; i <= idxColumn; i++ {
			columnName = nextColumn(line, &idx, ";")
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

		// Ler e somar `amountInjured`
		idx = 0
		var amountInjuredStr string
		for i := 0; i <= amountInjuredColumn; i++ {
			amountInjuredStr = nextColumn(line, &idx, ";")
		}
		amountInjured, _ := strconv.Atoi(amountInjuredStr)

		// Carregar ou inicializar o YearData para o ano atual
		yearDataInterface, _ := counts.LoadOrStore(date, &YearData{
			TotalAcciden: make(map[string]*AccidentData),
		})
		yearData := yearDataInterface.(*YearData)

		// Usar mutex para sincronizar o acesso ao mapa yearData
		yearData.mu.Lock()
		if _, exists := yearData.TotalAcciden[columnName]; !exists {
			yearData.TotalAcciden[columnName] = &AccidentData{}
		}
		yearData.TotalAcciden[columnName].TotalAccident++
		yearData.TotalAcciden[columnName].TotalDeath += amountDeath
		yearData.TotalAcciden[columnName].TotalInvolved += amountInvolved
		yearData.TotalAcciden[columnName].TotalInjured += amountInjured
		yearData.mu.Unlock()

		currentPos += int64(len(line))
	}
}
