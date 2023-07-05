package awards

import (
	"fmt"
	"testing"
)

func TestAward_Summarize(t *testing.T) {
	award := Award{
		ID:         "1",
		Film:       "Joker",
		Year:       2000,
		Award:      8,
		Nomination: 10,
	}

	want := fmt.Sprintf("%s - %d : has been nominated %d times and received %d awards", award.Film, award.Year, award.Nomination, award.Award)
	expectation := award.Summarize()
	if expectation != want {
		t.Fatalf(`award.Summarize() = %s, want match for %s`, expectation, want)
	}
}
