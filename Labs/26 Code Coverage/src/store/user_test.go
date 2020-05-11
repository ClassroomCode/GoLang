package store_test

import (
	"testing"

	"github.com/gopherguides/learn/_training/testing/coverage/src/store"
)

// section: validate
func TestUserValidate(t *testing.T) {
	// define our tests
	// section: tests
	tcs := []struct {
		name   string
		user   store.User
		expErr bool
	}{
		{name: "empty", user: store.User{}, expErr: true},
		{name: "first", user: store.User{First: "Rob"}, expErr: true},
		{name: "last", user: store.User{Last: "Pike"}, expErr: true},
		{name: "valid", user: store.User{First: "Rob", Last: "Pike"}, expErr: false},
	}
	// section: tests

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.user.Validate()
			if tc.expErr && err == nil {
				t.Error("expected validation error, got nil")
				return
			}
			if !tc.expErr && err != nil {
				t.Errorf("unexpected validation error, got: %s", err)
			}
		})
	}
}

// section: validate

func TestProperName(t *testing.T) {
	u := store.User{First: "Rob", Last: "Pike"}
	got, err := u.ProperName()
	if err != nil {
		t.Fatal(err)
	}
	exp := "Rob Pike"
	if got != exp {
		t.Errorf("unexpected propername.  got: %q, exp: %q", got, exp)
	}
}
