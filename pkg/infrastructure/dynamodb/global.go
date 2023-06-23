package dynamodb

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

var (
	dbClient *dynamodb.Client
)

func SetDB(client *dynamodb.Client) {
	dbClient = client
}

func GetDB() *dynamodb.Client {
	return dbClient
}
