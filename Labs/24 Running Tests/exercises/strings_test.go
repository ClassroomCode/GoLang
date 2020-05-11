// +build !integration

package exercise_test

import (
	"strings"
	"testing"
	"time"
)

func TestTitle(t *testing.T) {
	tests := []struct {
		value string
		exp   string
	}{
		{value: "rob", exp: "Rob"},
		{value: "sandy", exp: "Sandy"},
		{value: "margaret", exp: "Margaret"},
		{value: "bill", exp: "Bill"},
		{value: "tom", exp: "Tom"},
	}

	for _, test := range tests {
		t.Run(test.value, func(t *testing.T) {
			got := strings.Title(test.value)
			if got != test.exp {
				t.Logf("testing value: %s", test.value)
				t.Errorf("unexpected value.  got: %s, exp: %s", got, test.exp)
			}
		})
	}
}

func TestLong(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test due to short testing detected")
	}
	// long running test
	time.Sleep(3 * time.Second)
}
