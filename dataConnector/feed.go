package dataConnector

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/santosh1608/project-rss/dynamo"
	"github.com/santosh1608/project-rss/models"
)

func CreateFeed(feed *models.Feed) (*models.Feed, error) {
	id := uuid.NewString()
	pk := "user_" + feed.UserId + "#feeds"
	sk := "feed_" + id
	gs1pk := "feeds"
	gs1sk := "feed_" + id

	data := models.Feed{
		Pk:     pk,
		Sk:     sk,
		Gs1Pk:  gs1pk,
		Gs1Sk:  gs1sk,
		Id:     id,
		Name:   feed.Name,
		UserId: feed.UserId,
		Url:    feed.Url,
	}
	dataMarshed, err := attributevalue.MarshalMap(data)
	if err != nil {
		log.Println("Failed to Marshal item" + err.Error())
	}

	db, _ := dynamo.GetClient()
	table := os.Getenv("DYNAMODB_TABLE_NAME")

	_, err = db.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: &table,
		Item:      dataMarshed,
	})

	if err != nil {
		log.Println("Failed to Put item" + err.Error())
		return nil, err
	}
	return &data, nil
}

func GetAllFeeds() ([]*models.Feed, error) {
	db, err := dynamo.GetClient()
	table := os.Getenv("DYNAMODB_TABLE_NAME")

	if err != nil {
		fmt.Println("Error connecting to client")
		return nil, err
	}

	keys := map[string]string{
		"gs1pk": "feeds",
	}

	key, err := attributevalue.MarshalMap(keys)

	if err != nil {
		return nil, err
	}

	data, err := db.Query(context.Background(), &dynamodb.QueryInput{
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
		return nil, fmt.Errorf("no feeds found")
	}
	var items = []*models.Feed{}
	attributevalue.UnmarshalListOfMaps(data.Items, &items)

	return items, nil
}
