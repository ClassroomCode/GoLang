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

func (u User) Slug() string {
	return url.PathEscape(fmt.Sprintf("%s-%s-%s", u.ID, u.First, u.Last))
}

// #TODO:
// Satisfy the stringer interface (https://golang.org/pkg/fmt/#Stringer)
// so that the User struct will print
// User <ID> is <First> <Last>
//
// example:
//      User 1 is Rob Pike

type Product struct {
	ID   int
	Name string
}

func (p Product) Slug() string {
	return url.PathEscape(fmt.Sprintf("%d-%s", p.ID, p.Name))
}

// #TODO: Create an interface called `Slugger` that matches the behavior of the
//  `Slug` function that both the Product and User type have declared

func main() {
	u := User{ID: 1, First: "Rob", Last: "Pike"}
	fmt.Println(u)

	p := Product{ID: 10, Name: "Plush Gopher"}

	// #TODO: Declare a slice of `Slugger` and add the previously declared user and product to the slice

	for _, s := range sluggers {
		// #TODO: Print out the `Slug` for each item

		// #TODO Write a switch statement that switches on the `type`.
		//    - When the type is a user, print out the First and Last name
		//    - When the type is a product, print out he name
	}

}
