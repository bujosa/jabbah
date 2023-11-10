package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/memorydb"
	"github.com/go-redis/redis/v8"
)

func GetRedisClient() *redis.Client {
	LoadEnviroment()

	environment := os.Getenv("ENVIRONMENT")
	if environment == "" {
		environment = "development"
	}

	if environment == "development" {
		return getRedisClientDevelopment()
	}

	return getRedisClientProduction()
}

func getRedisClientDevelopment() *redis.Client {
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0, // Use the default database
	})

	return redisClient
}

func getRedisClientProduction() *redis.Client {
	cfg := GetAwsConfig()

	client := memorydb.NewFromConfig(cfg)

	// Get the endpoint of your Amazon MemoryDB for Redis cluster
	output, err := client.DescribeClusters(context.Background(), &memorydb.DescribeClustersInput{})
	if err != nil {
		log.Fatalf("Failed to describe clusters: %v", err)
	}

	clusterEndpoint := output.Clusters[0].ClusterEndpoint
	endpointString := *clusterEndpoint.Address
	endpointString += ":" + fmt.Sprint(clusterEndpoint.Port)

	// Create a new Redis client using the cluster endpoint
	redisClient := redis.NewClient(&redis.Options{
		Addr: endpointString, // Replace this with the endpoint of your Amazon MemoryDB for Redis cluster
		DB:   0,              // Use the default database
	})

	return redisClient
}
