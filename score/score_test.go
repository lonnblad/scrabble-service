package score_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lonnblad/scrabble-service/score"
)

type calculateTest struct {
	name  string
	word  string
	score int
}

var tests = []calculateTest{
	{
		name:  "abc",
		word:  "abc",
		score: 7,
	},
	{
		name:  "x_y",
		word:  "xxy",
		score: 12,
	},
}

func TestCalculate(t *testing.T) {
	for _, test := range tests {
		test := test

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			actual := score.Calculate(test.word)
			assert.Equal(t, test.score, actual)
		})
	}
}
