package spell_test

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/gopherguides/spell"
)

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

	word = "accommadate"
	got = spell.AutoCorrect(word)
	exp = "accommodate"
	if got != exp {
		t.Errorf("unexpected value for %q. got: %q, exp: %q", word, got, exp)
	}

	word = "acqire"
	got = spell.AutoCorrect(word)
	exp = "acquire"
	if got != exp {
		t.Errorf("unexpected value for %q. got: %q, exp: %q", word, got, exp)
	}

	word = "alot"
	got = spell.AutoCorrect(word)
	exp = "allot"
	if got != exp {
		t.Errorf("unexpected value for %q. got: %q, exp: %q", word, got, exp)
	}

	word = "amatuer"
	got = spell.AutoCorrect(word)
	exp = "amateur"
	if got != exp {
		t.Errorf("unexpected value for %q. got: %q, exp: %q", word, got, exp)
	}

	word = "apparint"
	got = spell.AutoCorrect(word)
	exp = "apparent"
	if got != exp {
		t.Errorf("unexpected value for %q. got: %q, exp: %q", word, got, exp)
	}

}

func TestAutoCorrectFromFile(t *testing.T) {
	data := loadFile(t, "./testdata/terms.txt")
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

// Convert the following function to be a `testing.Helper` function
func loadFile(t *testing.T, file string) string {
	t.Helper()
	b, err := ioutil.ReadFile(file)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}
