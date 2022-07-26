package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	const dictionary = "scrabble_dictionary.txt"
	const wordlist = "words.shakespeare.txt"

	words, err := readWords(dictionary)
	if err != nil {
		log.Printf("couldn't read words from file: %s, error: %s", dictionary, err.Error())
	}

	if err = uploadDictionary(words); err != nil {
		log.Printf("couldn't upload dictionary words from file: %s, error: %s", dictionary, err.Error())
	}

	if words, err = readWords(wordlist); err != nil {
		log.Printf("couldn't read words from file: %s, error: %s", wordlist, err.Error())
	}

	if err = uploadShakespeare(words); err != nil {
		log.Printf("couldn't upload shakespeare words from file: %s, error: %s", wordlist, err.Error())
	}
}

func readWords(filepath string) (words []string, err error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		err = fmt.Errorf("couldn't read file: %s, error: %w", filepath, err)
		return
	}

	words = strings.Split(string(content), "\n")
	return
}

func uploadShakespeare(words []string) error {
	return uploadWords("/shakespeare", words)
}

func uploadDictionary(words []string) error {
	return uploadWords("/dictionary", words)
}

func uploadWords(path string, words []string) error {
	const pageSize = 100

	for pageNo := 0; pageNo*pageSize < len(words); pageNo++ {
		startOfPage := pageNo * pageSize
		endOfPage := startOfPage + pageSize

		if endOfPage > len(words) {
			endOfPage = len(words)
		}

		if err := upload(path, pageNo, words[startOfPage:endOfPage]); err != nil {
			return fmt.Errorf("couldn't upload page: %d: %w", pageNo, err)
		}
	}

	return nil
}

func upload(path string, pageNo int, words []string) error {
	const baseURL = "https://5d178ebno6.execute-api.eu-north-1.amazonaws.com/alpha"

	requestBody := struct {
		Page  int      `json:"page"`
		Words []string `json:"words"`
	}{
		Page:  pageNo,
		Words: words,
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("couldn't marshal request body: %w", err)
	}

	_, err = http.Post(baseURL+path, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("couldn't post request: %w", err)
	}

	return nil
}
