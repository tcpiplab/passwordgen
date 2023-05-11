package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
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

type Response struct {
	ID     string `json:"id"`
	Object string `json:"object"`
}

// TestMakeChatGPTAPIRequest This function tests that the makeChatGPTAPIRequest()
// function is sending a readable request with an accurate response. It sets up a
// test server that returns a dummy response, creates a dummy
// CompletionCreateArgs struct, and calls the makeChatGPTAPIRequest function with
// the test server's URL and a dummy API key. The test function then checks if
// there were any errors and if the returned response matches the expected
// values.
func TestMakeChatGPTAPIRequest(t *testing.T) {
	// Create a test server that returns a dummy response
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := Response{
			ID:     "test_id",
			Object: "test_object",
		}
		jsonResponse, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := io.WriteString(w, string(jsonResponse))
		if err != nil {
			return
		}
	}))
	defer testServer.Close()

	// Create a dummy CompletionCreateArgs
	chatGPTRequestData := CompletionCreateArgs{
		Model:       "text-davinci-003",
		Prompt:      "Test prompt",
		MaxTokens:   10,
		Temperature: 0,
	}

	// Call the function with the test server URL and a dummy API key
	chatGPTResponseBody, errorString, apiRequestError := makeChatGPTAPIRequest(chatGPTRequestData, testServer.URL, "dummy_key")

	// Check if there was an error
	if apiRequestError {
		t.Errorf("Expected no error, got error: %s", errorString)
	}

	// Unmarshal the response and check if it matches the expected values
	var response Response
	err := json.Unmarshal(chatGPTResponseBody, &response)
	if err != nil {
		t.Errorf("Error unmarshaling JSON: %s", err)
	}

	if response.ID != "test_id" {
		t.Errorf("Expected ID 'test_id', got '%s'", response.ID)
	}

	if response.Object != "test_object" {
		t.Errorf("Expected Object 'test_object', got '%s'", response.Object)
	}
}

func TestSetupChatGPTAPI(t *testing.T) {
	// Save the current value of the GPT_API_KEY environment variable
	originalAPIKey, exists := os.LookupEnv("GPT_API_KEY")

	// Set a dummy API key for the test
	err1 := os.Setenv("GPT_API_KEY", "dummy_key")
	if err1 != nil {
		return
	}

	// Run the function
	openaiAPIURL, apiKey := setupChatGPTAPI()

	// Check if the returned URL is correct
	expectedURL := "https://api.openai.com/v1/completions"
	if openaiAPIURL != expectedURL {
		t.Errorf("Expected URL '%s', got '%s'", expectedURL, openaiAPIURL)
	}

	// Check if the returned API key is correct
	expectedAPIKey := "dummy_key"
	if apiKey != expectedAPIKey {
		t.Errorf("Expected API key '%s', got '%s'", expectedAPIKey, apiKey)
	}

	// Restore the original value of the GPT_API_KEY environment variable
	if exists {
		err2 := os.Setenv("GPT_API_KEY", originalAPIKey)
		if err2 != nil {
			return
		}
	} else {
		err3 := os.Unsetenv("GPT_API_KEY")
		if err3 != nil {
			return
		}
	}
}
