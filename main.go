package main

import (
	"context"
	"fmt"
	"jabbah/config"
	"log"
)

func main() {
	// Load Enviroment
	config.LoadEnviroment()

	// Create a new Redis client using the cluster endpoint
	redisClient := config.GetRedisClient()

	// Set a key-value pair
	err := redisClient.Set(context.Background(), "foo", "bar", 0).Err()
	if err != nil {
		log.Fatalf("failed to set key-value pair: %v", err)
	}

	// Get the value of a key
	val, err := redisClient.Get(context.Background(), "foo").Result()
	if err != nil {
		log.Fatalf("failed to get value of key: %v", err)
	}

	// Print the value
	fmt.Println(val) // "bar"

	// Close the Redis client
	err = redisClient.Close()
	if err != nil {
		log.Fatalf("failed to close Redis client: %v", err)
	}
}
