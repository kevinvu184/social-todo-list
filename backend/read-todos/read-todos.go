package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Todo struct {
	createdAt string
	done      bool
	name      string
	todoID    string
}

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "DUMMYIDEXAMPLE", SecretAccessKey: "DUMMYEXAMPLEKEY",
			},
		}),
	)
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}

	r := struct {
		TodoID string `dynamodbav:"todoID"`
	}{TodoID: "dc23d17e-798f-434e-b068-ab1d1121d006"}

	avs, err := attributevalue.MarshalMap(r)
	if err != nil {
		log.Fatal("failed to marshal Todo, %w", err)
	}
	log.Println(avs)

	svc := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.EndpointResolver = dynamodb.EndpointResolverFromURL("http://localhost:8000")
	})
	result, err := svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String("Todo"),
		Key:       avs,
	})
	if err != nil {
		log.Fatalf("Got error calling GetItem: %s", err)
	}
	log.Println(result.Item)
}
