package main

import "fmt"

// Function to add two numbers
func addTwoNumbers(x, y int) int {
	sum := x + y
	return sum
}

// Function to multiply two numbers
func multiplyTwoNumbers(x, y int) int {
	product := x * y
	return product
}

// section: main
func main() {
	/*
		In this example, we're commenting out the addTwoNumbers
		function because it is failing, therefore preventing it from executing.
		Only the multiplyTwoNumbers function will run

		a := addTwoNumbers(3, 5)
		fmt.Println(a)

	*/

	m := multiplyTwoNumbers(5, 9)
	fmt.Println(m)
}

// section: main
