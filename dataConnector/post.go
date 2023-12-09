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

func CreatePost(post *models.Post) (*models.Post, error) {
	id := uuid.NewString()

	resp, err := GetAllPostsByFeedId(post.FeedId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	for _, value := range resp {
		if value.Title == post.Title {
			return nil, fmt.Errorf("post already exists for this feedId")
		}
	}
	pk := "feed_" + post.FeedId + "#posts"
	sk := "post_" + id

	data := models.Post{
		Pk:     pk,
		Sk:     sk,
		Id:     id,
		Title:  post.Title,
		FeedId: post.FeedId,
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

func GetAllPostsByFeedId(feedId string) ([]*models.Post, error) {
	db, err := dynamo.GetClient()
	table := os.Getenv("DYNAMODB_TABLE_NAME")

	if err != nil {
		fmt.Println("Error connecting to client")
		return nil, err
	}

	keys := map[string]string{
		"pk": "feed_" + feedId + "#posts",
	}

	key, err := attributevalue.MarshalMap(keys)

	if err != nil {
		return nil, err
	}

	data, err := db.Query(context.Background(), &dynamodb.QueryInput{
		TableName: aws.String(table),
		KeyConditions: map[string]types.Condition{
			"pk": {
				ComparisonOperator: types.ComparisonOperator(*aws.String("EQ")),
				AttributeValueList: []types.AttributeValue{key["pk"]},
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if len(data.Items) == 0 {
		return nil, fmt.Errorf("no posts found")
	}
	var items = []*models.Post{}
	attributevalue.UnmarshalListOfMaps(data.Items, &items)

	return items, nil
}

func GetPostByFeedId(feedId, postId string) (*models.Post, error) {
	db, err := dynamo.GetClient()
	// table := os.Getenv("DYNAMODB_TABLE_NAME")
	fmt.Printf("%+v", db)
	if err != nil {
		fmt.Println("Error connecting to client")
		return nil, err
	}

	// keys := map[string]string{
	// 	"pk": "feed_" + feedId + "#posts",
	// 	"sk": "post_" + postId,
	// }

	keys := map[string]string{
		"pk": "user_a41d84e0-c596-4efe-b3d5-89aced8bd9b8",
		"sk": "user_a41d84e0-c596-4efe-b3d5-89aced8bd9b8",
	}

	key, err := attributevalue.MarshalMap(keys)

	if err != nil {
		return nil, err
	}

	data, err := db.GetItem(context.Background(), &dynamodb.GetItemInput{
		TableName: aws.String("project-rss"),
		Key:       key,
	})

	fmt.Println("DATADATADATA", data)

	if err != nil {
		return nil, err
	}

	if data.Item == nil {
		return nil, fmt.Errorf("no post found")
	}
	var item = models.Post{}
	attributevalue.UnmarshalMap(data.Item, &item)
	return &item, nil
}
