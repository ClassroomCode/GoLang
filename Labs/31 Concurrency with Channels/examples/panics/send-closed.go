package main

func main() {
	// section: code
	c1 := make(chan int)
	close(c1)
	c1 <- 1 // panic: send on a closed channel
	// section: code
}
