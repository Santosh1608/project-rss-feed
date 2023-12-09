package dynamo

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	db "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var client *db.Client

func GetClient() (*db.Client, error) {
	if client != nil {
		return client, nil
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile("admin"), config.WithRegion("ap-south-1"),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
		return nil, err
	}
	client = db.NewFromConfig(cfg)
	return client, nil
}

// func updateTable() {
// 	client, _ := GetClient()

// 	out, err := client.UpdateTable(context.Background(), &dynamodb.UpdateTableInput{

// 		TableName: aws.String("project-rss"),
// 		AttributeDefinitions: []types.AttributeDefinition{
// 			{
// 				AttributeName: aws.String("gs1pk"),
// 				AttributeType: types.ScalarAttributeTypeS,
// 			},
// 			{
// 				AttributeName: aws.String("gs1sk"),
// 				AttributeType: types.ScalarAttributeTypeS,
// 			},
// 		},
// 		GlobalSecondaryIndexUpdates: []types.GlobalSecondaryIndexUpdate{
// 			types.GlobalSecondaryIndexUpdate{
// 				Create: &types.CreateGlobalSecondaryIndexAction{
// 					IndexName: aws.String("gsi1"),
// 					KeySchema: []types.KeySchemaElement{
// 						types.KeySchemaElement{
// 							AttributeName: aws.String("gs1pk"),
// 							KeyType:       types.KeyTypeHash,
// 						},
// 						types.KeySchemaElement{
// 							AttributeName: aws.String("gs1sk"),
// 							KeyType:       types.KeyTypeRange,
// 						},
// 					},
// 					Projection: &types.Projection{
// 						ProjectionType: types.ProjectionTypeAll,
// 					},
// 					ProvisionedThroughput: &types.ProvisionedThroughput{
// 						ReadCapacityUnits:  aws.Int64(3),
// 						WriteCapacityUnits: aws.Int64(3),
// 					},
// 				},
// 			},
// 		},
// 	})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(out)
// }

// func main() {
// 	updateTable()
// }
