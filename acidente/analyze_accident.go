package accident

import (
	"log"
	"os"
	"sync"
)

func AnalyzeAccidentData(filePath, indexToColumn, dateColumn, filterValue, indexFilterValue string) map[string]*YearData {
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
	amountInjuredColumn := findColumnIndex(file, "qtde_feridosilesos")

	filterColumn := findColumnIndex(file, indexFilterValue)

	if indexFilterValue != "" {
		filterColumn = findColumnIndex(file, indexFilterValue)
		if filterColumn == -1 {
			log.Fatal("Coluna de filtro não encontrada no cabeçalho")
		}
	} else {
		filterColumn = -1
	}

	if idxColumn == -1 || dateColumnIndex == -1 || amountDeathColumn == -1 || amountInvolvedColumn == -1 || amountInjuredColumn == -1 {
		log.Fatal("Coluna definida ou data_acidente não encontrada no cabeçalho")
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

	// Processar cada parte do arquivo em paralelo
	for i := 0; i < numParts; i++ {
		startOffset := int64(i) * partSize
		endOffset := startOffset + partSize
		if i == numParts-1 {
			endOffset = fileSize
		}

		wg.Add(1)
		go processFilePart(filePath, startOffset, endOffset, idxColumn, dateColumnIndex, amountDeathColumn, amountInvolvedColumn, amountInjuredColumn, filterColumn, filterValue, &wg, &counts)
	}

	// Aguardar todas as goroutines
	wg.Wait()

	// Copiar os dados do sync.Map para um novo map[string]*YearData
	result := make(map[string]*YearData)
	counts.Range(func(key, value interface{}) bool {
		year := key.(string)
		yearData := value.(*YearData)
		result[year] = yearData
		return true
	})

	// Exibir o mapa de dados copiados
	return result
}
