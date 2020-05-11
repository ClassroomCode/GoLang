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

	fmt.Println(paint(c))
}

// section: paint
func paint(c Circle) string {
	paint := .5 * c.Area()
	return fmt.Sprintf("you need %.2f gallons of paint", paint)
}

// section: paint

// section: code

/*
// section: output
706.8583470577034
you need 353.43 gallons of paint
// section: output
*/
