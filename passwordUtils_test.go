package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRandStringPassword(t *testing.T) {

	// Test with a password length of 10
	passwordLength := 32

	// Seed the randomness
	//rand.Seed(time.Now().UnixNano())

	// Call the function to generate a random password string
	password := randStringPassword(passwordLength)

	// Check if the length of the generated password matches the expected length
	assert.Equal(t, passwordLength, len(password))

	// Check if the generated password contains only allowed characters
	assert.Regexp(t, "^[a-zA-Z0-9!@#^&*()\\[\\]{}%]*$", password)

	// Check for high entropy
	fmt.Printf("--- Testing entropy of: %s\n", password)
	assert.True(t, isHighEntropy(password))
}

func TestTrimPassword(t *testing.T) {

	// Test with a password that is longer than the requested length
	password := "ThisIsALongPassword"
	requestedPasswordLength := 5

	// Call the function to trim the password
	trimmedPassword := trimPassword(password, requestedPasswordLength)

	// Check if the length of the trimmed password matches the expected length
	assert.Equal(t, requestedPasswordLength, len(trimmedPassword))

	// Check if the trimmed password is a substring of the original password
	assert.Contains(t, password, trimmedPassword)

	// Add more tests for other password lengths and edge cases as needed
}

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
	// Check for high entropy
	fmt.Printf("--- Testing entropy of: %s\n", wordChain)
	assert.True(t, isHighEntropy(wordChain))
}

func TestCreateMixedPassword(t *testing.T) {

	// Test with a string input
	inputString := "password"

	// Call the function to create a mixed password
	mixedPassword := shuffleStringTransforms(inputString)

	// Check if the length of the mixed password matches the length of the input string
	//assert.Equal(t, len(inputString), len(mixedPassword))

	// Check if the mixed password contains the input string
	assert.Contains(t, mixedPassword, inputString)

	// Add more tests for other input strings and edge cases as needed
	// Check for high entropy
	fmt.Printf("--- Testing entropy of: %s\n", mixedPassword)
	assert.True(t, isHighEntropy(mixedPassword))
}

func TestIfMixedPasswords(t *testing.T) {
	rows := 10
	//requestedPasswordLength := 15

	// Test when mixedPasswords is false
	outputStr := createMixedPassword(false, true, rows)
	if outputStr != "" {
		t.Errorf("Expected empty string when mixedPasswords is false, but got %s", outputStr)
	}

	// Test when requestedPasswordLength is less than 12
	requestedPasswordLength = 10
	outputStr = createMixedPassword(true, false, rows)

	// Check if the outputStr contains only valid characters
	assert.Regexp(t, "^[a-zA-Z0-9-_+=/\\\\|~^$#@&*:.\"{}\\[\\]<>\\(\\)]*$", outputStr)

	// Check for high entropy
	fmt.Printf("--- Testing entropy of: %s\n", outputStr)
	assert.True(t, isHighEntropy(outputStr))

	// Test when requestedPasswordLength is between 12 and 20
	requestedPasswordLength = 18
	outputStr = createMixedPassword(true, false, rows)

	// Check if the outputStr contains only valid characters
	assert.Regexp(t, "^[a-zA-Z0-9-_+=/\\\\|~^$#@&*:.\"{}\\[\\]<>\\(\\)\\?\\!]*$", outputStr)

	// Check for high entropy
	fmt.Printf("--- Testing entropy of: %s\n", outputStr)
	assert.True(t, isHighEntropy(outputStr))

	// Test when requestedPasswordLength is greater than 20
	requestedPasswordLength = 25
	outputStr = createMixedPassword(true, false, rows)

	// Check if the outputStr contains only valid characters
	assert.Regexp(t, "^[a-zA-Z0-9-_+=/\\\\|~^$#@&*:.\"{}\\[\\]<>\\(\\)\\?%\\!]*$", outputStr)

	// Check for high entropy
	fmt.Printf("--- Testing entropy of: %s\n", outputStr)
	assert.True(t, isHighEntropy(outputStr))
}

func TestCreatePassphrase(t *testing.T) {
	passphrase := createPassphrase()
	words := strings.Split(passphrase, " ")
	if len(words) != 5 {
		t.Errorf("Expected 5 words in passphrase, but got %d", len(words))
	}
}
