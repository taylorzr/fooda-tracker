package main

import (
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
	Source        string        `json:"source"`
	Headers       []Header      `json:"headers"`
	CommonHeaders CommonHeaders `json:"commonHeaders"`
}

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CommonHeaders struct {
	From        []string `json:"from"`
	To          []string `json:"to"`
	DeliveredTo []string `json:"Delivered-To"`
	Subject     string   `json:"subject"`
}

func extractTime(email string) (time.Time, error) {
	r := regexp.MustCompile(`When\s+\w+, (\w{3}) (\d{1,2})\w{2}\s+BETWEEN .*\s+Your Order`)

	match := r.FindStringSubmatch(email)

	if len(match) == 0 {
		return time.Time{}, fmt.Errorf("Unable to match stuff")
	}

	timeString := fmt.Sprintf("%s %s %d 12:00", match[1], match[2], time.Now().Year())

	// Reference time: Mon Jan 2 15:04:05 -0700 MST 2006
	t, err := time.Parse("Jan 2 2006 15:04", timeString)

	if err != nil {
		fmt.Println("Unable to parse string")
		return time.Time{}, err
	}

	return t, nil
}

type Order struct {
	Email string
	Time  time.Time
}

func something(data string) (*Order, error) {
	email := Email{}

	err := json.Unmarshal([]byte(data), &email)

	if err != nil {
		fmt.Println("Unable to unmarshal message")
		return nil, err
	}

	fmt.Printf("From: %s\n", email.Mail.Source)

	orderTime, err := extractTime(email.Content)

	if err != nil {
		return nil, err
	}

	fmt.Printf("Time: %s\n", orderTime)

	db.SaveOrder(email.Mail.Source, orderTime)

	fmt.Println("Headers")
	fmt.Printf("%#v\n", email.Mail.Headers)
	fmt.Println("end")

	fmt.Println("Common headers")
	fmt.Printf("%#v\n", email.Mail.CommonHeaders)
	fmt.Println("end")

	return &Order{
		Email: email.Mail.Source,
		Time:  orderTime,
	}, nil
}

func main() {
	lambda.Start(handleSNS)
}

func handleSNS(snsEvent events.SNSEvent) error {
	fmt.Printf("Processing %d events\n", len(snsEvent.Records))

	for _, record := range snsEvent.Records {
		_, err := something(record.SNS.Message)

		if err != nil {
			return err
		}
	}

	return nil
}
