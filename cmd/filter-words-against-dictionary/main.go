package main

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type inputOutput struct {
	Words []string `json:"words"`
}

var dictionary = make(map[string]bool)

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = "eu-north-1"
		return nil
	})
	if err != nil {
		panic(err)
	}

	svc := dynamodb.NewFromConfig(cfg)

	var startKey map[string]types.AttributeValue

	for {
		out, err := svc.Scan(context.TODO(), &dynamodb.ScanInput{
			ExclusiveStartKey: startKey,
			TableName:         aws.String("dictionary"),
		})
		if err != nil {
			panic(err)
		}

		startKey = out.LastEvaluatedKey

		for _, item := range out.Items {
			if value := item["words"]; value != nil {
				val := value.(*types.AttributeValueMemberSS)

				for _, word := range val.Value {
					dictionary[strings.ToLower(word)] = true
				}
			}
		}

		if len(out.LastEvaluatedKey) == 0 {
			break
		}
	}
}

func handleRequest(ctx context.Context, in inputOutput) (out inputOutput, err error) {
	out.Words = make([]string, 0)

	for _, word := range in.Words {
		if dictionary[word] {
			out.Words = append(out.Words, word)
		}
	}

	return
}

func main() {
	lambda.Start(handleRequest)
}
