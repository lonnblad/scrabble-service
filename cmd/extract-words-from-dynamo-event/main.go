package main

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type output struct {
	Words []string `json:"words"`
}

func handleRequest(ctx context.Context, e events.DynamoDBEvent) (out output, err error) {
	uniqueWords := make(map[string]bool)

	for _, record := range e.Records {
		if record.EventName != string(events.DynamoDBOperationTypeInsert) {
			continue
		}

		value := record.Change.NewImage["words"]
		if value.DataType() != events.DataTypeStringSet {
			for _, word := range value.StringSet() {
				uniqueWords[strings.ToLower(word)] = true
			}
		}
	}

	out.Words = make([]string, 0, len(uniqueWords))
	for word := range uniqueWords {
		out.Words = append(out.Words, word)
	}

	return
}

func main() {
	lambda.Start(handleRequest)
}
