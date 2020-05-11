package main

import (
	"fmt"
)

func main() {
	fmt.Println(MeaningOfLife())
}

func MeaningOfLife() (rc int) {
	defer func() { rc++ }()
	return 41
}
