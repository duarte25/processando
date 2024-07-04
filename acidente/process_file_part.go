package acidente

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

func processFilePart(filePath, year string, startOffset, endOffset int64, idxColumn, dateColumnIndex int, wg *sync.WaitGroup, counts *sync.Map) {
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

	localCounts := make(map[string]int)
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

		localCounts[uf]++
		currentPos += int64(len(line))
	}

	/*
		percorrendo um mapa (localCounts) que contém contagens de acidentes por unidade
		e acumulando essas contagens em uma estrutura de dados chamada
		counts, que é um sync.Map. Vamos detalhar cada linha do código
	*/
	for unit, count := range localCounts {
		actual, _ := counts.LoadOrStore(unit, count)
		if actual != count {
			counts.Store(unit, actual.(int)+count)
		}
	}
}
