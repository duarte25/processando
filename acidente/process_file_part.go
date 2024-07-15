package acidente

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
	UFs map[string]*UFData
}

func processFilePart(filePath, year string, startOffset, endOffset int64, idxColumn, dateColumnIndex, amountDeathColumn, amountInvolvedColumn int, wg *sync.WaitGroup, counts *sync.Map) {
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

	groupYear := make(map[string]*YearData)

	localCounts := make(map[string]*UFData)
	currentPos := startOffset
	for currentPos < endOffset {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		idx := 0
		var date string
		for i := 0; i <= dateColumnIndex; i++ {
			date = nextColumn(line, &idx, ";")
		}

		if !strings.HasPrefix(date, year) {
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

		// Atualizar os dados no mapa localCounts
		if _, exists := localCounts[uf]; !exists {
			localCounts[uf] = &UFData{}
		}
		localCounts[uf].Count++
		localCounts[uf].TotalDeath += amountDeath
		localCounts[uf].TotalInvolved += amountInvolved

		// Verifica se o ano já existe no mapa groupYear
		if _, exists := groupYear[year]; !exists {
			groupYear[year] = &YearData{
				UFs: make(map[string]*UFData),
			}
		}

		// Atualiza os dados no mapa groupYear
		groupYear[year].UFs[uf] = localCounts[uf]

		currentPos += int64(len(line))
	}

	// Percorrendo localCounts e acumulando em counts
	for unit, data := range localCounts {
		actual, _ := counts.LoadOrStore(unit, data)
		if actual != data {
			storedData := actual.(*UFData)
			storedData.Count += data.Count
			storedData.TotalDeath += data.TotalDeath
			storedData.TotalInvolved += data.TotalInvolved
			counts.Store(unit, storedData)
		}
	}

	// Exibe o conteúdo do groupYear (opcional)
	for year, yearData := range groupYear {
		fmt.Printf("Year: %s\n", year)
		for uf, data := range yearData.UFs {
			fmt.Printf("UF: %s, Count: %d, TotalDeath: %d, TotalInvolved: %d\n",
				uf, data.Count, data.TotalDeath, data.TotalInvolved)
		}
	}
}
