package main

import (
	"fmt"
	"processando/acidente"
	"time"
)

func main() {
	start := time.Now()

	acidente.Acidente("Acidentes_DadosAbertos_20230412.csv")

	end := time.Now() // Captura o tempo de término

	duration := end.Sub(start) // Calcula a diferença de tempo

	fmt.Printf("Tempo gasto no precesso sequencial: %v\n", duration)
}
