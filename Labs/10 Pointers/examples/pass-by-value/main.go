package main

import "fmt"

type Beatle struct {
	Name string
}

func main() {
	b := Beatle{Name: "Ringo"}
	changeBeatleName(b)
	fmt.Println(b.Name) // Ringo
}

// section: change
func changeBeatleName(b Beatle) {
	b.Name = "George"
	fmt.Println(b.Name) // George
}

// section: change
