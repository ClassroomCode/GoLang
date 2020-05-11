package spell_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/gopherguides/spell"
)

// section: exercise
func TestAutoCorrect(t *testing.T) {
	word := "accaptable"
	got := spell.AutoCorrect(word)
	exp := "acceptable"
	if got != exp {
		t.Errorf("unexpected value for %q. got: %q, exp: %q", word, got, exp)
	}

	word = "accidentilly"
	got = spell.AutoCorrect(word)
	exp = "accidentally"
	if got != exp {
		t.Errorf("unexpected value for %q. got: %q, exp: %q", word, got, exp)
	}

	/*
		TODO Test for the following misspellings:
			"accommadate":  "accommodate",
			"acqire":       "acquire",
			"alot":         "allot",
			"amatuer":      "amateur",
			"apparint":     "apparent",
	*/

}

func TestAutoCorrectFromFile(t *testing.T) {
	// TODO Call `loadFile` as a test helper
	data := loadFile("./testdata/terms.txt")
	terms := strings.Split(data, "\n")
	for i, term := range terms {
		if len(term) > 0 {
			f := strings.Fields(term)
			if len(f) != 2 {
				t.Errorf("unexpected input: %s, line: %d", term, i+1)
				continue
			}
			if got, exp := spell.AutoCorrect(f[0]), f[1]; got != exp {
				t.Errorf("unexpected value for %q. got: %q, exp: %q", term, got, exp)
			}
		}
	}
}

// TODO: Convert the following function to be a `testing.Helper` function
func loadFile(file string) string {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// section: exercise
