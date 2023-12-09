package dataConnector

import (
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/santosh1608/project-rss/dynamo"
	"github.com/santosh1608/project-rss/models"
)

func GetUserByEmail(email string) (*models.User, error) {
	client, err := dynamo.GetClient()
	table := os.Getenv("DYNAMODB_TABLE_NAME")
	var item models.User

	if err != nil {
		fmt.Println("Error connecting to client")
		return nil, err
	}

	keys := map[string]string{
		"gs1pk": "user_" + email,
	}

	key, err := attributevalue.MarshalMap(keys)

	if err != nil {
		return nil, err
	}

	data, err := client.Query(context.Background(), &dynamodb.QueryInput{
		TableName: aws.String(table),
		IndexName: aws.String("gsi1"),
		KeyConditions: map[string]types.Condition{
			"gs1pk": {
				ComparisonOperator: types.ComparisonOperator(*aws.String("EQ")),
				AttributeValueList: []types.AttributeValue{key["gs1pk"]},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if len(data.Items) == 0 {
		return nil, fmt.Errorf("GetItem: Data not found")
	}

	err = attributevalue.UnmarshalMap(data.Items[0], &item)

	if err != nil {
		return &item, fmt.Errorf("UnmarshalMap %v", err)
	}

	return &item, nil
}
