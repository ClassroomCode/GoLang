package main

import (
	"fmt"

	"github.com/gopherguides/foo"
)

func main() {
	u := foo.User{
		First: "Rob",
		Last:  "Pike",
	}

	fmt.Println(u)
}
