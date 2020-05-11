package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"golang.org/x/sync/errgroup"
)

// query holds the found channel
// the word to query and the text
// of the book
type query struct {
	found chan string
	word  string
	text  string
}

func main() {
	hamlet, err := download()
	if err != nil {
		log.Fatal(err)
	}

	words := os.Args[1:]
	if len(words) == 0 {
		words = []string{"Hamlet", "Mark", "King", "Queen"}
	}

	// create a new context that timeouts after 5 seconds if
	// it hasn't found 50 results

	wg, ctx := errgroup.WithContext(ctx)

	found := make(chan string)

	for _, word := range words {
		func(word string) {
			wg.Go(func() error {
				// add some sleep for randomness
				time.Sleep(sleep())
				// call search with a context that contains the query
				// information for that word
				return search(ctx)
			})
		}(word)
	}

	go watch(ctx, found, cancel)

	if err := wg.Wait(); err != nil {
		cancel()
		log.Fatal(err)
	}
}

func watch(ctx context.Context, found chan string, cancel context.CancelFunc) {
	var count int
	for {
		select {
		case msg := <-found:
			// * print out the line number the word was found, and the text of that line:
			//   `3545: QUEEN GERTRUDE O Hamlet, speak no more`
			// * stop when it finds 50 results across all of the words
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				if err != context.Canceled {
					log.Fatal(err)
				}
			}
			return
		}
	}
}

// In a real world app you wouldn't pass the query struct via the context,
// but rather as a seperate argument.
func search(ctx context.Context) error {
	// search for a word in the text
	// if found write it to the found channel
}

func download() (string, error) {
	const hamlet = "http://hamlet.gopherguides.com"

	res, err := http.Get(hamlet)
	if err != nil {
		return "", err
	}
	if res.StatusCode != 200 {
		return "", fmt.Errorf("couldn't download %s (%d)", hamlet, res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func sleep() time.Duration {
	d := time.Duration(rand.Intn(5))
	return d * time.Microsecond
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
