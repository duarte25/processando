package acidente

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Result struct {
	NumRows  int
	MediaRow float64
}

// ggfggh
func processRow(text string, columnIndex int) string {

	startPos := 0

	for i := 0; i < columnIndex; i++ {
		pos := strings.Index(text[startPos:], ";")

		if pos == -1 {
			return ""
		}

		startPos += pos + 1
	}

	endPos := strings.Index(text[startPos:], ";")
	if endPos == -1 {
		return text[startPos:]
	}
	return text[startPos : startPos+endPos]
}

func Acidente(file string) {

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	stateCounts := make(map[string]int)

	dateColumnIndex := 4
	ufAcidenteColumn := 3

	for scanner.Scan() {
		row := scanner.Text()

		date := processRow(row, dateColumnIndex)

		if strings.HasPrefix(date, "2022") {
			ufAcidente := processRow(row, ufAcidenteColumn)
			stateCounts[ufAcidente]++
		}

	}

	for state, count := range stateCounts {
		fmt.Printf("%s: %d\n", state, count)
	}
}

// num_acidente;chv_localidade;data_acidente;uf_acidente;ano_acidente;mes_acidente;mes_ano_acidente;codigo_ibge;
// dia_semana;fase_dia;tp_acidente;cond_meteorologica;end_acidente;num_end_acidente;cep_acidente;bairro_acidente;
// km_via_acidente;latitude_acidente;longitude_acidente;hora_acidente;tp_rodovia;cond_pista;tp_cruzamento;tp_pavimento;
// tp_curva;lim_velocidade;tp_pista;ind_guardrail;ind_cantcentral;ind_acostamento;qtde_acidente;qtde_acid_com_obitos
// ;qtde_envolvidos;qtde_feridosilesos;qtde_obitos
