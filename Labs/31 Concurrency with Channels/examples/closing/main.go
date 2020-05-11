package main

import "fmt"

func echo(s string, c chan string) {
	for i := 0; i < cap(c); i++ {
		// write the string down the channel
		c <- s
	}
	close(c)
}

func main() {
	// make a channel with a capacity of 10
	c := make(chan string, 10)

	// launch the goroutine
	go echo("hello!", c)

	for s := range c {
		fmt.Println(s)
	}
}
