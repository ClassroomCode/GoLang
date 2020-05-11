package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// section: user
func (u User) Greet() {
	fmt.Printf("Hello, %s\n", u.Name)
}

// section: user

type Admin struct {
	User
	Permissions map[string]bool
}

// section: greeter
type greeter interface {
	Greet()
}

// section: greeter

// section: func
func greet(g greeter) {
	g.Greet()
}

// section: func

func main() {
	a := Admin{
		User: User{
			Name: "Homer",
		},
	}
	// section: greet
	greet(a)
	// section: greet
}
