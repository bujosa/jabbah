package config

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

func GetAwsConfig() aws.Config {
	// Load the SDK's configuration from environment and shared config files = ~/.aws/config and ~/.aws/credentials
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	fmt.Println("Region: ", cfg.Region)

	return cfg
}
