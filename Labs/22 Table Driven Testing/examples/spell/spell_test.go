package spell_test

import (
	"testing"

	"github.com/gopherguides/spell"
)

// section: test
func TestAutoCorrect(t *testing.T) {
	// section: first
	word := "accaptable"
	got := spell.AutoCorrect(word)
	exp := "acceptable"
	if got != exp {
		t.Errorf("unexpected value for %q. got: %q, exp: %q", word, got, exp)
	}
	// section: first

	// section: second
	word = "accidentilly"
	got = spell.AutoCorrect(word)
	exp = "accidentally"
	if got != exp {
		t.Errorf("unexpected value for %q. got: %q, exp: %q", word, got, exp)
	}

	word = "accommadate"
	got = spell.AutoCorrect(word)
	exp = "accommodate"
	if got != exp {
		t.Errorf("unexpected value for %q. got: %q, exp: %q", word, got, exp)
	}
	// section: second

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

// section: test

// section: table
func TestAutoCorrectTable(t *testing.T) {
	// Define an anonymous struct that defines the test cases
	// tcs stands for "test cases"
	tcs := []struct {
		word string
		exp  string
	}{
		{word: "accaptable", exp: "acceptable"},
		{word: "accidentilly", exp: "accidentally"},
		{word: "accommadate", exp: "accommodate"},
		{word: "acqire", exp: "acquire"},
		{word: "alot", exp: "allot"},
		{word: "amatuer", exp: "amateur"},
		{word: "apparint", exp: "apparent"},
	}

	for _, tc := range tcs { // tc stands for "test case"
		got := spell.AutoCorrect(tc.word)
		if got != tc.exp {
			t.Errorf("unexpected value for %q. got: %q, exp: %q", tc.word, got, tc.exp)
		}
	}
}

// section: table

// section: sub
func TestAutoCorrectTableSub(t *testing.T) {
	// Define an anonymous struct that defines the test cases
	// tcs stands for "test cases"
	tcs := []struct {
		word string
		exp  string
	}{
		{word: "accaptable", exp: "acceptable"},
		{word: "accidentilly", exp: "accidentally"},
		{word: "accommadate", exp: "accommodate"},
		{word: "acqire", exp: "acquire"},
		{word: "alot", exp: "allot"},
		{word: "amatuer", exp: "amateur"},
		{word: "apparint", exp: "apparent"},
	}

	for _, tc := range tcs { // tc stands for "test case"
		t.Run(tc.word, func(t *testing.T) {
			got := spell.AutoCorrect(tc.word)
			if got != tc.exp {
				t.Errorf("unexpected value for %q. got: %q, exp: %q", tc.word, got, tc.exp)
			}
		})
	}
}

// section: sub

/*
// section: subtestoutput
=== RUN   TestAutoCorrectTableSub
=== RUN   TestAutoCorrectTableSub/accaptable
=== RUN   TestAutoCorrectTableSub/accidentilly
=== RUN   TestAutoCorrectTableSub/accommadate
=== RUN   TestAutoCorrectTableSub/acqire
=== RUN   TestAutoCorrectTableSub/alot
=== RUN   TestAutoCorrectTableSub/amatuer
=== RUN   TestAutoCorrectTableSub/apparint
--- PASS: TestAutoCorrectTableSub (0.00s)
    --- PASS: TestAutoCorrectTableSub/accaptable (0.00s)
    --- PASS: TestAutoCorrectTableSub/accidentilly (0.00s)
    --- PASS: TestAutoCorrectTableSub/accommadate (0.00s)
    --- PASS: TestAutoCorrectTableSub/acqire (0.00s)
    --- PASS: TestAutoCorrectTableSub/alot (0.00s)
    --- PASS: TestAutoCorrectTableSub/amatuer (0.00s)
    --- PASS: TestAutoCorrectTableSub/apparint (0.00s)
PASS
ok      github.com/gopherguides/learn/_training/testing/table/src/spell 0.006s
// section: subtestoutput
*/

// section: parallelbug
func TestAutoCorrectTableParallelBug(t *testing.T) {
	// Define an anonymous struct that defines the test cases
	// tcs stands for "test cases"
	tcs := []struct {
		word string
		exp  string
	}{
		{word: "accaptable", exp: "acceptable"},
		{word: "accidentilly", exp: "accidentally"},
		{word: "accommadate", exp: "accommodate"},
		{word: "acqire", exp: "acquire"},
		{word: "alot", exp: "allot"},
		{word: "amatuer", exp: "amateur"},
		{word: "apparint", exp: "apparent"},
	}

	for _, tc := range tcs { // tc stands for "test case"
		t.Run(tc.word, func(t *testing.T) {
			t.Parallel()
			got := spell.AutoCorrect(tc.word)
			t.Logf("testing %q", tc.word)
			if got != tc.exp {
				t.Errorf("unexpected value for %q. got: %q, exp: %q", tc.word, got, tc.exp)
			}
		})
	}
}

// section: parallelbug

/*
// section: parallelbugoutput
=== RUN   TestAutoCorrectTableParallelBug
=== RUN   TestAutoCorrectTableParallelBug/accaptable
=== PAUSE TestAutoCorrectTableParallelBug/accaptable
=== RUN   TestAutoCorrectTableParallelBug/accidentilly
=== PAUSE TestAutoCorrectTableParallelBug/accidentilly
=== RUN   TestAutoCorrectTableParallelBug/accommadate
=== PAUSE TestAutoCorrectTableParallelBug/accommadate
=== RUN   TestAutoCorrectTableParallelBug/acqire
=== PAUSE TestAutoCorrectTableParallelBug/acqire
=== RUN   TestAutoCorrectTableParallelBug/alot
=== PAUSE TestAutoCorrectTableParallelBug/alot
=== RUN   TestAutoCorrectTableParallelBug/amatuer
=== PAUSE TestAutoCorrectTableParallelBug/amatuer
=== RUN   TestAutoCorrectTableParallelBug/apparint
=== PAUSE TestAutoCorrectTableParallelBug/apparint
=== CONT  TestAutoCorrectTableParallelBug/accaptable
=== CONT  TestAutoCorrectTableParallelBug/acqire
=== CONT  TestAutoCorrectTableParallelBug/apparint
=== CONT  TestAutoCorrectTableParallelBug/alot
=== CONT  TestAutoCorrectTableParallelBug/amatuer
=== CONT  TestAutoCorrectTableParallelBug/accommadate
=== CONT  TestAutoCorrectTableParallelBug/accidentilly
--- PASS: TestAutoCorrectTableParallelBug (0.00s)
    --- PASS: TestAutoCorrectTableParallelBug/acqire (0.00s)
        spell_test.go:167: testing "apparint"
    --- PASS: TestAutoCorrectTableParallelBug/accaptable (0.00s)
        spell_test.go:167: testing "apparint"
    --- PASS: TestAutoCorrectTableParallelBug/apparint (0.00s)
        spell_test.go:167: testing "apparint"
    --- PASS: TestAutoCorrectTableParallelBug/amatuer (0.00s)
        spell_test.go:167: testing "apparint"
    --- PASS: TestAutoCorrectTableParallelBug/alot (0.00s)
        spell_test.go:167: testing "apparint"
    --- PASS: TestAutoCorrectTableParallelBug/accidentilly (0.00s)
        spell_test.go:167: testing "apparint"
    --- PASS: TestAutoCorrectTableParallelBug/accommadate (0.00s)
        spell_test.go:167: testing "apparint"
PASS
ok      github.com/gopherguides/learn/_training/testing/table/src/spell 0.008s
// section: parallelbugoutput
*/

// section: parallel
func TestAutoCorrectTableParallel(t *testing.T) {
	// Define an anonymous struct that defines the test cases
	// tcs stands for "test cases"
	tcs := []struct {
		word string
		exp  string
	}{
		{word: "accaptable", exp: "acceptable"},
		{word: "accidentilly", exp: "accidentally"},
		{word: "accommadate", exp: "accommodate"},
		{word: "acqire", exp: "acquire"},
		{word: "alot", exp: "allot"},
		{word: "amatuer", exp: "amateur"},
		{word: "apparint", exp: "apparent"},
	}

	for _, tc := range tcs { // tc stands for "test case"
		t.Run(tc.word, func(t *testing.T) {
			tc := tc // rebind tc into this lexical scope
			t.Parallel()
			got := spell.AutoCorrect(tc.word)
			t.Logf("testing %q", tc.word)
			if got != tc.exp {
				t.Errorf("unexpected value for %q. got: %q, exp: %q", tc.word, got, tc.exp)
			}
		})
	}
}

// section: parallel

/*

// section: paralleloutput
=== RUN   TestAutoCorrectTableParallel
=== RUN   TestAutoCorrectTableParallel/accaptable
=== PAUSE TestAutoCorrectTableParallel/accaptable
=== RUN   TestAutoCorrectTableParallel/accidentilly
=== PAUSE TestAutoCorrectTableParallel/accidentilly
=== RUN   TestAutoCorrectTableParallel/accommadate
=== PAUSE TestAutoCorrectTableParallel/accommadate
=== RUN   TestAutoCorrectTableParallel/acqire
=== PAUSE TestAutoCorrectTableParallel/acqire
=== RUN   TestAutoCorrectTableParallel/alot
=== PAUSE TestAutoCorrectTableParallel/alot
=== RUN   TestAutoCorrectTableParallel/amatuer
=== PAUSE TestAutoCorrectTableParallel/amatuer
=== RUN   TestAutoCorrectTableParallel/apparint
=== PAUSE TestAutoCorrectTableParallel/apparint
=== CONT  TestAutoCorrectTableParallel/accaptable
=== CONT  TestAutoCorrectTableParallel/apparint
=== CONT  TestAutoCorrectTableParallel/acqire
=== CONT  TestAutoCorrectTableParallel/alot
=== CONT  TestAutoCorrectTableParallel/accommadate
=== CONT  TestAutoCorrectTableParallel/amatuer
=== CONT  TestAutoCorrectTableParallel/accidentilly
--- PASS: TestAutoCorrectTableParallel (0.00s)
    --- PASS: TestAutoCorrectTableParallel/accaptable (0.00s)
        spell_test.go:244: testing "accaptable"
    --- PASS: TestAutoCorrectTableParallel/apparint (0.00s)
        spell_test.go:244: testing "apparint"
    --- PASS: TestAutoCorrectTableParallel/alot (0.00s)
        spell_test.go:244: testing "alot"
    --- PASS: TestAutoCorrectTableParallel/acqire (0.00s)
        spell_test.go:244: testing "acqire"
    --- PASS: TestAutoCorrectTableParallel/amatuer (0.00s)
        spell_test.go:244: testing "amatuer"
    --- PASS: TestAutoCorrectTableParallel/accommadate (0.00s)
        spell_test.go:244: testing "accommadate"
    --- PASS: TestAutoCorrectTableParallel/accidentilly (0.00s)
        spell_test.go:244: testing "accidentilly"
PASS
ok      github.com/gopherguides/learn/_training/testing/table/src/spell 0.008s
// section: paralleloutput
*/
