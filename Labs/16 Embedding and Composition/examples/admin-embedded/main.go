package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// section: admin
type Admin struct {
	User
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

	fmt.Println(a.Name)
}

// section: main
