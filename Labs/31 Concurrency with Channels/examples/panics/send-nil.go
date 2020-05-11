package main

func main() {
	// section: code
	var c3 chan int
	c3 <- 1 // panic: send on a nil channel
	// section: code
}
