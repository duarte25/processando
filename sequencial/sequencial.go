package sequencial

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
