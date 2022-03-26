package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"order-book-wallet-integration/pkg"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, event events.SQSEvent) error {

	var orders pkg.Orders
	for _, record := range event.Records {
		if err := json.Unmarshal([]byte(record.Body), &orders); err != nil {
			return err
		}

	}
	return nil

}
