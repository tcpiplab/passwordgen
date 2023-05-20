package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/inflection"
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

// createGrammaticalPasswordAI This function uses the OpenAI API to generate a
// grammatically correct sentence based on the user's definition of
// 'grammatical.'
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
			MaxTokens: 17,
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
			MaxTokens: 14,
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

func createMemorablePasswordAI() string {

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

	promptSentence = "Please return two nouns that are related to each other, " +
		"and one adjective that is related to at least one of the nouns. " +
		"Put them in a simple list, separated by commas. And don't include anything else."

	chatGPTRequestData = CompletionCreateArgs{
		Model:  "text-davinci-003",
		Prompt: promptSentence,
		// Any amount of tokens < 12 will truncate some sentences.
		MaxTokens: 12,
		// The best outcomes seem to be with temperature set to 0.8 for this prompt.
		// Otherwise you get duplicate answers.
		Temperature: 1.3,
	}

	// Make the actual API call
	chatGPTResponseBody, errorString, apiRequestError := makeChatGPTAPIRequest(chatGPTRequestData, openaiAPIURL, apiKey)

	// If the API call returned and error, return the error string
	if apiRequestError {
		return errorString
	}

	commaSeparatedWords := extractGPTJson(string(chatGPTResponseBody))

	// Remove any surrounding single quotes. This happens sometimes.
	commaSeparatedWords = strings.Trim(commaSeparatedWords, "'")

	// Split the string on the comma and space.
	separatedWords := strings.Split(commaSeparatedWords, ",")

	// FIXME: Create fallback lookups from local dictionary. If these are missing it errors out.
	noun1 := separatedWords[0]
	noun2 := separatedWords[1]
	adjective := separatedWords[2]

	noun1 = strings.TrimSpace(noun1)
	noun2 = strings.TrimSpace(noun2)
	adjective = strings.TrimSpace(adjective)

	noun1 = capitalizeFirstLetter(noun1)
	noun2 = capitalizeFirstLetter(noun2)
	adjective = capitalizeFirstLetter(adjective)

	noun1Plural := inflection.Plural(noun1)
	noun2Plural := inflection.Plural(noun2)

	// initialize global pseudo random generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	delimiter := RandomDelimiter()

	digit := strconv.Itoa(r.Intn(10))

	// Use the plural versions of the nouns unless the digit is 1
	if digit != "1" {

		noun1 = noun1Plural
		noun2 = noun2Plural
	}

	// Randomly decide between variations in word order
	wordOrder := r.Intn(3)

	if wordOrder == 0 {

		return adjective + noun1 + delimiter + digit + noun2

	} else if wordOrder == 1 {

		return digit + noun1 + delimiter + adjective + noun2

	} else {

		return noun1 + delimiter + digit + noun2 + "," + adjective
	}

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
