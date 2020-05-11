package main

import (
	"log"
	"training/cmd"
)

// section: main
func main() {
	c := cmd.New()
	if err := c.Open(); err != nil {
		log.Fatal(c)
	}

	wait := make(chan struct{})
	<-wait

	if err := c.Close(); err != nil {
		log.Fatal(err)
	}
	log.Println("successfuly shutdown")
}

// section: main

/*
// section: output
2020/04/20 11:09:09 monitor check
2020/04/20 11:09:10 monitor check
2020/04/20 11:09:11 monitor check
^Csignal: interrupt
// section: output
*/
