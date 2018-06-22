package db

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	db    *dynamodb.DynamoDB
	table = "fooda"
)

func init() {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	db = dynamodb.New(sess)
}

func GetUsers() ([]string, error) {
	r, err := db.Scan(&dynamodb.ScanInput{
		TableName: aws.String("fooda_users"),
	})

	if err != nil {
		return nil, err
	}

	users := []string{}

	for _, item := range r.Items {
		users = append(users, *item["email"].S)
	}

	return users, nil
}

func GetTodaysOrders() ([]string, error) {
	today := time.Now().Format("2006-01-02")

	fmt.Println(today)

	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("fooda_orders"),
		Key: map[string]*dynamodb.AttributeValue{
			"date": {
				S: aws.String(today),
			},
		},
	})

	if err != nil {
		return nil, err
	}

	if len(result.Item) == 0 {
		return []string{}, nil
	}

	orderers := []string{}

	for _, orderer := range result.Item["orderers"].SS {
		orderers = append(orderers, *orderer)
	}

	return orderers, nil
}

type Order struct {
	Date     string   `json:"date"`
	Orderers []string `json:"orderers" dynamodbav:"orderers,stringset"`
}

func SaveOrder(email string, t time.Time) error {
	order := Order{
		Date:     t.Format("2006-01-02"),
		Orderers: []string{email},
	}

	item, err := dynamodbattribute.MarshalMap(order)

	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("fooda_orders"),
	}

	_, err = db.PutItem(input)

	return err
}
