package main

import (
	"fmt"
	"strings"
)

type User struct {
	First string
	Last  string
}

// Pointer Receiver
func (u *User) Titleize() {
	u.First = strings.Title(u.First)
	u.Last = strings.Title(u.Last)
}

// Value Receiver
func (u User) Reset() {
	u.First = ""
	u.Last = ""
}

func main() {
	u := User{First: "cory", Last: "lanou"}
	u.Titleize()
	fmt.Println(u)

	// Reset can't mutate the data as it's defined with a value receiver
	u.Reset()
	fmt.Println(u)
}
