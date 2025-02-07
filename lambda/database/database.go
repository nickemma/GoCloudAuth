package database

// This is going to be the logic that directly communicates with
// our database

import (
	"fmt"
	"lambda-func/types"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// We are going to define a constant for our table name
const (
	TABLE_NAME = "userTable"
)

// We are going to define an interface for our UserStore
type UserStore interface {
	DoesUserExist(username string) (bool, error)
	InsertUser(user types.User) error
	GetUser(username string) (types.User, error)
}

// We are going to define a struct for our DynamoDBClient
type DynamoDBClient struct {
	databaseStore *dynamodb.DynamoDB
}

// We are going to define a constructor function for our DynamoDBClient
func NewDynamoDB() DynamoDBClient {
	dbSession := session.Must(session.NewSession())
	db := dynamodb.New(dbSession)

	return DynamoDBClient{
		databaseStore: db,
	}
}

// - checking if a user with this username already exists in our DB
func (u DynamoDBClient) DoesUserExist(username string) (bool, error) {
	result, err := u.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	})
	// - if they do, we return an error
	if err != nil {
		return true, err
	}

	if result.Item == nil {
		return false, nil
	}

	return true, nil
}

// - if they DON'T, that is when we INSERT the user into the DB
func (u DynamoDBClient) InsertUser(user types.User) error {
	// we want to create a new item to insert into the table
	item := &dynamodb.PutItemInput{
		TableName: aws.String(TABLE_NAME),
		Item: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(user.Username),
			},
			"password": {
				S: aws.String(user.PasswordHash),
			},
		},
	}

	// we want to insert the item into the table
	_, err := u.databaseStore.PutItem(item)
	if err != nil {
		return err
	}

	return nil
}

// - if they DO, we want to get the user from the DB
func (u DynamoDBClient) GetUser(username string) (types.User, error) {
	var user types.User
	// we want to get the item from the table
	result, err := u.databaseStore.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(TABLE_NAME),
		Key: map[string]*dynamodb.AttributeValue{
			"username": {
				S: aws.String(username),
			},
		},
	})
	// we want to unmarshal the item into our user struct
	if err != nil {
		return user, err
	}

	if result.Item == nil {
		return user, fmt.Errorf("user not found")
	}
	// we want to unmarshal the item into our user struct
	err = dynamodbattribute.UnmarshalMap(result.Item, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}
