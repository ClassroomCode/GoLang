package main

import "fmt"

func main() {
	beatles := map[string]string{}

	beatles["Paul"] = "bass"

	paul := beatles["Paul"]
	fmt.Println(paul) // bass
}
