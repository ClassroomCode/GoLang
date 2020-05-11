package main

import (
	"fmt"
	"net/url"
)

type User struct {
	ID    int
	First string
	Last  string
}

func (u User) String() string {
	return fmt.Sprintf("User %d is %s %s", u.ID, u.First, u.Last)
}

func (u User) Slug() string {
	return url.PathEscape(fmt.Sprintf("%s-%s-%s", u.ID, u.First, u.Last))
}

type Product struct {
	ID   int
	Name string
}

func (p Product) Slug() string {
	return url.PathEscape(fmt.Sprintf("%d-%s", p.ID, p.Name))
}

// Create an interface called `Slugger` that declares a function called `Slug() string`
type Slugger interface {
	Slug() string
}

func main() {
	u := User{ID: 1, First: "Rob", Last: "Pike"}
	fmt.Println(u)

	p := Product{ID: 10, Name: "Plush Gopher"}

	sluggers := []Slugger{u, p}
	for _, s := range sluggers {
		fmt.Println(s.Slug())
		switch t := s.(type) {
		case User:
			fmt.Println(t.First, t.Last)
		case Product:
			fmt.Println(t.Name)
		}
	}
}
