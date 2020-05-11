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

// section: admin
func (a Admin) Greet() {
	fmt.Printf("Admin: %s\n", a.Name)
}

// section: admin

// section: main
func main() {
	a := Admin{
		User: User{
			Name: "Homer",
		},
	}

	a.Greet() // Admin: Homer
}

// section: main
