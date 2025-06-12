package sequencial

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Totais struct {
	QtdeAcidente      int
	QtdeEnvolvidos    int
	QtdeFeridosIlesos int
	QtdeObitos        int
}

type AcidenteData struct {
	Geral             Totais
	UF                map[string]*Totais
	CondMeteorologica map[string]*Totais
	DiaSemana         map[string]*Totais
	IndGuardrail      map[string]*Totais
	IndCantCentral    map[string]*Totais
	MesAcidente       map[string]*Totais
	FaseDia           map[string]*Totais
	IndAcostamento    map[string]*Totais
	LimVelocidade     map[string]*Totais
	CondPista         map[string]*Totais
}

func newAcidenteData() *AcidenteData {
	return &AcidenteData{
		UF:                make(map[string]*Totais),
		CondMeteorologica: make(map[string]*Totais),
		DiaSemana:         make(map[string]*Totais),
		IndGuardrail:      make(map[string]*Totais),
		IndCantCentral:    make(map[string]*Totais),
		MesAcidente:       make(map[string]*Totais),
		FaseDia:           make(map[string]*Totais),
		IndAcostamento:    make(map[string]*Totais),
		LimVelocidade:     make(map[string]*Totais),
		CondPista:         make(map[string]*Totais),
	}
}

func getOrCreate(m map[string]*Totais, key string) *Totais {
	if key == "" {
		key = "NÃO INFORMADO"
	}
	if _, ok := m[key]; !ok {
		m[key] = &Totais{}
	}
	return m[key]
}

func processRow(text string, columnIndex int, minColumns int) string {
	cols := strings.Split(text, ";")
	if len(cols) < minColumns {
		log.Printf("Linha com colunas insuficientes: %s", text)
		return ""
	}
	if columnIndex >= len(cols) {
		return ""
	}
	return strings.TrimSpace(cols[columnIndex])
}

func parseInt(value string) int {
	n, err := strconv.Atoi(strings.TrimSpace(value))
	if err != nil {
		return 0
	}
	return n
}

func Acidente(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	acidentesPorAno := make(map[string]*AcidenteData)

	// Ler cabeçalho e mapear colunas dinamicamente
	if !scanner.Scan() {
		log.Fatal("Arquivo vazio ou falha ao ler cabeçalho")
	}
	headers := strings.Split(scanner.Text(), ";")
	headerMap := make(map[string]int)
	for i, h := range headers {
		headerMap[strings.TrimSpace(h)] = i
	}

	// Definir índices com base no cabeçalho
	// dateIndex, _ := headerMap["data_acidente"]
	anoIndex, _ := headerMap["ano_acidente"]
	mesIndex, _ := headerMap["mes_acidente"]
	diaSemanaIndex, _ := headerMap["dia_semana"]
	faseDiaIndex, _ := headerMap["fase_dia"]
	condMeteoIndex, _ := headerMap["cond_meteorologica"]
	ufIndex, _ := headerMap["uf_acidente"]
	condPistaIndex, _ := headerMap["cond_pista"]
	indGuardrailIndex, _ := headerMap["ind_guardrail"]
	indCantCentralIndex, _ := headerMap["ind_cantcentral"]
	indAcostamentoIndex, _ := headerMap["ind_acostamento"]
	limVelocidadeIndex, _ := headerMap["lim_velocidade"]
	qtdeEnvolvidosIndex, _ := headerMap["qtde_envolvidos"]
	qtdeFeridosIlesosIndex, _ := headerMap["qtde_feridosilesos"]
	qtdeObitosIndex, _ := headerMap["qtde_obitos"]

	// Loop pelas linhas restantes
	for scanner.Scan() {
		row := scanner.Text()

		// Extrair ano
		year := processRow(row, anoIndex, 10)
		if year == "" {
			continue
		}

		if _, ok := acidentesPorAno[year]; !ok {
			acidentesPorAno[year] = newAcidenteData()
		}
		data := acidentesPorAno[year]

		qtdEnvol := parseInt(processRow(row, qtdeEnvolvidosIndex, 35))
		qtdFeridos := parseInt(processRow(row, qtdeFeridosIlesosIndex, 35))
		qtdObitos := parseInt(processRow(row, qtdeObitosIndex, 35))

		// Atualizar totais gerais
		data.Geral.QtdeAcidente++
		data.Geral.QtdeEnvolvidos += qtdEnvol
		data.Geral.QtdeFeridosIlesos += qtdFeridos
		data.Geral.QtdeObitos += qtdObitos

		attrList := []struct {
			m   map[string]*Totais
			idx int
		}{
			{data.UF, ufIndex},
			{data.CondMeteorologica, condMeteoIndex},
			{data.DiaSemana, diaSemanaIndex},
			{data.FaseDia, faseDiaIndex},
			{data.CondPista, condPistaIndex},
			{data.IndGuardrail, indGuardrailIndex},
			{data.IndCantCentral, indCantCentralIndex},
			{data.IndAcostamento, indAcostamentoIndex},
			{data.LimVelocidade, limVelocidadeIndex},
			{data.MesAcidente, mesIndex},
		}

		for _, attr := range attrList {
			valor := processRow(row, attr.idx, 35)
			t := getOrCreate(attr.m, valor)
			t.QtdeAcidente++
			t.QtdeEnvolvidos += qtdEnvol
			t.QtdeFeridosIlesos += qtdFeridos
			t.QtdeObitos += qtdObitos
		}
	}

	// Exibir os dados
	for year, data := range acidentesPorAno {
		fmt.Printf("\nAno: %s\n", year)
		fmt.Printf("Geral:\n")
		fmt.Printf("  Acidentes: %d\n", data.Geral.QtdeAcidente)
		fmt.Printf("  Envolvidos: %d\n", data.Geral.QtdeEnvolvidos)
		fmt.Printf("  Feridos ilesos: %d\n", data.Geral.QtdeFeridosIlesos)
		fmt.Printf("  Óbitos: %d\n", data.Geral.QtdeObitos)

		printMap("UF", data.UF)
		printMap("Condições Meteorológicas", data.CondMeteorologica)
		printMap("Dia da Semana", data.DiaSemana)
		printMap("Fase do Dia", data.FaseDia)
		printMap("Condição da Pista", data.CondPista)
		printMap("Guarda-corpo", data.IndGuardrail)
		printMap("Canteiro Central", data.IndCantCentral)
		printMap("Acostamento", data.IndAcostamento)
		printMap("Limite de Velocidade", data.LimVelocidade)
		printMap("Mês do Acidente", data.MesAcidente)
	}
}

func printMap(title string, m map[string]*Totais) {
	fmt.Printf("\n%s:\n", title)
	for k, v := range m {
		fmt.Printf("  %s: Acidentes=%d, Envolvidos=%d, Feridos=%d, Óbitos=%d\n",
			k, v.QtdeAcidente, v.QtdeEnvolvidos, v.QtdeFeridosIlesos, v.QtdeObitos)
	}
}
