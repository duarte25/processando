package acidente

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

const (
	maxGoroutines = 100 // Limite de goroutines simultâneas
	numParts      = 8   // Número de partes para dividir o arquivo
)

func Acidente(filePath string) map[string]int {
	// Abre o arquivo e insere no file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Erro ao abrir o arquivo:", err)
	}
	defer file.Close()

	// Encontrar o índice da coluna desejada
	idxColumn := findColumnIndex(file, "uf_acidente")
	dateColumnIndex := findColumnIndex(file, "data_acidente")
	if idxColumn == -1 || dateColumnIndex == -1 {
		log.Fatal("Coluna uf_acidente ou data_acidente não encontrada no cabeçalho")
	}

	// Dividir o arquivo em várias partes
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal("Erro ao obter informações do arquivo:", err)
	}

	// Dividindo para conquistar
	fileSize := fileInfo.Size()
	partSize := fileSize / numParts

	var wg sync.WaitGroup
	var counts sync.Map

	// Processar cada parte do arquivo em paralelo
	for i := 0; i < numParts; i++ {
		startOffset := int64(i) * partSize
		endOffset := startOffset + partSize
		if i == numParts-1 {
			endOffset = fileSize
		}

		wg.Add(1)
		go processFilePart(filePath, startOffset, endOffset, idxColumn, dateColumnIndex, &wg, &counts)
	}

	// Aguardar todas as goroutines
	wg.Wait()

	// Coletar resultados do sync.Map
	result := make(map[string]int)
	counts.Range(func(key, value interface{}) bool {
		result[key.(string)] = value.(int)
		return true
	})

	return result
}

func processFilePart(filePath string, startOffset, endOffset int64, idxColumn, dateColumnIndex int, wg *sync.WaitGroup, counts *sync.Map) {
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

		// Filtrar para alguma coluna por exemplo data
		dateStartPos := 0
		for i := 0; i < dateColumnIndex; i++ {
			pos := strings.Index(line[dateStartPos:], ";")
			if pos == -1 {
				break
			}
			dateStartPos += pos + 1
		}

		dateEndPos := strings.Index(line[dateStartPos:], ";")
		if dateEndPos == -1 {
			dateEndPos = len(line)
		} else {
			dateEndPos += dateStartPos
		}

		date := line[dateStartPos:dateEndPos]
		if !strings.HasPrefix(date, "2022") {
			currentPos += int64(len(line))
			continue
		}

		// Processar uf_acidente
		startPos := 0
		for i := 0; i < idxColumn; i++ {
			pos := strings.Index(line[startPos:], ";")
			if pos == -1 {
				break
			}
			startPos += pos + 1
		}

		endPos := strings.Index(line[startPos:], ";")
		if endPos == -1 {
			endPos = len(line)
		} else {
			endPos += startPos
		}

		uf := line[startPos:endPos]
		localCounts[uf]++

		currentPos += int64(len(line))
	}

	for uf, count := range localCounts {
		actual, _ := counts.LoadOrStore(uf, count)
		if actual != count {
			counts.Store(uf, actual.(int)+count)
		}
	}
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
