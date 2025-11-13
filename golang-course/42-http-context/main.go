package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	// Simulate some work with the context
	select {
	case <-time.After(2 * time.Second):
		println("Completed work")
	case <-ctx.Done():
		println("Context cancelled:", ctx.Err().Error())
	}
}
