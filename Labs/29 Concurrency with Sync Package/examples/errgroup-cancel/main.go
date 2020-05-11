package main

import (
	"context"
	"fmt"
	"log"

	"github.com/markbates/randx"
	"golang.org/x/sync/errgroup"
)

// section: main
func main() {
	ctx := context.Background()
	wg, ctx := errgroup.WithContext(ctx)

	go func() {
		// wait for errgroup to return an error
		// and cancel its context
		<-ctx.Done()
		fmt.Println("error returned from errgroup", ctx.Err())
	}()

	for i := 0; i < 10; i++ {
		x := i
		wg.Go(func() error {
			return doWork(ctx, x)
		})
	}

	// wait for errgroup to finish/error
	if err := wg.Wait(); err != nil {
		log.Fatal(err)
	}
}

// section: main

func doWork(ctx context.Context, x int) error {
	for {
		select {
		case <-ctx.Done():
			// context was canceled
			return nil
		default:
			fmt.Printf("processing %d\n", x)
			// sleep randomly
			randx.RandomSleep(50)
			// if error, return it to the errgroup
			if err := randx.RandomError(42); err != nil {
				return err
			}
		}
	}
}
