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

type (
	User struct {
		Email  string  `json:"email"`
		Orders []Order `json:"orders"`
	}

	Order struct {
		Time  string `json:"time"`
		Items string `json:"items"`
	}
)

type Orders struct {
	Date     string   `json:"date"`
	Orderers []string `json:"orderers"`
}

func main() {
	// user, err := getUser("zach.taylor@avant.com")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%v\n", user)

	// newUser, err := createUser("notzach.taylor@avant.com")

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(newUser)
	// updateUser()
	getTodaysOrders()
	users, _ := getUsers()
	fmt.Println(users)
}

func getUser(email string) (User, error) {
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
	})

	if err != nil {
		return User{}, err
	}

	user := User{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &user)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	return user, nil
}

func updateUser() {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":order": {
				L: []*dynamodb.AttributeValue{
					{
						M: map[string]*dynamodb.AttributeValue{
							"ordered_at": {
								S: aws.String(time.Now().Format(time.RFC3339))},
						},
					},
				},
			},
		},
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String("zach.taylor@avant.com"),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("set orders = list_append(orders, :order)"),
	}

	_, err := db.UpdateItem(input)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Successfully updated 'zach.taylor@avant.com'")
}

func getUsers() ([]string, error) {
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

func getTodaysOrders() {
	result, err := db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("fooda_orders"),
		Key: map[string]*dynamodb.AttributeValue{
			"date": {
				S: aws.String("2014-11-15"),
			},
		},
	})

	if err != nil {
		panic(err)
	}

	orders := Orders{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &orders)

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	fmt.Printf("%#v\n", orders)
}

func createUser(email string) (User, error) {
	user := User{
		Email: email,
		Orders: []Order{
			Order{Time: "2014-11-16", Items: "Some more items"},
		},
	}

	item, err := dynamodbattribute.MarshalMap(user)

	if err != nil {
		return user, err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(table),
		Item:      item,
	}

	_, err = db.PutItem(input)

	return user, err
}

/*
When I receive an order email
store date & email

At 9:45am each morning, notify each user if they haven't ordered
  - need to be able to retrieve all users
	- need to be able to retrieve all orders for today
*/
