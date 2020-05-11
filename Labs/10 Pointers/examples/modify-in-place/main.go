package main

import "fmt"

type Beatle struct {
	Name string
}

func main() {
	b := Beatle{Name: "Ringo"}
	fmt.Println(b.Name) // Ringo
	b.Name = "George"
	fmt.Println(b.Name) // George
}
