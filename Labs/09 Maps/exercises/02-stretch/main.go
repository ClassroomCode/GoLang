package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func main() {
	users := make(map[int]User)

	u := User{ID: 1, Name: "Cory Bates"}
	users[u.ID] = u

	u = User{ID: 2, Name: "Mark Raymond"}
	users[u.ID] = u

	u = User{ID: 3, Name: "Tim LaNou"}
	users[u.ID] = u

	// Look up Cory

	// Update Cory's name to "Cory LaNou"

	// Set the user of ID `1` to the updated user

	// Fix the remaining users names:
	// ID:2, Mark -> "Mark Bates"
	// ID:3, Tim -> "Tim Raymond"

	fmt.Println(users)
}
