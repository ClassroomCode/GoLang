package main

import "fmt"

// section: user
type User struct {
	Name string
	Age  int
}

// section: user

// section: admin
type Admin struct {
	User        User
	Permissions map[string]bool
}

// section: admin

// section: main
func main() {
	a := Admin{
		User: User{
			Name: "Homer",
			Age:  40,
		},
		Permissions: map[string]bool{},
	}

	fmt.Println(a.User.Name)
}

// section: main
