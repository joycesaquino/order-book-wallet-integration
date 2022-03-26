package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joycesaquino/order-book-wallet-integration/pkg/types"
)

func main() {
	lambda.Start(Handler)
}

func Handler(ctx context.Context, event events.SQSEvent) error {

	var orders types.Orders
	for _, record := range event.Records {
		if err := json.Unmarshal([]byte(record.Body), &orders); err != nil {
			return err
		}

	}
	return nil

}
