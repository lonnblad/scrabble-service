package score

func Calculate(word string) (score int) {
	letterDistribution := scrabbleLetterDistribution()
	letterScores := letterScores()

	for _, char := range word {
		if letterDistribution[char] > 0 {
			score += letterScores[char]
		}

		letterDistribution[char]--
	}

	return
}

func letterScores() map[rune]int {
	return map[rune]int{
		'a': 1,
		'b': 3,
		'c': 3,
		'd': 2,
		'e': 1,
		'f': 4,
		'g': 2,
		'h': 4,
		'i': 1,
		'j': 8,
		'k': 5,
		'l': 1,
		'm': 3,
		'n': 1,
		'o': 1,
		'p': 3,
		'q': 10,
		'r': 1,
		's': 1,
		't': 1,
		'v': 1,
		'w': 4,
		'u': 4,
		'x': 8,
		'y': 4,
		'z': 10,
	}
}

func scrabbleLetterDistribution() map[rune]int {
	return map[rune]int{
		'a': 9,
		'b': 2,
		'c': 2,
		'd': 1,
		'e': 12,
		'f': 2,
		'g': 3,
		'h': 2,
		'i': 9,
		'j': 1,
		'k': 1,
		'l': 4,
		'm': 2,
		'n': 6,
		'o': 8,
		'p': 2,
		'q': 1,
		'r': 6,
		's': 4,
		't': 6,
		'v': 4,
		'w': 2,
		'u': 2,
		'x': 1,
		'y': 2,
		'z': 1,
	}
}
