package main

import (
	"context"
	"encoding/json"
	"fmt"
	"regexp"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/taylorzr/fooda-tracker/db"
	// "github.com/taylorzr/fooda-tracker/fooda"
)

type Email struct {
	// Receipt Receipt `json:"receipt"`
	Mail    Mail   `json:"mail"`
	Content string `json:"content"`
}

type Mail struct {
	Source string `json:"source"`
}

func HandleRequest(ctx context.Context, snsEvent events.SNSEvent) {
	fmt.Printf("Processing %d events\n", len(snsEvent.Records))

	for _, record := range snsEvent.Records {
		snsRecord := record.SNS

		email := Email{}

		err := json.Unmarshal([]byte(snsRecord.Message), &email)

		if err != nil {
			fmt.Println("Unable to unmarshal message")
			panic(err)
		}

		fmt.Printf("From: %s\n", email.Mail.Source)
		t := extractTime(email.Content)
		fmt.Printf("Time: %s\n", t)

		db.SaveOrder(email.Mail.Source, t)
		// TODO: Save stuff to dynamodb

		// fmt.Printf("\n\nContent:\n%s", email.Content)
	}
}

// func parseEmail() {
// 	message, err := mail.ReadMessage(file)

// 	if err != nil {
// 		panic("Oh no")
// 	}

// 	body, err := ioutil.ReadAll(message.Body)

// 	if err != nil {
// 		panic("Oh no 2")
// 	}

// 	fmt.Println(string(body))
// }

func extractTime(email string) time.Time {
	r := regexp.MustCompile(`When\r\n\w+, (\w{3}) (\d{1,2})\w{2}\r\nBETWEEN .*\r\nYour Order\r\n(.*)\r\n`)

	match := r.FindStringSubmatch(email)

	timeString := fmt.Sprintf("%s %s %d 12:00", match[1], match[2], time.Now().Year())

	// Reference time: Mon Jan 2 15:04:05 -0700 MST 2006
	t, err := time.Parse("Jan 2 2006 15:04", timeString)

	if err != nil {
		fmt.Println("Unable to parse string")
		panic(err)
	}

	return t
}

func main() {
	lambda.Start(HandleRequest)
}
