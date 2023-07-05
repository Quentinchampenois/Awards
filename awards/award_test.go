package awards

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAward_Summarize(t *testing.T) {
	t.Run("multiple"+
		"multiple awards", func(t *testing.T) {
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
	})
	t.Run("single award", func(t *testing.T) {
		award := Award{
			ID:         "1",
			Film:       "Joker",
			Year:       2000,
			Award:      1,
			Nomination: 10,
		}

		want := fmt.Sprintf("%s - %d : has been nominated %d times and received %d award", award.Film, award.Year, award.Nomination, award.Award)
		expectation := award.Summarize()
		if expectation != want {
			t.Fatalf(`award.Summarize() = %s, want match for %s`, expectation, want)
		}
	})
	t.Run("single nomination", func(t *testing.T) {
		award := Award{
			ID:         "1",
			Film:       "Joker",
			Year:       2000,
			Award:      1,
			Nomination: 1,
		}

		want := fmt.Sprintf("%s - %d : has been nominated %d time and received %d award", award.Film, award.Year, award.Nomination, award.Award)
		expectation := award.Summarize()
		if expectation != want {
			t.Fatalf(`award.Summarize() = %s, want match for %s`, expectation, want)
		}
	})
	t.Run("multiple nominations", func(t *testing.T) {
		award := Award{
			ID:         "1",
			Film:       "Joker",
			Year:       2000,
			Award:      1,
			Nomination: 10,
		}

		want := fmt.Sprintf("%s - %d : has been nominated %d times and received %d award", award.Film, award.Year, award.Nomination, award.Award)
		expectation := award.Summarize()
		if expectation != want {
			t.Fatalf(`award.Summarize() = %s, want match for %s`, expectation, want)
		}
	})

}

func TestSortByYear(t *testing.T) {
	awards := []Award{{
		ID:         "0",
		Award:      3,
		Nomination: 10,
		Film:       "Foobar",
		Year:       2020,
	},
		{
			ID:         "1",
			Award:      1,
			Nomination: 3,
			Film:       "Barfoo",
			Year:       2020,
		},
		{
			ID:         "2",
			Award:      10,
			Nomination: 3,
			Film:       "not2020",
			Year:       2018,
		},
	}
	subject := Awards{Awards: awards}
	want := []Award{awards[2], awards[0], awards[1]}
	expectation := subject.SortByYear()
	assert.Equal(t, expectation, want)
}
