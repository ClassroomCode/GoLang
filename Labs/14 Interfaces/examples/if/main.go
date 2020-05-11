package main

import (
	"fmt"
)

type User struct {
	First string
	Last  string
}

func (u User) String() string {
	return fmt.Sprintf("User name is %s %s", u.First, u.Last)
}

func main() {
	u := User{First: "Rob", Last: "Pike"}
	print(u)
	print("bob")
}

// section: print
func print(s interface{}) {
	stringer, ok := s.(fmt.Stringer)
	if ok {
		fmt.Println(stringer.String())
		return
	}
	fmt.Println("not a stringer")
}

// section: print

// section: printinline
func printInline(s interface{}) {
	if stringer, ok := s.(fmt.Stringer); ok {
		fmt.Println(stringer.String())
		return
	}
	fmt.Println("not a stringer")
}

// section: printinline
