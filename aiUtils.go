package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func createGrammaticalPasswordAI(nonSensicalSentence string) string {

	openaiAPIURL := "https://api.openai.com/v1/completions"

	type CompletionCreateArgs struct {
		Model       string  `json:"model"`
		Prompt      string  `json:"prompt"`
		MaxTokens   int     `json:"max_tokens"`
		Temperature float64 `json:"temperature"`
	}

	apiKey, exists := os.LookupEnv("GPT_API_KEY")
	if !exists {
		log.Fatal("Error: Environment variable GPT_API_KEY does not exist.")
	}

	// Continue execution if the environment variable exists

	promptSentence := "Change the subject in the following nonsensical sentence so that it makes more sense. " +
		"Change the adverb, adjective, noun, or verb if they don't sound like they belong together: '" +
		nonSensicalSentence + "'"

	data := CompletionCreateArgs{
		Model:  "text-davinci-003",
		Prompt: promptSentence,
		// Any amount of tokens < 12 will truncate some sentences.
		MaxTokens: 12,
		// The best outcomes seem to be with temperature set to 0.
		Temperature: 0,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return "Error marshaling JSON"
	}

	req, err := http.NewRequest("POST", openaiAPIURL, bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return "Error creating request"
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return "Error making request"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "Error reading response body"
	}

	rewrittenSentence := extractGPTJson(string(body))

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
