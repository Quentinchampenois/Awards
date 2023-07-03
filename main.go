package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Oscar struct {
	ID         string
	Film       string
	Year       int
	Award      int
	Nomination int
}

func (o *Oscar) Summarize() string {
	var timeStr string
	if o.Nomination > 1 {
		timeStr = "times"
	} else {
		timeStr = "time"
	}

	var awardStr string
	if o.Award > 1 {
		awardStr = "awards"
	} else {
		awardStr = "award"
	}
	return fmt.Sprintf("%s - %d : has been nominated %d %s and received %d %s", o.Film, o.Year, o.Nomination, timeStr, o.Award, awardStr)
}

func sortByYears(oscars []Oscar) []Oscar {
	sort.SliceStable(oscars, func(i, j int) bool {
		return oscars[i].Year < oscars[j].Year
	})

	return oscars
}

func groupByYear(oscars []Oscar, year int) []Oscar {
	var group []Oscar

	for _, oscar := range oscars {
		if oscar.Year == year {
			group = append(group, oscar)
		}
	}

	return group
}

func bestForTheYear(oscars []Oscar, year int) Oscar {
	var best Oscar
	oscars = groupByYear(oscars, year)

	for _, oscar := range oscars {
		if best.Award < oscar.Award {
			best = oscar
		}
	}

	return best
}

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
	var oscars []Oscar
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
		oscar := Oscar{
			ID:         row[0],
			Film:       row[1],
			Year:       year,
			Award:      award,
			Nomination: nomination,
		}

		lines += 1
		oscars = append(oscars, oscar)
	}

	//sortByYears(oscars)
	oscars = groupByYear(oscars, 2012)
	for _, oscar := range oscars {
		fmt.Println(oscar.Summarize())
	}

	best := bestForTheYear(oscars, 2012)

	fmt.Println(best.Summarize())
	fmt.Println("Found films : ", len(oscars))
	fmt.Println("Total lines : ", lines)
}
