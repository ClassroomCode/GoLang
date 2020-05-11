package main

import (
	"fmt"
	"strings"
)

func main() {
	zero()
}

func zero() {
	// section: code
	counts := make(map[string]int)

	sentence := "The quick brown fox jumps over the lazy dog"

	words := strings.Fields(strings.ToLower(sentence))

	for _, w := range words {
		counts[w]++
	}
	// section: code

	fmt.Println(counts)
}

/*
// section: output
map[brown:1 dog:1 fox:1 jumps:1 lazy:1 over:1 quick:1 the:2]
// section: output
*/
