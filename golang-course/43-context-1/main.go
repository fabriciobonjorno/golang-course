package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*6)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// Simulate booking a hotel
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking cancelled:")
		return
	case <-time.After(5 * time.Second):
		// Booking successful
		fmt.Println("Hotel Booked")
	}
}
