package main

import (
	"fmt"
	"time"
)

func main() {
	logChan := make(chan string, 5)
	quit := make(chan struct{})

	// Logger reading fro the logChan
	go func() {
		for msg := range logChan {
			fmt.Println("delivered message: ", msg)
			// create some back pressure
			time.Sleep(10 * time.Millisecond)
		}
		close(quit)
	}()

	// simulate some load on the logger
	go func() {
		for i := 0; i < 20; i++ {
			log(i, logChan)
		}
		// this will close the logger
		close(logChan)
	}()

	<-quit

}

func log(i int, c chan string) {
	select {
	// write the message to the logger
	case c <- fmt.Sprintf("this is log message %d", i):
	default:
		fmt.Printf("log channel buffer full: dropped message due to back pressure %d\n", i)
	}
}
