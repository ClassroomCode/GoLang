package race_test

import (
	"fmt"
	"sync"
	"testing"
)

func TestRace(t *testing.T) {
	var data = map[int]int{}

	get := func(i int) int {
		return data[i]
	}

	set := func(i int) {
		data[i] = i
	}

	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(2)
		go func(i int) {
			defer wg.Done()
			set(i)
		}(i)
		go func(i int) {
			defer wg.Done()
			fmt.Println(get(i))
		}(i)
	}
	wg.Wait()

}
