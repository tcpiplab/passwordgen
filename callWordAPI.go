package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

type Words struct {
	Word  string `json:"word"`
	Score int    `json:"score"`
}

func callWordApi() string {
	// Set up the HTTP client and request
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.datamuse.com/sug", nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	searchString := randomWordStem()
	//searchString := "del"

	// Set any headers or query parameters as needed
	//req.Header.Add("Authorization", "Bearer <token>")
	q := req.URL.Query()
	q.Add("s", searchString)
	q.Add("max", "10")
	req.URL.RawQuery = q.Encode()

	// Send the request and read the response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Parse the JSON response
	var words []Words
	err = json.NewDecoder(resp.Body).Decode(&words)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	longestWord := ""
	for _, word := range words {
		if len(word.Word) > len(longestWord) {
			longestWord = word.Word
		}
	}

	output := strings.ReplaceAll(longestWord, " ", "_")

	return output
}

func randomWordChain(requestedPasswordLength int) string {
	//// Generate a random number between 2 and 6
	//rand.Seed(time.Now().UnixNano())
	//numWords := rand.Intn(5) + 2

	// Call callWordApi() and concatenate the returned words into a string
	var buffer bytes.Buffer

	// Choose a single delimiter to place between the words
	delimiters := "-_=+/\\|~^$#@&*:."
	delimiter := string(delimiters[rand.Intn(len(delimiters))])

	var word string

	for i := 0; i < requestedPasswordLength; i += len(word) {
		word = callWordApi()

		if len(word) > 2 {

			buffer.WriteString(word)

			if i != requestedPasswordLength {
				// Add a delimiter between the words except for the last word
				if i != requestedPasswordLength-1 {

					buffer.WriteString(delimiter)
				}
			}
		}
	}

	// Replace spaces with an underscore character
	output := strings.ReplaceAll(buffer.String(), " ", "_")

	// Truncate the resulting word-chain password to the specified length
	// by removing characters from the right side
	if len(output) > requestedPasswordLength {

		output = strings.TrimSpace(output[:requestedPasswordLength])
	}

	return output
}
