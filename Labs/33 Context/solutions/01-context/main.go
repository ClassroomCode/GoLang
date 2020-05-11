package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	wg, ctx := errgroup.WithContext(ctx)

	found := make(chan string)

	for _, word := range words {
		func(word string) {
			wg.Go(func() error {
				// add some sleep for randomness
				time.Sleep(sleep())
				ctx = context.WithValue(ctx, "query", query{
					found: found,
					word:  word,
					text:  hamlet,
				})
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
			if count > 49 {
				cancel()
				return
			}
			count++
			fmt.Printf("%d) %s\n", count, msg)
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
	q, ok := ctx.Value("query").(query)
	if !ok {
		return fmt.Errorf("expected query to be of type 'query' got '%T' instead", q)
	}
	for i, line := range strings.Split(q.text, "\n") {
		select {
		case <-ctx.Done():
			return nil
		default:
			if strings.Contains(line, q.word) {
				q.found <- fmt.Sprintf("[%s] %d: %s", q.word, i+1, line)
			}
			time.Sleep(sleep())
		}
	}

	return nil
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
