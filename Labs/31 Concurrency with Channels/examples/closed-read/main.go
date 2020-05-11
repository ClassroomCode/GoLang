package main

import "fmt"

func main() {
	c := make(chan int)
	go count(c)
	for {
		i, ok := <-c
		if !ok {
			fmt.Println("closed channel")
			return
		}
		fmt.Printf("read %d from channel\n", i)
	}
}

func count(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
