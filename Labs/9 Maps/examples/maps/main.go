package main

import "fmt"

func main() {
	beatles := map[string]string{}

	beatles["John"] = "guitar"
	beatles["Paul"] = "bass"
	beatles["George"] = "guitar"
	beatles["Ringo"] = "drums"

	fmt.Println(beatles)
}
