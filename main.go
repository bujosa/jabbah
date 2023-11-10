package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/memorydb"
	"github.com/go-redis/redis/v8"
)

func main() {
	// Load the SDK's configuration from environment and shared config, and
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	fmt.Println("Region: ", cfg.Region)

	// Create a new MemoryDB client
	client := memorydb.NewFromConfig(cfg)

	// Get the endpoint of your Amazon MemoryDB for Redis cluster
	output, err := client.DescribeClusters(context.Background(), &memorydb.DescribeClustersInput{})
	if err != nil {
		log.Fatalf("Failed to describe clusters: %v", err)
	}

	clusterEndpoint := output.Clusters[0].ClusterEndpoint
	endpointString := *clusterEndpoint.Address
	endpointString += ":" + fmt.Sprint(clusterEndpoint.Port)
	fmt.Println("Endpoint: ", endpointString)

	// Create a new Redis client using the cluster endpoint
	redisClient := redis.NewClient(&redis.Options{
		Addr: endpointString, // Replace this with the endpoint of your Amazon MemoryDB for Redis cluster
		DB:   0,              // Use the default database
	})

	// Set a key-value pair
	err = redisClient.Set(context.Background(), "foo", "bar", 0).Err()
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
