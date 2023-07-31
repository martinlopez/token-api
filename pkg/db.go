package pkg

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go/aws"
)

type DBSecrets struct {
	Username string `json:"username"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Host     string `json:"host"`
}

func GetDBSecrets() (*DBSecrets, error) {
	secretName := "dev/core"
	region := "us-east-1"

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return nil, err
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		// For a list of exceptions thrown, see
		// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
		return nil, err
	}

	// Decrypts secret using the associated KMS key.
	secretString := *result.SecretString
	var dbConfig DBSecrets
	err = json.Unmarshal([]byte(secretString), &dbConfig)
	if err != nil {
		log.Println("error unmarshalling db secrets")
		return nil, err
	}
	return &dbConfig, nil
}
