package main

import "fmt"

type User struct {
	Name string
	Age  int
}

// section: user
func (u User) String() string {
	return u.Name
}

// section: user

// section: admin
type Admin struct {
	User
	Permissions map[string]bool
}

func (a Admin) String() string {
	return fmt.Sprintf("Admin: %s", a.User.String())
}

// section: admin

func main() {
	a := Admin{
		User: User{
			Name: "Homer",
		},
	}

	// section: print
	fmt.Println(a.String()) // Admin: Homer
	// section: print
}
