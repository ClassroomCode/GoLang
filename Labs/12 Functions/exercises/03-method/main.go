package main

import "fmt"

type User struct {
	First string
	Last  string
	Email string
}

// Create a method on User called `FullName` that takes no arguments,
// and returns the First and Last name as a string
// Example User{First: "Rob", Last: "Pike", Email: "commander@google.com"} would return "Rob Pike"

// Create a method on User called `FormattedEmail` that takes no arguments,
// and returns a valid formatted email as a string
// Example User{First: "Rob", Last: "Pike", Email: "commander@google.com"} would return "Rob Pike <commander@google.com>"

// Create a method on User called `Reset` that sets all the fields to a blank string.
// Hint: we need to change the value so we need a special receiver for that...

func main() {
	u := User{"Rob", "Pike", "commander@google.com"}
	fmt.Printf("%+v\n", u)
	fmt.Println("Full Name", u.FullName())
	fmt.Println("Formatted Email", u.FormattedEmail())
	u.Reset()
	fmt.Printf("%+v\n", u)
}
