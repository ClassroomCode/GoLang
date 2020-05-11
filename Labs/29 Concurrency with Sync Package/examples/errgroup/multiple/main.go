package main

import (
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// section: main
func main() {
	var wg errgroup.Group
	for i := 0; i < 5; i++ {
		i := i // capture the range variable
		wg.Go(func() error {
			fmt.Printf("processing %d\n", i)
			time.Sleep(10 * time.Duration(i) * time.Millisecond)
			fmt.Printf("finished %d\n", i)
			return fmt.Errorf("routine %d error'd out", i)
		})
	}
	if err := wg.Wait(); err != nil {
		fmt.Println(err)
	}
}

// section: main

/*
// section: output
processing 4
processing 2
processing 0
finished 0
processing 1
processing 3
finished 1
finished 2
finished 3
finished 4
routine 0 error'd out
// section: output
*/
