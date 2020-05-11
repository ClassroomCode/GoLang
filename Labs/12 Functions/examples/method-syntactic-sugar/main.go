package main

import (
	"fmt"
)

type User struct {
	First string
	Last  string
	Email string
}

func (u User) Notify(msg string) error {
	if err := sendEmail(u.Email, msg); err != nil {
		return err
	}
	return nil
}

func main() {
	// Define a user
	u := User{First: "Cory", Last: "LaNou", Email: "cory@gopherguides.com"}

	// call the "Notify" as a method on "User"
	if err := u.Notify("This is a message as a method"); err != nil {
		fmt.Printf("oh no, something went wrong! %v", err)
		return
	}

	// call the "Notify" method as a function from the concrete type "User"
	if err := User.Notify(u, "This is a message as a method expression"); err != nil {
		fmt.Printf("oh no, something went wrong! %v", err)
		return
	}

	// Notice that a method is just "syntactic sugar" by hanging a
	// function off of a concrete type

	fmt.Println("notification successfully sent")
}

// sendEmail simulates sending an email.
func sendEmail(em, msg string) error {
	fmt.Printf("sending message %q to %q\n", msg, em)
	return nil
}
