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
	count         int
	totalDeath    int
	totalInvolved int
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
		// Inclui dessa forma pois é a ultima coluna e então para não acontecer erros na hora de calcular
		// Erro na quebra de linha inclui esse trimspace mas somente nele acredito que os outros não quebrarão
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

		// Atualizar os dados no mapa
		if _, exists := localCounts[uf]; !exists {
			localCounts[uf] = &UFData{}
		}
		localCounts[uf].count++
		localCounts[uf].totalDeath += amountDeath
		localCounts[uf].totalInvolved += amountInvolved

		currentPos += int64(len(line))
	}

	/*
		percorrendo um mapa (localCounts) que contém contagens de acidentes por unidade
		e acumulando essas contagens em uma estrutura de dados chamada
		counts, que é um sync.Map. Vamos detalhar cada linha do código
	*/
	for unit, data := range localCounts {
		actual, _ := counts.LoadOrStore(unit, data)
		if actual != data {
			storedData := actual.(*UFData)
			storedData.count += data.count
			storedData.totalDeath += data.totalDeath
			storedData.totalInvolved += data.totalInvolved
			counts.Store(unit, storedData)
		}
	}
}
