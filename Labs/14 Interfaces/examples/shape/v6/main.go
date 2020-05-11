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

// section: string
func (c Circle) String() string {
	return fmt.Sprintf("Circle{Radius: %.2f}", c.Radius)
}

// section: string

type Rectangle struct {
	Height float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// section: string
func (r Rectangle) String() string {
	return fmt.Sprintf("Rectangle{Height: %.2f, Width: %.2f}", r.Height, r.Width)
}

// section: string

type Sizer interface {
	Area() float64
}

func main() {
	c := Circle{Radius: 15}
	fmt.Println(c)
	fmt.Println(c.Area())

	fmt.Println(paint(c))

	r := Rectangle{Height: 20, Width: 30}
	fmt.Println(r)
	fmt.Println(r.Area())

	fmt.Println(paint(r))
}

func paint(s Sizer) string {
	paint := .5 * s.Area()
	return fmt.Sprintf("you need %.2f gallons of paint", paint)
}

// section: code

/*
// section: output
Circle{Radius: 15.00}
706.8583470577034
you need 353.43 gallons of paint
Rectangle{Height: 20.00, Width: 30.00}
600
you need 300.00 gallons of paint
// section: output
*/
