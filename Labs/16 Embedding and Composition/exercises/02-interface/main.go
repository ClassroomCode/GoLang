package main

import (
	"errors"
	"fmt"
)

// #TODO: Declare an interface called `Addressable` that defines the `Address` method
// from the Email

type Invite struct {
	ID int
	Email
}

type Email struct {
	DisplayName string
	Username    string
	Domain      string
}

func (e Email) Address() string {
	return fmt.Sprintf("%s <%s@%s>", e.DisplayName, e.Username, e.Domain)
}

// #TODO: Declare a type called `spamTrack` from the `Invite` type

// #TODO: create a method `Address` to override the Invite address method
// - Retrieve the address from the original Invite
// - add in `+spam` identifier in the email to change all emails from `name@domain.com` to `name+spam@domain.com`
//   hint: you can use the following command: strings.ReplaceAll(e, "@", "+spam@")
// - return the new `spam tracker` email instead of the original email

func main() {
	invite := Invite{
		ID: 1,
		Email: Email{
			DisplayName: "John Smith",
			Username:    "jsmith",
			Domain:      "yahoo.com",
		},
	}
	if err := SendEmail(invite); err != nil {
		fmt.Printf("error sending email: %s", err)
	}

	// #TODO: repeat the above command to send an email, but this time override the invite with the new `spamTrack` type
	// to insert a spam tracking email
}

// #TODO: Change the signature of the `SendEmail` function to take
// an `Addressable` interface, instead of the concrete type `Invite`
func SendEmail(i Invite) error {
	address := i.Address()
	if address == "" {
		return errors.New("no email address provided")
	}
	fmt.Printf("sent email to: %s\n", i.Address())
	return nil
}
