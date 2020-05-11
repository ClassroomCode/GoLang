package main

func main() {
	// section: code
	c2 := make(chan int)
	close(c2)
	close(c2) // panic: closeing a channel that is already closed
	// section: code
}
