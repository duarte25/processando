package main

import (
	"fmt"
	"processando/acidente"
	"time"
)

func main() {
	// Contagem de todo o processo
	start := time.Now()

	// Chamar a função para processar os acidentes
	result := acidente.Acidente("./Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "data_acidente", "2022")

	// Conte quanto tempo durou o processo
	end := time.Now()
	duration := end.Sub(start)

	fmt.Println("Resultado: ", result)

	fmt.Printf("Tempo gasto: %v\n", duration)
}
