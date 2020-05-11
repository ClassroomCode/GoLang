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
	u = users[1]

	// Update Cory's name to "Cory LaNou"
	u.Name = "Cory LaNou"

	// Set the user of ID `1` to the updated user
	users[u.ID] = u

	// Fix the remaining users names:
	// ID:2, Mark -> "Mark Bates"
	// ID:3, Tim -> "Tim Raymond"

	u = users[2]
	u.Name = "Mark Bates"
	users[u.ID] = u

	u = users[3]
	u.Name = "Tim Raymond"
	users[u.ID] = u

	fmt.Println(users)
}
