package domain_test

import (
	"encoding/json"
	"os"
	"testing"
)

type User struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

// section: test
func Test_JSON_Valid(t *testing.T) {
	u := &User{}

	// call a helper method to load the file and decode it
	decode(t, "testdata/valid.json", u)
	if got, exp := u.First, "Jenny"; got != exp {
		t.Errorf("unexpected first name.  got: %s, exp: %s", got, exp)
	}
	if got, exp := u.Last, "Smith"; got != exp {
		t.Errorf("unexpected last name.  got: %s, exp: %s", got, exp)
	}
}

// section: test

// section: helper
func decode(t *testing.T, fn string, u *User) {
	r, err := os.Open(fn)
	if err != nil {
		t.Fatal(err)
	}
	d := json.NewDecoder(r)
	err = d.Decode(u)
	if err != nil {
		t.Fatal(err)
	}
}

// section: helper
