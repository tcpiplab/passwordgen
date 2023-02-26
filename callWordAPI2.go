package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Word struct {
	Word  string `json:"word"`
	Score int    `json:"score"`
}

func main() {
	// Set up the HTTP client and request
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://wordsapiv1.p.rapidapi.com/words?random=true", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set any headers or query parameters as needed
	req.Header.Add("X-RapidAPI-Host", "wordsapiv1.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", "<your-rapidapi-key>")

	// Send the request and read the response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var words []Word
	err = json.NewDecoder(resp.Body).Decode(&words)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Print the response data
	for _, word := range words {
		fmt.Printf("Word: %s\nScore: %d\n\n", word.Word, word.Score)
	}
}
