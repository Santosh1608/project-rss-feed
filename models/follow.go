package models

type Follow struct {
	Pk     string `dynamodbav:"pk"`
	Sk     string `dynamodbav:"sk"`
	Id     string `dynamodbav:"id"`
	FeedId string `dynamodbav:"feedId"`
	UserId string `dynamodbav:"userId"`
}
