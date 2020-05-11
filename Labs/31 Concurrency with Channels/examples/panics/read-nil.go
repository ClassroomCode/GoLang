package main

func main() {
	// section: code
	var c4 chan bool
	<-c4 // panic: read from a nil channel: all goroutines are asleep
	// section: code
}
