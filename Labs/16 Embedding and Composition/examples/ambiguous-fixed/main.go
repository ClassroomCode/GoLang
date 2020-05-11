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

func (a Admin) Greet() {
	if a.User.Name != "" {
		a.User.Greet()
		return
	}
	a.Guest.Greet()
}

func main() {
	a := Admin{
		User: User{Name: "Cory"},
	}
	a.Greet()
}
