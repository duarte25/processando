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
	acidente.Acidente("./Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "data_acidente", "2022")

	// Conte quanto tempo durou o processo
	end := time.Now()
	duration := end.Sub(start)

	// Exibir os resultados
	// for uf, data := range result {
	// 	fmt.Printf("%s: count=%d, totalDeath=%d, totalInvolved=%d\n", uf, data.count, data.totalDeath, data.totalInvolved)
	// }

	fmt.Printf("Tempo gasto: %v\n", duration)
}
