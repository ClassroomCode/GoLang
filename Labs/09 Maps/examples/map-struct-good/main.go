package main

import "fmt"

// section: code
type simple struct {
	ID   int
	Name string
}

func main() {
	data := map[int]simple{}
	s1 := simple{ID: 1, Name: "Cory"}
	data[1] = s1

	if value, ok := data[10]; ok {
		fmt.Printf("%+v", value)
		return
	}
	fmt.Println("value not found")
}

// section: code
