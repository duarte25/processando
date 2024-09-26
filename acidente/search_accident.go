package accident

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// SearchAccident lê um arquivo CSV e conta quantas vezes o desireValue aparece na coluna "cep_acidente"
func SearchAccident(filePath, desireValue string) {
	// Abre o arquivo e insere no file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Erro ao abrir o arquivo:", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	// Lê o cabeçalho para encontrar o índice da coluna "cep_acidente"
	headerLine, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal("Erro ao ler o cabeçalho:", err)
	}

	// Divide o cabeçalho por vírgulas
	headerFields := strings.Split(strings.TrimSpace(headerLine), ";")

	// Encontra o índice da coluna "cep_acidente"
	idxColumn := -1
	for i, field := range headerFields {
		if field == "cep_acidente" {
			idxColumn = i
			break
		}
	}

	if idxColumn == -1 {
		log.Fatal("Coluna 'cep_acidente' não encontrada no arquivo")
	}

	// Variável para contar quantas vezes o desireValue aparece
	count := 0

	// Lê linha por linha do arquivo
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break // Fim do arquivo ou erro
		}

		// Divide a linha por vírgulas
		fields := strings.Split(strings.TrimSpace(line), ",")

		// Verifica se o índice está dentro do limite da linha
		fmt.Println(fields)
		if idxColumn < len(fields) {
			// Compara o valor na coluna "cep_acidente" com o valor desejado
			if fields[idxColumn] == desireValue {
				count++
			}
		}
	}

	// Exibe o resultado
	fmt.Printf("O CEP %s apareceu %d vezes no arquivo.\n", desireValue, count)
}
