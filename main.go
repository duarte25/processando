package main

import (
	"fmt"
	"processando/acidente"
	"time"
)

func main() {
	// Contagem de todo o processo
	start := time.Now()

	result := acidente.Acidente("./Acidentes_DadosAbertos_20230412.csv")

	// Conte quanto tempo durou o processo
	end := time.Now()
	duration := end.Sub(start)

	// Encontrar o estado com o maior número de ocorrências
	var maxUf string
	var maxCount int
	for uf, count := range result {
		if count > maxCount {
			maxUf = uf
			maxCount = count
		}
	}

	// Imprimir o resultado
	fmt.Printf("Estado com maior número de ocorrências: %s (%d ocorrências)\n", maxUf, maxCount)
	fmt.Println("Resultado: ", result)

	fmt.Printf("Tempo gasto: %v\n", duration)
}
