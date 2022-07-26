package main

import (
	"context"
	"sort"
	"strconv"

	"github.com/aws/aws-lambda-go/lambda"
)

type input struct {
	Word  string `json:"word"`
	Score int    `json:"score"`
}

type output struct {
	Word  string `json:"word"`
	Score string `json:"score"`
}

type ByScore []input

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score < a[j].Score }

func handleRequest(_ context.Context, in []input) (out []output, err error) {
	if len(in) == 0 {
		return make([]output, 0), nil
	}

	sort.Sort(ByScore(in))

	for _, word := range in {
		out = append(out, output{
			Word:  word.Word,
			Score: strconv.Itoa(word.Score),
		})
	}

	return
}

func main() {
	lambda.Start(handleRequest)
}
