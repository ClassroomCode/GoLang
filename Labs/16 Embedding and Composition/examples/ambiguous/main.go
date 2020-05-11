package main

import (
	"fmt"
)

type User struct {
	Name string
}

func (u User) Greet() {
	fmt.Println("hello: ", u.Name)
}

type Guest struct{}

func (g Guest) Greet() {
	fmt.Println("hello guest")
}

type Admin struct {
	User
	Guest
}

func main() {
	a := Admin{
		User: User{Name: "Cory"},
	}
	a.Greet()
}
