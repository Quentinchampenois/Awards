package awards

import (
	"fmt"
	"sort"
)

type Award struct {
	ID         string
	Film       string
	Year       int
	Award      int
	Nomination int
}

func (o *Award) Summarize() string {
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

type Awards struct {
	Awards []Award
}

func (a *Awards) SortByYear() []Award {
	awards := a.Awards

	sort.SliceStable(awards, func(i, j int) bool {
		return awards[i].Year < awards[j].Year
	})

	return awards
}

func (a *Awards) GroupByYear(year int) []Award {
	var group []Award

	for _, award := range a.Awards {
		if award.Year == year {
			group = append(group, award)
		}
	}

	return group
}

func (a *Awards) BestForTheYear(year int) Award {
	var best Award
	a.Awards = a.GroupByYear(year)

	for _, award := range a.Awards {
		if best.Award < award.Award {
			best = award
		}
	}

	return best
}
