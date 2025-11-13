package main

import "context"

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "password123")
	println("Token from context:", ctx.Value("token").(string))
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// Simulate booking a hotel using the token from context
	token := ctx.Value("token").(string)
	println("Booking hotel with token:", token)
}
