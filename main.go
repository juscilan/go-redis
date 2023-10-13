package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	// Initialize a Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6380", // Redis server address
		Password: "",               // No password
		DB:       0,                // Default DB
	})

	// Close the client when the program finishes
	defer client.Close()

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Set a key-value pair in Redis
	err := client.Set(ctx, "mykey", "Hello, Redis!", 0).Err()
	if err != nil {
		fmt.Println("Error setting key:", err)
		return
	}
	fmt.Println("Key set successfully!")

	// Retrieve the value from Redis
	val, err := client.Get(ctx, "mykey").Result()
	if err != nil {
		fmt.Println("Error getting key:", err)
		return
	}
	fmt.Printf("Value: %s\n", val)
}
