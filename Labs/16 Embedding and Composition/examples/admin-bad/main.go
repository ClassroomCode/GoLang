package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Admin struct {
	User
	Permissions map[string]bool
}

// section: user
func greetUser(u User) {
	fmt.Println(u.Name)
}

// section: user

func main() {
	a := Admin{
		User: User{
			Name: "Homer",
		},
	}
	// section: greet
	greetUser(a)
	// section: greet
}
