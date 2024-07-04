package main

import (
	"fmt"
	"processando/acidente"
	"time"
)

func main() {
	// Contagem de todo o processo
	start := time.Now()

	// Passarei sempre date aqui mas estou pensando em injetar lá no algoritmo pois provavelmente sempre usarei ele.. Vamos ve
	result := acidente.Acidente("./Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "data_acidente", "2022")

	// Conte quanto tempo durou o processo
	end := time.Now()
	duration := end.Sub(start)

	for uf, count := range result {
		fmt.Printf("%s: %d\n", uf, count)
	}

	fmt.Printf("Tempo gasto: %v\n", duration)
}
