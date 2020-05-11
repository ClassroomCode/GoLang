package main

import (
	"fmt"
	"log"

	"github.com/gopherguides/learn/_training/fundamentals/embedding-and-composition/src/private-interface/notify"
)

type User struct {
	ID          int
	First, Last string
}

func (u User) Notify(msg string) error {
	_, err := fmt.Printf("sending notification: %q\n", msg)
	return err
}

func main() {
	u := User{
		ID:    1,
		First: "Rob",
		Last:  "Pike",
	}
	if err := notify.Update(u); err != nil {
		log.Fatal(err)
	}
}
