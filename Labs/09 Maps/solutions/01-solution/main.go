package main

import (
	"fmt"
	"sort"
)

func main() {

	beatles := map[string]string{}

	beatles["John"] = "guitar"
	beatles["Paul"] = "bass"
	beatles["George"] = "guitar"
	beatles["Ringo"] = "drums"

	// Loop through the map and print out the key and value

	for key, value := range beatles {
		fmt.Printf("key: %q, value: %q\n", key, value)
	}

	// Look up the key `Bob` and detect that it wasn't found by printing `not found`
	if bob, ok := beatles["Bob"]; !ok {
		fmt.Println("not found")
	} else {
		fmt.Println(bob)
	}

	// Declare a slice of strings called `keys` and collect all the keys from the map.
	keys := []string{}
	for k := range beatles {
		keys = append(keys, k)
	}

	// sort the keys
	sort.Strings(keys)

	// print out the keys
	fmt.Println(keys)

}
