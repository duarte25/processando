package acidente

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

func Acidente(filePath, indexToColumn, dateColumn string) map[string]YearData {
	// Abre o arquivo e insere no file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Erro ao abrir o arquivo:", err)
	}
	defer file.Close()

	// Encontrar o índice da coluna desejada
	idxColumn := findColumnIndex(file, indexToColumn)
	dateColumnIndex := findColumnIndex(file, dateColumn)

	amountDeathColumn := findColumnIndex(file, "qtde_obitos")
	amountInvolvedColumn := findColumnIndex(file, "qtde_envolvidos")

	if idxColumn == -1 || dateColumnIndex == -1 || amountDeathColumn == -1 || amountInvolvedColumn == -1 {
		log.Fatal("Coluna uf_acidente ou data_acidente não encontrada no cabeçalho")
	}

	// Dividir o arquivo em várias partes
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal("Erro ao obter informações do arquivo:", err)
	}

	// Dividindo para conquistar
	fileSize := fileInfo.Size()
	partSize := fileSize / int64(numParts)

	var wg sync.WaitGroup
	var counts sync.Map
	var mu sync.Mutex

	// Processar cada parte do arquivo em paralelo
	for i := 0; i < numParts; i++ {
		startOffset := int64(i) * partSize
		endOffset := startOffset + partSize
		if i == numParts-1 {
			endOffset = fileSize
		}

		wg.Add(1)
		go processFilePart(filePath, startOffset, endOffset, idxColumn, dateColumnIndex, amountDeathColumn, amountInvolvedColumn, &wg, &counts, &mu)
	}

	// Aguardar todas as goroutines
	wg.Wait()

	// Copiar os dados do sync.Map para um novo map[string]YearData
	result := make(map[string]YearData)
	counts.Range(func(key, value interface{}) bool {
		year := key.(string)
		yearData := value.(*YearData)
		result[year] = *yearData
		return true
	})

	// Exibir o mapa de dados copiados
	return result
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
