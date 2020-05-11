package store

import (
	"errors"
	"fmt"
)

// section: user
type User struct {
	First string
	Last  string
}

// section: validate
func (u User) Validate() error {
	if u.First == "" {
		return errors.New("first name can't be blank")
	}
	if u.Last == "" {
		return errors.New("last name can't be blank")
	}
	return nil
}

// section: validate

func (u User) ProperName() (string, error) {
	err := u.Validate()
	if err != nil {
		return "", fmt.Errorf("failed to generate propername: %s", err)
	}
	return fmt.Sprintf("%s %s", u.First, u.Last), nil
}

// section: user
