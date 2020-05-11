package main

import (
	"context"
	"errors"
	"fmt"
)

// section: type
type key int

// section: type

// section: const
const userKey key = 0

// section: const

type User struct {
	ID   int
	Name string
}

// section: with-user
func WithUser(ctx context.Context, u *User) context.Context {
	return context.WithValue(ctx, userKey, u)
}

// section: with-user

// section: from-context
func FromContext(ctx context.Context) (*User, error) {
	u, ok := ctx.Value(userKey).(*User)
	if !ok {
		return nil, errors.New("user not found")
	}
	return u, nil
}

// section: from-context

// section: main
func main() {
	u := &User{
		ID:   0,
		Name: "Tim",
	}

	ctx := WithUser(context.Background(), u)
	printUser(ctx)

}

func printUser(ctx context.Context) {
	u, err := FromContext(ctx)
	if err != nil {
		fmt.Println("err getting user:", err)
		return
	}
	fmt.Println("user:", u)
}

// section: main
