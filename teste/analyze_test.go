package accident

import (
	"log"
	accident "processando/acidente"
	"processando/src/configs"
	"testing"
)

func TestAnalyzeAccidentData(t *testing.T) {
	// Carregar as configurações
	err := configs.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar configurações: %v", err)
	}

	// Executa a função que queremos testar
	results := accident.AnalyzeAccidentData("../Acidentes_DadosAbertos_20230412.csv", "uf_acidente", "ano_acidente", "", "")

	// Verifica os resultados esperados para o ano de 2022
	expectedData2022 := map[string]*accident.AccidentData{
		"AC": {
			TotalAccident: 3775,
			TotalDeath:    47,
			TotalInvolved: 7725,
		},
		"AL": {
			TotalAccident: 2568,
			TotalDeath:    642,
			TotalInvolved: 5181,
		},
		"AM": {
			TotalAccident: 1,
			TotalDeath:    1,
			TotalInvolved: 1,
		},
		"MG": {
			TotalAccident: 256298,
			TotalDeath:    1993,
			TotalInvolved: 483857,
		},
		"SP": {
			TotalAccident: 121838,
			TotalDeath:    3475,
			TotalInvolved: 67653,
		},
		"RJ": {
			TotalAccident: 15299,
			TotalDeath:    1361,
			TotalInvolved: 15262,
		},
	}

	resultData2022, ok := results["2022"]
	if !ok {
		t.Fatalf("Resultados para o ano de 2022 não encontrados")
	}

	for uf, expectedUFData := range expectedData2022 {

		resultUFData, ok := resultData2022.TotalAcciden[uf]
		if !ok {
			t.Fatalf("Resultados para o estado %s não encontrados no ano de 2022", uf)
		}
		if resultUFData.TotalAccident != expectedUFData.TotalAccident {
			t.Errorf("Ano 2022, UF %s: esperava %d acidentes, obteve %d", uf, expectedUFData.TotalAccident, resultUFData.TotalAccident)
		}
		if resultUFData.TotalDeath != expectedUFData.TotalDeath {
			t.Errorf("Ano 2022, UF %s: esperava %d óbitos, obteve %d", uf, expectedUFData.TotalDeath, resultUFData.TotalDeath)
		}
		if resultUFData.TotalInvolved != expectedUFData.TotalInvolved {
			t.Errorf("Ano 2022, UF %s: esperava %d envolvidos, obteve %d", uf, expectedUFData.TotalInvolved, resultUFData.TotalInvolved)
		}
	}
}
