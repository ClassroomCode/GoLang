package orders

import (
	"context"
	"fmt"
)

var nextID int

// section: order-id
func NextOrderID(ctx context.Context) context.Context {
	defer func() { nextID++ }()
	return context.WithValue(ctx, "id", nextID)
}

// section: order-id

func CreateOrder(ctx context.Context) {
	id := ctx.Value("id").(int)
	fmt.Println("Order ID:", id)
}
