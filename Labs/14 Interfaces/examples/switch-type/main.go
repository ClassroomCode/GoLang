package main

import (
	"encoding/json"
	"fmt"
)

// section: main
type Beatle struct {
	Name       string
	Instrument string
}

func (b Beatle) String() string {
	s, _ := json.Marshal(b)
	return string(s)
}

func main() {
	print(42)
	print("Ringo Rules")
	print(3.141592)
	print(Beatle{Name: "John", Instrument: "Guitar"})
	// INT: 42
	// Ringo Rules
	// FLOAT: 3.141592
}

func print(i interface{}) {
	switch t := i.(type) {
	case int:
		t = t + 3
		fmt.Printf("INT: %d\n", t)
	case float64:
		fmt.Printf("FLOAT: %f\n", t)
	case fmt.Stringer:
		fmt.Printf("%T\n", t)
		fmt.Println(t.String())
	default:
		fmt.Println(t)
	}
}

// section: main
