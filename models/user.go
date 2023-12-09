package models

type User struct {
	Pk       string `dynamodbav:"pk" json:"-"`
	Sk       string `dynamodbav:"sk" json:"-"`
	Gs1Pk    string `dynamodbav:"gs1pk" json:"-"`
	Gs1Sk    string `dynamodbav:"gs1sk" json:"-"`
	Id       string `dynamodbav:"id" json:"id"`
	Name     string `dynamodbav:"name" json:"name"`
	Email    string `dynamodbav:"email" json:"email"`
	Password string `dynamodbav:"password"`
}
