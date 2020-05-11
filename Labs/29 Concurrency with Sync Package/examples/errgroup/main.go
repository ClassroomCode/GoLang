package main

import (
	"fmt"
	"log"

	"github.com/markbates/randx"
	"golang.org/x/sync/errgroup"
)

// section: main
func main() {
	var wg errgroup.Group
	for i := 0; i < 100; i++ {
		x := i
		wg.Go(func() error {
			randx.RandomSleep(50)
			fmt.Printf("processing %d\n", x)
			return randx.RandomError(42)
		})
	}
	if err := wg.Wait(); err != nil {
		log.Fatal(err)
	}
}

// section: main
