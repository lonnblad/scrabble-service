package main

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lonnblad/scrabble-service/score"
)

type input struct {
	Word string `json:"word"`
}

type output struct {
	Word  string `json:"word"`
	Score int    `json:"score"`
}

func handleRequest(_ context.Context, ev input) (out output, err error) {
	out.Word = ev.Word
	out.Score = score.Calculate(ev.Word)

	return
}

func main() {
	lambda.Start(handleRequest)
}
