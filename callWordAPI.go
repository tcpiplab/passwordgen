package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
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

	return longestWord
}

func randomWordChain() string {
	// Generate a random number between 2 and 5
	rand.Seed(time.Now().UnixNano())
	numWords := rand.Intn(4) + 2

	// Call callWordApi() and concatenate the returned words into a string
	var buffer bytes.Buffer
	for i := 0; i < numWords; i++ {
		word := callWordApi()
		buffer.WriteString(word)
		if i != numWords-1 {
			// Add a delimiter between the words except for the last word
			delimiters := "-_=+/\\|~^$#@&*:."
			delimiter := string(delimiters[rand.Intn(len(delimiters))])
			buffer.WriteString(delimiter)
		}
	}
	return buffer.String()
}
