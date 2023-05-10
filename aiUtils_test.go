package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateGrammaticalPasswordAI(t *testing.T) {
	// Define the test cases
	testCases := []struct {
		name                     string
		nonSensicalSentence      string
		grammaticalAIWithNumbers bool
	}{
		{
			name:                     "Test without numbers",
			nonSensicalSentence:      "The quick brown fox jumps high over the moon.",
			grammaticalAIWithNumbers: false,
		},
		{
			name:                     "Test with numbers",
			nonSensicalSentence:      "The quick brown fox jumps high over the moon.",
			grammaticalAIWithNumbers: true,
		},
	}

	// Run the test cases
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := createGrammaticalPasswordAI(testCase.nonSensicalSentence, testCase.grammaticalAIWithNumbers)

			// Since the result is not deterministic, we'll just check if it's not an empty string
			if result == "" {
				t.Errorf("Expected a non-empty string, got an empty string")
			}
		})
	}
}

func TestExtractGPTJson(t *testing.T) {
	// Define the test cases
	testCases := []struct {
		name     string
		jsonData string
		expected string
	}{
		{
			name: "Valid JSON",
			jsonData: `{
				"id": "example_id",
				"object": "text.completion",
				"created": 1624390123,
				"model": "text-davinci-002",
				"choices": [
					{
						"text": "\n\nHello, world!",
						"index": 0,
						"logprobs": null,
						"finish_reason": "stop"
					}
				],
				"usage": {
					"prompt_tokens": 5,
					"completion_tokens": 20,
					"total_tokens": 25
				}
			}`,
			expected: "Hello, world!",
		},
		{
			name:     "Invalid JSON",
			jsonData: `{"invalid": "json}`,
			expected: "Error unmarshaling JSON",
		},
		{
			name: "Empty choices",
			jsonData: `{
				"id": "example_id",
				"object": "text.completion",
				"created": 1624390123,
				"model": "text-davinci-002",
				"choices": [],
				"usage": {
					"prompt_tokens": 5,
					"completion_tokens": 20,
					"total_tokens": 25
				}
			}`,
			expected: "",
		},
	}

	// Run the test cases
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := extractGPTJson(testCase.jsonData)

			// Check if the result matches the expected value
			if result != testCase.expected {
				t.Errorf("Expected '%s', got '%s'", testCase.expected, result)
			}
		})
	}
}

func Test_makeChatGPTAPIRequest(t *testing.T) {
	type args struct {
		chatGPTRequestData CompletionCreateArgs
		openaiAPIURL       string
		apiKey             string
	}
	tests := []struct {
		name  string
		args  args
		want  []byte
		want1 string
		want2 bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := makeChatGPTAPIRequest(tt.args.chatGPTRequestData, tt.args.openaiAPIURL, tt.args.apiKey)
			assert.Equalf(t, tt.want, got, "makeChatGPTAPIRequest(%v, %v, %v)", tt.args.chatGPTRequestData, tt.args.openaiAPIURL, tt.args.apiKey)
			assert.Equalf(t, tt.want1, got1, "makeChatGPTAPIRequest(%v, %v, %v)", tt.args.chatGPTRequestData, tt.args.openaiAPIURL, tt.args.apiKey)
			assert.Equalf(t, tt.want2, got2, "makeChatGPTAPIRequest(%v, %v, %v)", tt.args.chatGPTRequestData, tt.args.openaiAPIURL, tt.args.apiKey)
		})
	}
}

func Test_setupChatGPTAPI(t *testing.T) {
	tests := []struct {
		name  string
		want  string
		want1 string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := setupChatGPTAPI()
			assert.Equalf(t, tt.want, got, "setupChatGPTAPI()")
			assert.Equalf(t, tt.want1, got1, "setupChatGPTAPI()")
		})
	}
}
