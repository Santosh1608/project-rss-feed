package models

type Feed struct {
	Pk     string `dynamodbav:"pk"`
	Sk     string `dynamodbav:"sk"`
	Gs1Pk  string `dynamodbav:"gs1pk"`
	Gs1Sk  string `dynamodbav:"gs1sk"`
	Id     string `dynamodbav:"id"`
	Name   string `dynamodbav:"name"`
	Url    string `dynamodbav:"url"`
	UserId string `dynamodbav:"userId"`
}
