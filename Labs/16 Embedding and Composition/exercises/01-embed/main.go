package main

import (
	"errors"
	"fmt"
)

// #TODO: Create a struct type called `Email`
// Add the following fields:
// - DisplayName: string
// - Username: string
// - Domain: string

// #TODO: Create a method on the Email type called `Address` that takes no arguments
// It should return a string in the format:
// `DisplayName <Username@Domain>`
// You can use `fmt.Sprintf` to create the return string

// #TODO: Create a struct type called Invite
// Add the following fields:
// - ID: int
// - Email (embedded type Email)

func main() {
	// #TODO: initialize a variable called `invite` of Type Invite with the following values:
	// - ID: 1
	// - Email
	//   - Display Name: "John Smith"
	//   - Username: "jsmith"
	//   - Domain: "yahoo.com"

	if err := SendEmail(invite); err != nil {
		fmt.Printf("error sending email: %s", err)
	}
}

func SendEmail(i Invite) error {
	address := i.Address()
	if address == "" {
		return errors.New("no email address provided")
	}
	fmt.Printf("sent email to: %s", i.Address())
	return nil
}
