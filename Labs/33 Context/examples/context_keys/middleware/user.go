package middleware

import (
	"context"
	"fmt"
)

// section: authenticate
func Authenticate(ctx context.Context) context.Context {
	return context.WithValue(ctx, "id", 42)
}

// section: authenticate

func PrintUserID(ctx context.Context) {
	userID := ctx.Value("id")
	fmt.Println("User ID:", userID)
}
