package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	u := User{"Homer Simpson", "homer@example.com"}
	fmt.Println(u)
}
