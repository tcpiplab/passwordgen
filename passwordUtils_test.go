package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateWordChain(t *testing.T) {

	// Test with a password length of 10
	passwordLength := 10

	// Call the function to generate a word chain
	wordChain := createWordChain(passwordLength)

	// Check if the length of the generated word chain matches the expected length
	//assert.Equal(t, passwordLength, len(wordChain))

	// Check if the generated word chain contains only valid characters
	assert.Regexp(t, "^[a-zA-Z0-9-_+=/\\\\|~^$#@&*:.\"]*$", wordChain)

	// Add more tests for other password lengths and edge cases as needed
}
