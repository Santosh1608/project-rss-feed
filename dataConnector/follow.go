package dataConnector

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
	"github.com/santosh1608/project-rss/dynamo"
	"github.com/santosh1608/project-rss/models"
)

func FollowFeed(userId, feedId string) (*models.Follow, error) {
	id := uuid.NewString()
	pk := "user_" + userId + "#follows"
	sk := "feed_" + feedId

	data := models.Follow{
		Pk:     pk,
		Sk:     sk,
		Id:     id,
		UserId: userId,
		FeedId: feedId,
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
