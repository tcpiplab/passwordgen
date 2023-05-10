package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// CompletionCreateArgs is used for the ChatGPT API.
// It is defined here at the package level so any function can access it
type CompletionCreateArgs struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens"`
	Temperature float64 `json:"temperature"`
}

func createGrammaticalPasswordAI(nonSensicalSentence string, grammaticalAIWithNumbers bool) string {

	openaiAPIURL, apiKey := setupChatGPTAPI()

	// Check if the API key exists
	if apiKey == "" {
		log.Println("Error: API key is missing. Please set the API key and try again.")
		os.Exit(1)
	}
	// Continue execution if the environment variable exists

	var promptSentence string

	// Declare a variable of type CompletionCreateArgs that we'll use below
	var chatGPTRequestData CompletionCreateArgs

	if grammaticalAIWithNumbers == true {

		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randomInteger := r.Intn(101)
		randomIntegerString := strconv.Itoa(randomInteger)

		// Ask to add a number to the sentence it is rewriting
		// Note that the nonsensical sentence was actually already rewritten by AI once
		promptSentence = "Please rewrite this sentence so that it is less nonsensical and more coherent: " +
			"'" + nonSensicalSentence + "'" + "Also, please add in the digit " + randomIntegerString + " to quantify the subject of the sentence."

		// The sentences will be longer, so we need to allocate more tokens
		chatGPTRequestData = CompletionCreateArgs{
			Model:  "text-davinci-003",
			Prompt: promptSentence,
			// Any amount of tokens < 12 will truncate some sentences.
			MaxTokens: 16,
			// The best outcomes seem to be with temperature set to 0.
			Temperature: 0,
		}

	} else if grammaticalAIWithNumbers == false {

		// Just ask it to rewrite the sentence
		promptSentence = "Change the subject in the following nonsensical sentence so that it makes more sense. " +
			"Change the adverb, adjective, noun, or verb if they don't sound like they belong together: '" +
			nonSensicalSentence + "'"

		chatGPTRequestData = CompletionCreateArgs{
			Model:  "text-davinci-003",
			Prompt: promptSentence,
			// Any amount of tokens < 12 will truncate some sentences.
			MaxTokens: 12,
			// The best outcomes seem to be with temperature set to 0.
			Temperature: 0,
		}

	}

	// Make the actual API call
	chatGPTResponseBody, errorString, apiRequestError := makeChatGPTAPIRequest(chatGPTRequestData, openaiAPIURL, apiKey)

	// If the API call returned and error, return the error string
	if apiRequestError {
		return errorString
	}

	rewrittenSentence := extractGPTJson(string(chatGPTResponseBody))

	// Remove any surrounding single quotes. This happens sometimes.
	rewrittenSentence = strings.Trim(rewrittenSentence, "'")

	// If the rewrittenSentence is missing a terminating punctuation mark
	// then add a trailing period character.
	if !strings.HasSuffix(rewrittenSentence, ".") && !strings.HasSuffix(rewrittenSentence, "?") {
		rewrittenSentence += "."
	}

	//fmt.Println(nonSensicalSentence)
	//fmt.Println(rewrittenSentence)

	return rewrittenSentence
}

func makeChatGPTAPIRequest(chatGPTRequestData CompletionCreateArgs, openaiAPIURL string, apiKey string) ([]byte, string, bool) {

	// Convert the struct data into JSON
	chatGPTRequestJSON, err := json.Marshal(chatGPTRequestData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return nil, "Error marshaling JSON", true
	}

	// Create a new HTTP request using our JSON as the POST body
	chatGPTRequest, err := http.NewRequest("POST", openaiAPIURL, bytes.NewBuffer(chatGPTRequestJSON))

	if err != nil {
		fmt.Println("Error creating HTTP request:", err)
		return nil, "Error creating HTTP request", true
	}

	// Set up headers for content type and authorization
	chatGPTRequest.Header.Set("Content-Type", "application/json")
	chatGPTRequest.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	//improvedProgressBar(1)

	// Send the HTTP request
	httpClient := &http.Client{}
	resp, err := httpClient.Do(chatGPTRequest)
	if err != nil {
		fmt.Println("Error making HTTP request:", err)
		return nil, "Error making HTTP request", true
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Problem closing HTTP response chatGPTResponseBody after read.")
		}
	}(resp.Body)

	// Read the HTTP response body
	chatGPTResponseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response chatGPTResponseBody:", err)
		return nil, "Error reading response chatGPTResponseBody", true
	}

	// return response body, empty error string, and "false" if there were no errors
	return chatGPTResponseBody, "", false
}

func setupChatGPTAPI() (string, string) {

	openaiAPIURL := "https://api.openai.com/v1/completions"

	apiKey, exists := os.LookupEnv("GPT_API_KEY")
	if !exists {
		log.Fatal("Error: Environment variable GPT_API_KEY does not exist.")
	}
	return openaiAPIURL, apiKey
}

func extractGPTJson(jsonData string) string {

	var sentence string

	type Response struct {
		ID      string `json:"id"`
		Object  string `json:"object"`
		Created int64  `json:"created"`
		Model   string `json:"model"`
		Choices []struct {
			Text         string      `json:"text"`
			Index        int         `json:"index"`
			Logprobs     interface{} `json:"logprobs"`
			FinishReason string      `json:"finish_reason"`
		} `json:"choices"`
		Usage struct {
			PromptTokens     int `json:"prompt_tokens"`
			CompletionTokens int `json:"completion_tokens"`
			TotalTokens      int `json:"total_tokens"`
		} `json:"usage"`
	}

	var response Response
	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return "Error unmarshaling JSON"
	}

	if len(response.Choices) > 0 {
		sentence = response.Choices[0].Text
		//fmt.Println("Extracted sentence:", sentence)

		// Remove two leading newline characters
		sentence = strings.TrimPrefix(sentence, "\n\n")

	} else {
		fmt.Println(response)
		fmt.Println("No choices found in the JSON")
	}
	return sentence
}
