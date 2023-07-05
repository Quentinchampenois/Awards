package main

import (
	"awards/awards"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("src/oscar.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error opening file:", err)
			return
		}
	}(file)

	reader := csv.NewReader(file)
	reader.Comma = ','

	lines := 0
	var allAwards awards.Awards
	for {
		row, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading row:", err)
			continue
		}

		year, err := strconv.Atoi(row[2])
		if err != nil {
			continue
		}
		award, err := strconv.Atoi(row[3])
		if err != nil {
			continue
		}
		nomination, err := strconv.Atoi(row[4])
		if err != nil {
			continue
		}

		item := awards.Award{
			ID:         row[0],
			Film:       row[1],
			Year:       year,
			Award:      award,
			Nomination: nomination,
		}

		lines += 1
		allAwards.Awards = append(allAwards.Awards, item)
	}

	grouped := allAwards.GroupByYear(2012)
	for _, award := range grouped {
		fmt.Println(award.Summarize())
	}

	best := allAwards.BestForTheYear(2012)

	fmt.Println(best.Summarize())
	fmt.Println("Found films : ", len(allAwards.Awards))
	fmt.Println("Total lines : ", lines)
}
