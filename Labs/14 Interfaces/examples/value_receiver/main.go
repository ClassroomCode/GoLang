package main

import "fmt"

type User struct {
	First string
	Last  string
}

func (u User) String() string {
	return fmt.Sprintf("%s %s", u.First, u.Last)
}

func pretty(v fmt.Stringer) {
	fmt.Println(v.String())
}

func main() {
	u := &User{First: "Rob", Last: "Pike"}
	pretty(u)
}
