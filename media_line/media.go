package media

import (
	"bufio"
	"log"
	"math"
	"os"
	"strings"
)

type Result struct {
	NumRows           int
	DonationMonthFreq map[string]int
	MediaRow          float64
}

func processRow(text string) (calcRow int) {

	//calcRow = len(text)

	posicaoBarra := strings.LastIndex(text, "|")
	parte := text[posicaoBarra:]

	calcRow = len(parte)

	return calcRow
}

// sequential processes a file line by line using processRow.
func Media(file string) Result {
	res := Result{DonationMonthFreq: map[string]int{}}

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	var totalChars int64

	for scanner.Scan() {
		row := scanner.Text()
		calcRow := processRow(row)

		// update numRows
		res.NumRows++
		totalChars += int64(calcRow)
	}

	res.MediaRow = math.Round(float64(totalChars) / float64(res.NumRows))

	return res
}
