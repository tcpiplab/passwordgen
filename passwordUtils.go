package main

import (
	"bytes"
	"math/rand"
	"strings"
	"time"
)

func randStringPassword(lengthOfRandString int) string {

	// Set allowed characters
	var allowedCharacters = []int32("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#^&*()[]{}%")

	// Make a list of type int32 of the length the user requested their passwords should be
	listOfInt32Characters := make([]int32, lengthOfRandString)

	// Seed the randomness
	rand.Seed(time.Now().UnixNano())

	for i := range listOfInt32Characters {

		// Grab random chars and put them in the list. But only from the set of allowed characters
		listOfInt32Characters[i] = allowedCharacters[rand.Intn(len(allowedCharacters))]
	}

	// Return a new random password string
	return string(listOfInt32Characters)
}

func trimPassword(password string, requestedPasswordLength int) string {
	if requestedPasswordLength >= len(password) {
		return password
	}

	rand.Seed(time.Now().UnixNano())
	trimPosition := rand.Intn(len(password) - requestedPasswordLength + 1)

	switch trimPosition {
	case 0:
		return password[:requestedPasswordLength]
	case len(password) - requestedPasswordLength:
		return password[len(password)-requestedPasswordLength:]
	default:
		trimStart := trimPosition / 2
		trimEnd := trimStart + requestedPasswordLength
		return password[trimStart:trimEnd]
	}
}

// randomWordChain() generates a random word-chain password of the specified length.
//
//	Parameters:
//	requestedPasswordLength - the length of the password to generate
//
//	Returns:
//	A string representing the generated password
func randomWordChain(requestedPasswordLength int) string {

	var buffer bytes.Buffer

	// Choose a single delimiter to place between the words
	delimiters := "-_=+/\\|~^$#@&*:."
	delimiter := string(delimiters[rand.Intn(len(delimiters))])

	var word string

	for i := 0; i < requestedPasswordLength; i += len(word) {

		// Grab a word from the compressed dictionary
		word = getWordFromCompressedDictionary(dictionaryData)

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

	// Colorize word-chain output
	colorizeCharactersWindows(requestedPasswordLength, output)

	return output
}

func createMixedPassword(str string) string {
	// create a slice of functions
	listOfFunctions := []func(string) string{
		padString,
		surroundString,
		//randomCase,
		randPadString,
	}

	// apply each function to the string in the shuffled order
	for _, f := range listOfFunctions {
		str = f(str)
	}

	return str
}
