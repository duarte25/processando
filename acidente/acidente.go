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

func Acidente(filePath, indexToColumn, dateColumn, year string) map[string]int {
	// Abre o arquivo e insere no file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Erro ao abrir o arquivo:", err)
	}
	defer file.Close()

	// Encontrar o índice da coluna desejada
	idxColumn := findColumnIndex(file, indexToColumn)
	dateColumnIndex := findColumnIndex(file, dateColumn)
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
		go processFilePart(filePath, year, startOffset, endOffset, idxColumn, dateColumnIndex, &wg, &counts)
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

/*
num_acidente;chv_localidade;data_acidente;uf_acidente;ano_acidente;mes_acidente;mes_ano_acidente;codigo_ibge;
dia_semana;fase_dia;tp_acidente;cond_meteorologica;end_acidente;num_end_acidente;cep_acidente;bairro_acidente;
km_via_acidente;latitude_acidente;longitude_acidente;hora_acidente;tp_rodovia;cond_pista;tp_cruzamento;
tp_pavimento;tp_curva;lim_velocidade;tp_pista;ind_guardrail;ind_cantcentral;ind_acostamento;qtde_acidente;
qtde_acid_com_obitos;qtde_envolvidos;qtde_feridosilesos;qtde_obitos
*/

// uf_acidente, ano_acidente,
