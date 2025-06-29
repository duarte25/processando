package accident

import (
	"os"
	"testing"
)

func TestAccident(t *testing.T) {
	// Carregar as configurações
	t.Setenv("ACIDENTE_FILE", "../acidentes.csv") // Caso bug no mod inclua o path do arquivo

	file := os.Getenv("ACIDENTE_FILE")

	// Executa a nova função AnalyzeAll
	results := AnalyzeAll(file)

	t.Run("TestTrackConditionData", func(t *testing.T) {
		expectedTrackData2022 := map[string]*AccidentData{
			"SECA": {
				TotalAccident: 209975,
				TotalDeath:    4507,
				TotalInvolved: 338186,
				TotalInjured:  333679,
			},
			"COM LAMA": {
				TotalAccident: 168,
				TotalDeath:    2,
				TotalInvolved: 252,
				TotalInjured:  250,
			},
			"ESCORREGADIA": {
				TotalAccident: 717,
				TotalDeath:    4,
				TotalInvolved: 842,
				TotalInjured:  838,
			},
			"MOLHADA": {
				TotalAccident: 25644,
				TotalDeath:    386,
				TotalInvolved: 42153,
				TotalInjured:  41767,
			},
			"NAO INFORMADO": {
				TotalAccident: 655243,
				TotalDeath:    11570,
				TotalInvolved: 1077137,
				TotalInjured:  1065567,
			},
			"OBSTRUIDA": {
				TotalAccident: 241,
				TotalDeath:    8,
				TotalInvolved: 229,
				TotalInjured:  221,
			},
			"COM MATERIAL GRANULADO": {
				TotalAccident: 65,
				TotalDeath:    3,
				TotalInvolved: 127,
				TotalInjured:  124,
			},
			"DESCONHECIDO": {
				TotalAccident: 89861,
				TotalDeath:    3726,
				TotalInvolved: 42625,
				TotalInjured:  38899,
			},
			"COM BURACO": {
				TotalAccident: 243,
				TotalDeath:    24,
				TotalInvolved: 429,
				TotalInjured:  405,
			},
		}

		// Obter os resultados específicos para "cond_pista" e o ano de 2022
		yearData2022, ok := results.YearDataByTrack["2022"]
		if !ok {
			t.Fatalf("Resultados para 'cond_pista' no ano de 2022 não encontrados")
		}

		validateAccidentData(t, "condição de pista", yearData2022.TotalAcciden, expectedTrackData2022)
	})

	t.Run("TestUFData", func(t *testing.T) {
		expectedUFData2022 := map[string]*AccidentData{
			"PE": {
				TotalAccident: 9325,
				TotalDeath:    1086,
				TotalInvolved: 10425,
				TotalInjured:  9339,
			},
			"RS": {
				TotalAccident: 1526,
				TotalDeath:    1703,
				TotalInvolved: 3649,
				TotalInjured:  1946,
			},
			"MA": {
				TotalAccident: 6218,
				TotalDeath:    1073,
				TotalInvolved: 6825,
				TotalInjured:  5752,
			},
			"BA": {
				TotalAccident: 3656,
				TotalDeath:    340,
				TotalInvolved: 8310,
				TotalInjured:  7970,
			},
			"AL": {
				TotalAccident: 2568,
				TotalDeath:    642,
				TotalInvolved: 5181,
				TotalInjured:  4539,
			},
			"DF": {
				TotalAccident: 55518,
				TotalDeath:    289,
				TotalInvolved: 106844,
				TotalInjured:  106555,
			},
			"MG": {
				TotalAccident: 256403,
				TotalDeath:    1995,
				TotalInvolved: 484019,
				TotalInjured:  482024,
			},
			"RO": {
				TotalAccident: 16574,
				TotalDeath:    404,
				TotalInvolved: 31805,
				TotalInjured:  31401,
			},
			"AM": {
				TotalAccident: 1,
				TotalDeath:    1,
				TotalInvolved: 1,
				TotalInjured:  0,
			},
			"MS": {
				TotalAccident: 21168,
				TotalDeath:    142,
				TotalInvolved: 30554,
				TotalInjured:  30412,
			},
			"AP": {
				TotalAccident: 864,
				TotalDeath:    2,
				TotalInvolved: 1147,
				TotalInjured:  1145,
			},
			"TO": {
				TotalAccident: 3317,
				TotalDeath:    193,
				TotalInvolved: 6404,
				TotalInjured:  6211,
			},
			"RJ": {
				TotalAccident: 16902,
				TotalDeath:    1524,
				TotalInvolved: 16865,
				TotalInjured:  15341,
			},
			"CE": {
				TotalAccident: 9582,
				TotalDeath:    622,
				TotalInvolved: 19506,
				TotalInjured:  18884,
			},
			"SC": {
				TotalAccident: 168604,
				TotalDeath:    925,
				TotalInvolved: 233786,
				TotalInjured:  232861,
			},
			"AC": {
				TotalAccident: 3775,
				TotalDeath:    47,
				TotalInvolved: 7725,
				TotalInjured:  7678,
			},
		}

		yearData2022, ok := results.YearDataByUF["2022"]
		if !ok {
			t.Fatalf("Resultados para 'uf_acidente' no ano de 2022 não encontrados")
		}

		validateAccidentData(t, "UF", yearData2022.TotalAcciden, expectedUFData2022)
	})

	t.Run("TestSpeedData", func(t *testing.T) {
		expectedSpeedData2022 := map[string]*AccidentData{
			"40 KMH": {
				TotalAccident: 10131,
				TotalDeath:    133,
				TotalInvolved: 24844,
				TotalInjured:  24711,
			},
			"NAO INFORMADO": {
				TotalAccident: 958374,
				TotalDeath:    19106,
				TotalInvolved: 1443425,
				TotalInjured:  1424319,
			},
			"80 KMH": {
				TotalAccident: 3646,
				TotalDeath:    541,
				TotalInvolved: 9555,
				TotalInjured:  9014,
			},
			"60 KMH": {
				TotalAccident: 6564,
				TotalDeath:    311,
				TotalInvolved: 16135,
				TotalInjured:  15824,
			},
			"110 KMH": {
				TotalAccident: 813,
				TotalDeath:    111,
				TotalInvolved: 1968,
				TotalInjured:  1857,
			},
			"30 KMH": {
				TotalAccident: 2629,
				TotalDeath:    28,
				TotalInvolved: 6053,
				TotalInjured:  6025,
			},
		}

		yearData2022, ok := results.YearDataBySpeed["2022"]
		if !ok {
			t.Fatalf("Resultados para 'lim_velocidade' no ano de 2022 não encontrados")
		}

		validateAccidentData(t, "velocidade", yearData2022.TotalAcciden, expectedSpeedData2022)
	})

	t.Run("TestShoulderData", func(t *testing.T) {
		expectedShoulderData2022 := map[string]*AccidentData{
			"NAO": {
				TotalAccident: 66117,
				TotalDeath:    2470,
				TotalInvolved: 156999,
				TotalInjured:  154529,
			},
			"NAO INFORMADO": {
				TotalAccident: 631919,
				TotalDeath:    11336,
				TotalInvolved: 952826,
				TotalInjured:  941490,
			},
			"DESCONHECIDO": {
				TotalAccident: 203138,
				TotalDeath:    4027,
				TotalInvolved: 231179,
				TotalInjured:  227152,
			},
			"SIM": {
				TotalAccident: 80983,
				TotalDeath:    2397,
				TotalInvolved: 160976,
				TotalInjured:  158579,
			},
		}

		yearData2022, ok := results.YearDataByShoulder["2022"]
		if !ok {
			t.Fatalf("Resultados para 'ind_acostamento' no ano de 2022 não encontrados")
		}

		validateAccidentData(t, "acostamento", yearData2022.TotalAcciden, expectedShoulderData2022)
	})

	t.Run("TestPhaseDayData", func(t *testing.T) {
		expectedPhaseDayData2022 := map[string]*AccidentData{
			"TARDE": {
				TotalAccident: 363622,
				TotalDeath:    4575,
				TotalInvolved: 569266,
				TotalInjured:  564691,
			},
			"NAO INFORMADO": {
				TotalAccident: 9720,
				TotalDeath:    1183,
				TotalInvolved: 11007,
				TotalInjured:  9824,
			},
			"NOITE": {
				TotalAccident: 261627,
				TotalDeath:    6573,
				TotalInvolved: 396173,
				TotalInjured:  389600,
			},
			"DESCONHECIDO": {
				TotalAccident: 471,
				TotalDeath:    357,
				TotalInvolved: 463,
				TotalInjured:  106,
			},
			"MANHA": {
				TotalAccident: 280169,
				TotalDeath:    3966,
				TotalInvolved: 436045,
				TotalInjured:  432079,
			},
			"MADRUGADA": {
				TotalAccident: 66548,
				TotalDeath:    3576,
				TotalInvolved: 89026,
				TotalInjured:  85450,
			},
		}

		yearData2022, ok := results.YearDataByPhaseDay["2022"]
		if !ok {
			t.Fatalf("Resultados para 'fase_dia' no ano de 2022 não encontrados")
		}

		validateAccidentData(t, "fase_dia", yearData2022.TotalAcciden, expectedPhaseDayData2022)
	})

	t.Run("TestMonthData", func(t *testing.T) {
		expectedMonthData2022 := map[string]*AccidentData{
			"04": {
				TotalAccident: 83866,
				TotalDeath:    1760,
				TotalInvolved: 124631,
				TotalInjured:  122871,
			},
			"03": {
				TotalAccident: 88931,
				TotalDeath:    1769,
				TotalInvolved: 148059,
				TotalInjured:  146290,
			},
			"12": {
				TotalAccident: 57809,
				TotalDeath:    1029,
				TotalInvolved: 88954,
				TotalInjured:  87925,
			},
			"11": {
				TotalAccident: 75935,
				TotalDeath:    1169,
				TotalInvolved: 109175,
				TotalInjured:  108006,
			},
			"02": {
				TotalAccident: 79153,
				TotalDeath:    1717,
				TotalInvolved: 130229,
				TotalInjured:  128512,
			},
			"09": {
				TotalAccident: 87821,
				TotalDeath:    1751,
				TotalInvolved: 136429,
				TotalInjured:  134678,
			},
			"08": {
				TotalAccident: 89190,
				TotalDeath:    1761,
				TotalInvolved: 119317,
				TotalInjured:  117556,
			},
			"05": {
				TotalAccident: 90243,
				TotalDeath:    1863,
				TotalInvolved: 133207,
				TotalInjured:  131344,
			},
			"06": {
				TotalAccident: 84442,
				TotalDeath:    1855,
				TotalInvolved: 124413,
				TotalInjured:  122558,
			},
			"01": {
				TotalAccident: 74668,
				TotalDeath:    1604,
				TotalInvolved: 120469,
				TotalInjured:  118865,
			},
			"07": {
				TotalAccident: 86988,
				TotalDeath:    1951,
				TotalInvolved: 144423,
				TotalInjured:  142472,
			},
			"10": {
				TotalAccident: 83111,
				TotalDeath:    2001,
				TotalInvolved: 122674,
				TotalInjured:  120673,
			},
		}

		yearData2022, ok := results.YearDataByMonth["2022"]
		if !ok {
			t.Fatalf("Resultados para 'mes' no ano de 2022 não encontrados")
		}

		validateAccidentData(t, "mes", yearData2022.TotalAcciden, expectedMonthData2022)
	})

	t.Run("TestSidewalkData", func(t *testing.T) {
		expectedSidewalkData2022 := map[string]*AccidentData{
			"NAO": {
				TotalAccident: 68022,
				TotalDeath:    2620,
				TotalInvolved: 168675,
				TotalInjured:  166055,
			},
			"NAO INFORMADO": {
				TotalAccident: 796087,
				TotalDeath:    16142,
				TotalInvolved: 1094374,
				TotalInjured:  1078232,
			},
			"DESCONHECIDO": {
				TotalAccident: 67574,
				TotalDeath:    899,
				TotalInvolved: 149364,
				TotalInjured:  148465,
			},
			"SIM": {
				TotalAccident: 50474,
				TotalDeath:    569,
				TotalInvolved: 89567,
				TotalInjured:  88998,
			},
		}

		yearData2022, ok := results.YearDataByMedian["2022"]
		if !ok {
			t.Fatalf("Resultados para 'ind_cantcentral' no ano de 2022 não encontrados")
		}

		validateAccidentData(t, "ind_cantcentral", yearData2022.TotalAcciden, expectedSidewalkData2022)
	})

	t.Run("TestHighwayData", func(t *testing.T) {
		expectedHighwayData2022 := map[string]*AccidentData{
			"CONCRETO": {
				TotalAccident: 1717,
				TotalDeath:    9,
				TotalInvolved: 2452,
				TotalInjured:  2443,
			},
			"ASFALTO": {
				TotalAccident: 271347,
				TotalDeath:    4637,
				TotalInvolved: 425994,
				TotalInjured:  421357,
			},
			"NAO INFORMADO": {
				TotalAccident: 633207,
				TotalDeath:    11832,
				TotalInvolved: 1053061,
				TotalInjured:  1041229,
			},
			"PARALELEPIPEDO": {
				TotalAccident: 2046,
				TotalDeath:    20,
				TotalInvolved: 2423,
				TotalInjured:  2403,
			},
			"DESCONHECIDO": {
				TotalAccident: 69415,
				TotalDeath:    3598,
				TotalInvolved: 12123,
				TotalInjured:  8525,
			},
			"TERRA": {
				TotalAccident: 3315,
				TotalDeath:    121,
				TotalInvolved: 4515,
				TotalInjured:  4394,
			},
			"CASCALHO": {
				TotalAccident: 1110,
				TotalDeath:    13,
				TotalInvolved: 1412,
				TotalInjured:  1399,
			},
		}

		yearData2022, ok := results.YearDataByHighway["2022"]
		if !ok {
			t.Fatalf("Resultados para 'tp_pavimento' no ano de 2022 não encontrados")
		}

		validateAccidentData(t, "tp_pavimento", yearData2022.TotalAcciden, expectedHighwayData2022)
	})

	t.Run("TestGuardrailData", func(t *testing.T) {
		expectedGuardrailData2022 := map[string]*AccidentData{
			"NAO": {
				TotalAccident: 5719,
				TotalDeath:    924,
				TotalInvolved: 6693,
				TotalInjured:  5769,
			},
			"NAO INFORMADO": {
				TotalAccident: 908843,
				TotalDeath:    18403,
				TotalInvolved: 1345879,
				TotalInjured:  1327476,
			},
			"DESCONHECIDO": {
				TotalAccident: 67574,
				TotalDeath:    899,
				TotalInvolved: 149364,
				TotalInjured:  148465,
			},
			"SIM": {
				TotalAccident: 21,
				TotalDeath:    4,
				TotalInvolved: 44,
				TotalInjured:  40,
			},
		}

		yearData2022, ok := results.YearDataByGuardrail["2022"]
		if !ok {
			t.Fatalf("Resultados para 'ind_guardrail' no ano de 2022 não encontrados")
		}

		validateAccidentData(t, "ind_guardrail", yearData2022.TotalAcciden, expectedGuardrailData2022)
	})

	t.Run("TestWeekDayData", func(t *testing.T) {
		expectedDayWeekData2022 := map[string]*AccidentData{
			"QUINTA-FEIRA": {
				TotalAccident: 140618,
				TotalDeath:    2190,
				TotalInvolved: 216534,
				TotalInjured:  214344,
			},
			"DOMINGO": {
				TotalAccident: 118552,
				TotalDeath:    4506,
				TotalInvolved: 176180,
				TotalInjured:  171674,
			},
			"SEGUNDA-FEIRA": {
				TotalAccident: 142571,
				TotalDeath:    2582,
				TotalInvolved: 218428,
				TotalInjured:  215846,
			},
			"SEXTA-FEIRA": {
				TotalAccident: 159734,
				TotalDeath:    2801,
				TotalInvolved: 246568,
				TotalInjured:  243767,
			},
			"QUARTA-FEIRA": {
				TotalAccident: 137008,
				TotalDeath:    2169,
				TotalInvolved: 210275,
				TotalInjured:  208106,
			},
			"SABADO": {
				TotalAccident: 146965,
				TotalDeath:    3969,
				TotalInvolved: 223910,
				TotalInjured:  219941,
			},
			"TERCA-FEIRA": {
				TotalAccident: 136709,
				TotalDeath:    2013,
				TotalInvolved: 210085,
				TotalInjured:  208072,
			},
		}

		yearData2022, ok := results.YearDataByDayWeek["2022"]
		if !ok {
			t.Fatalf("Resultados para 'dia_semana' no ano de 2022 não encontrados")
		}

		validateAccidentData(t, "dia_semana", yearData2022.TotalAcciden, expectedDayWeekData2022)
	})

	t.Run("TestClimateData", func(t *testing.T) {
		expectedClimateData2022 := map[string]*AccidentData{
			"NEVOEIRO  NEVOA OU FUMACA": {
				TotalAccident: 899,
				TotalDeath:    37,
				TotalInvolved: 1514,
				TotalInjured:  1477,
			},
			"NEVE": {
				TotalAccident: 3,
				TotalDeath:    0,
				TotalInvolved: 6,
				TotalInjured:  6,
			},
			"VENTOS FORTES": {
				TotalAccident: 34,
				TotalDeath:    0,
				TotalInvolved: 32,
				TotalInjured:  32,
			},
			"OUTRAS CONDICOES": {
				TotalAccident: 80895,
				TotalDeath:    614,
				TotalInvolved: 115626,
				TotalInjured:  115012,
			},
			"NAO INFORMADO": {
				TotalAccident: 457678,
				TotalDeath:    8696,
				TotalInvolved: 693935,
				TotalInjured:  685239,
			},
			"GAROACHUVISCO": {
				TotalAccident: 2479,
				TotalDeath:    98,
				TotalInvolved: 5215,
				TotalInjured:  5117,
			},
			"NUBLADO": {
				TotalAccident: 6144,
				TotalDeath:    150,
				TotalInvolved: 13428,
				TotalInjured:  13278,
			},
			"CLARO": {
				TotalAccident: 267516,
				TotalDeath:    5623,
				TotalInvolved: 499252,
				TotalInjured:  493629,
			},
			"DESCONHECIDAS": {
				TotalAccident: 145762,
				TotalDeath:    4694,
				TotalInvolved: 133216,
				TotalInjured:  128522,
			},
			"CHUVA": {
				TotalAccident: 20747,
				TotalDeath:    318,
				TotalInvolved: 39756,
				TotalInjured:  39438,
			},
		}

		yearData2022, ok := results.YearDataByClimate["2022"]
		if !ok {
			t.Fatalf("Resultados para 'cond_meteorologica' no ano de 2022 não encontrados")
		}

		validateAccidentData(t, "cond_meteorologica", yearData2022.TotalAcciden, expectedClimateData2022)
	})
}

func validateAccidentData(t *testing.T, category string, actualData, expectedData map[string]*AccidentData) {
	t.Helper() // Marca a função como um helper de teste

	for condition, expectedConditionData := range expectedData {
		resultConditionData, ok := actualData[condition]
		if !ok {
			t.Errorf("%s '%s': Resultados não encontrados. Esperava dados para '%s'", category, condition, condition)
			continue // Pular para a próxima condição se não for encontrada
		}

		// Validar TotalAccident
		if resultConditionData.TotalAccident != expectedConditionData.TotalAccident {
			t.Errorf("%s '%s': esperava %d acidentes, obteve %d",
				category, condition, expectedConditionData.TotalAccident, resultConditionData.TotalAccident)
		}

		// Validar TotalDeath
		if resultConditionData.TotalDeath != expectedConditionData.TotalDeath {
			t.Errorf("%s '%s': esperava %d óbitos, obteve %d",
				category, condition, expectedConditionData.TotalDeath, resultConditionData.TotalDeath)
		}

		// Validar TotalInvolved
		if resultConditionData.TotalInvolved != expectedConditionData.TotalInvolved {
			t.Errorf("%s '%s': esperava %d envolvidos, obteve %d",
				category, condition, expectedConditionData.TotalInvolved, resultConditionData.TotalInvolved)
		}

		// Validar TotalInjured
		if resultConditionData.TotalInjured != expectedConditionData.TotalInjured {
			t.Errorf("%s '%s': esperava %d feridos, obteve %d",
				category, condition, expectedConditionData.TotalInjured, resultConditionData.TotalInjured)
		}
	}
}
