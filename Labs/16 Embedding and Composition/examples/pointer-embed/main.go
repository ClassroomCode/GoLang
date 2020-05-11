package main

import "fmt"

type User struct {
	name string
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) String() string {
	return u.name
}

// section: pointer
type Admin struct {
	*User
	Perms map[string]bool
}

// section: pointer

// section: main
func main() {
	a := &Admin{}
	fmt.Println(a.String())
}

// section: main
