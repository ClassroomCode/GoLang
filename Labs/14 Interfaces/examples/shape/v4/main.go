package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func main() {
	c := Circle{Radius: 15}
	fmt.Println(c.Area())

	fmt.Println(paint(c))

	r := Rectangle{Height: 20, Width: 30}
	fmt.Println(r.Area())

	// section: error
	fmt.Println(paint(r))
	// section: error
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
cannot use r (type Rectangle) as type Circle in argument to paint
// section: output
*/
