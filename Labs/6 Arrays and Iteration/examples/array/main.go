package main

import "fmt"

func main() {
	// section: main
	names := [4]string{}
	names[0] = "John"
	names[1] = "Paul"
	names[2] = "George"
	names[3] = "Ringo"
	// section: main

	fmt.Println(names)

	i := 0
	for i < len(names) {
		if i == 3 {
			break
		}
		fmt.Println(names[i])
		i++
	}
}
