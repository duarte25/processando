package sequencial

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type result struct {
	numRows           int
	peopleCount       int
	commonName        string
	commonNameCount   int
	donationMonthFreq map[string]int
}

// processRow takes a pipe-separated line and returns the firstName, fullName, and month.
// this function was created to be somewhat compute intensive and not accurate.
func processRow(text string) (firstName, fullName, month string) {
	row := strings.Split(text, "|")

	// extract full name
	fullName = strings.Replace(strings.TrimSpace(row[7]), " ", "", -1)

	// extract first name
	name := strings.TrimSpace(row[7])
	if name != "" {
		startOfName := strings.Index(name, ", ") + 2
		if endOfName := strings.Index(name[startOfName:], " "); endOfName < 0 {
			firstName = name[startOfName:]
		} else {
			firstName = name[startOfName : startOfName+endOfName]
		}
		if strings.HasSuffix(firstName, ",") {
			firstName = strings.Replace(firstName, ",", "", -1)
		}
	}

	// extract month
	date := strings.TrimSpace(row[13])
	if len(date) == 8 {
		month = date[:2]
	} else {
		month = "--"
	}

	return firstName, fullName, month
}

// sequential processes a file line by line using processRow.
func Sequential(file string) result {
	res := result{donationMonthFreq: map[string]int{}}

	start := time.Now()

	end := time.Now() // Captura o tempo de término

	duration := end.Sub(start) // Calcula a diferença de tempo

	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tempo gasto abrindo arquivo: %v\n", duration)

	// track full names
	fullNamesRegister := make(map[string]bool)

	// track first name frequency
	firstNameMap := make(map[string]int)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		row := scanner.Text()
		firstName, fullName, month := processRow(row)

		// add fullname
		fullNamesRegister[fullName] = true

		// update common firstName
		firstNameMap[firstName]++
		if firstNameMap[firstName] > res.commonNameCount {
			res.commonName = firstName
			res.commonNameCount = firstNameMap[firstName]
		}
		// add month freq
		res.donationMonthFreq[month]++
		// update numRows
		res.numRows++
		res.peopleCount = len(fullNamesRegister)
	}

	return res
}
