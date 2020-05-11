package main

import (
	"fmt"
	"math"
)

// section: code
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	c := Circle{Radius: 15}
	fmt.Println(c.Area())
}

// section: code

/*
// section: output
706.8583470577034
// section: output
*/
