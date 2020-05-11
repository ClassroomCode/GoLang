package email_test

import (
	"testing"

	"github.com/gopherguides/email"
)

func Test_ValidateEmail(t *testing.T) {
	tcs := []struct {
		name      string
		em        string
		expectErr bool
	}{
		{name: "blank", em: "", expectErr: true},
		{name: "at", em: "yahoo.com", expectErr: true},
		{name: "period", em: "homer@yahoocom", expectErr: true},
		{name: "valid", em: "homer@yahoo.com"},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			err := email.Validate(tc.em)
			if err == nil && tc.expectErr {
				t.Fatal("expected error, got nil")
			}
			if err != nil && !tc.expectErr {
				t.Fatal(err)
			}
		})
	}
}
