package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/taylorzr/fooda-tracker/fooda"
)

func HandleRequest(ctx context.Context) ([]string, error) {
	return fooda.Remind()
}

func main() {
	lambda.Start(HandleRequest)
}
