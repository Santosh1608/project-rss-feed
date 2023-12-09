package models

type Post struct {
	Pk     string `dynamodbav:"pk"`
	Sk     string `dynamodbav:"sk"`
	Id     string `dynamodbav:"id"`
	Title  string `dynamodbav:"title"`
	FeedId string `dynamodbav:"feedId"`
}
