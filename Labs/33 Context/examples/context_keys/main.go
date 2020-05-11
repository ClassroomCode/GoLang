package main

import (
	"context"

	"github.com/gopherguides/learn/_training/fundamentals/context/src/context_keys/middleware"
	"github.com/gopherguides/learn/_training/fundamentals/context/src/context_keys/orders"
)

// section: main
func main() {
	ctx := context.Background()
	ctx = middleware.Authenticate(ctx)

	middleware.PrintUserID(ctx)

	ctx = orders.NextOrderID(ctx)
	orders.CreateOrder(ctx)

	middleware.PrintUserID(ctx)
}

// section: main
