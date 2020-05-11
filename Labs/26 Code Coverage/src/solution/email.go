package email

import (
	"errors"
	"strings"
)

func Validate(e string) error {
	if len(e) == 0 {
		return errors.New("email can't be blank")
	}
	if !strings.Contains(e, "@") {
		return errors.New("invalid email, missing '@' symbol")
	}
	if !strings.Contains(e, ".") {
		return errors.New("invalid email, missing '.' symbol")
	}
	return nil
}
