package models

import "fmt"

// user is unexported
type user struct {
	Name string
	Age  int
}

func (u user) Greet() {
	fmt.Printf("Hello, %s\n", u.Name)
}

type Admin struct {
	user        // embedding the unexported type
	Permissions map[string]bool
}

func DefaultAdmin() Admin {
	return Admin{
		user: user{
			Name: "Homer",
		},
	}

}
