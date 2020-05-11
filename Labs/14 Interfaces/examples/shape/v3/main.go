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

// section: rectangle
type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// section: rectangle

func main() {
	c := Circle{Radius: 15}
	fmt.Println(c.Area())

	fmt.Println(paint(c))

	r := Rectangle{Height: 20, Width: 30}
	fmt.Println(r.Area())
}

func paint(c Circle) string {
	paint := .5 * c.Area()
	return fmt.Sprintf("you need %.2f gallons of paint", paint)
}

// section: code

/*
// section: output
706.8583470577034
you need 353.43 gallons of paint
600
// section: output
*/
