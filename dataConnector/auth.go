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
	"github.com/santosh1608/project-rss/requests"
)

func Register(user *requests.Register) (*models.User, error) {
	id := uuid.NewString()
	pk := "user_" + id
	sk := "user_" + id
	gs1pk := "user_" + user.Email
	gs1sk := user.Email

	data := models.User{
		Pk:       pk,
		Sk:       sk,
		Gs1Pk:    gs1pk,
		Gs1Sk:    gs1sk,
		Id:       id,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
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
