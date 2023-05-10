package main

import (
	"math/rand"
	"time"
)

func padString(s string) string {
	//leftChar := '['
	var rightChar rune

	// Initialize a slice containing the characters to choose from
	leftMostChars := []rune{'{', '[', '(', '<'}

	// Seed the random number generator
	//rand.Seed(time.Now().UnixNano())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Randomly select a character from the slice
	leftChar := leftMostChars[r.Intn(len(leftMostChars))]

	rightChar = matchCharacterPair(leftChar, rightChar)

	// Add the left and right characters to the input string
	paddedStr := string(leftChar) + s + string(rightChar)

	//fmt.Printf(paddedStr)
	return paddedStr
}

func matchCharacterPair(leftChar rune, rightChar rune) rune {

	rightMostChars := []rune{'}', ']', ')', '>'}

	if leftChar == '(' {
		rightChar = ')'
	} else if leftChar == '{' {
		rightChar = '}'
	} else if leftChar == '[' {
		rightChar = ']'
	} else if leftChar == '<' {
		rightChar = '>'
	} else {
		rightChar = rightMostChars[rand.Intn(len(rightMostChars))]
	}
	return rightChar
}

func surroundString(input string) string {

	// Only run this transformation routine about 25% of the times it is called
	// Generate a random number between 0 and 3
	//rand.Seed(time.Now().UnixNano())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := r.Intn(4)

	// Only perform the routines if the random number is 0
	if randomNumber == 0 {

		// Initialize a slice containing the valid surrounding characters
		chars := []rune{'_', '-', '*', '&', '^', '%', '$', '#', '@', '!', '~', '?', '.'}

		// Seed the random number generator
		//rand.Seed(time.Now().UnixNano())
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		// Randomly select a character from the slice of valid characters
		randChar := chars[r.Intn(len(chars))]

		// Surround the input string with the randomly selected character
		output := string(randChar) + input + string(randChar)

		//fmt.Printf("Input string: %s\nOutput string: %s\n", input, output)

		return output

	} else {

		return input
	}

}

// const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const charset = "1234567890"

func randomStringNumbers(length int) string {
	// Seed the random number generator
	//rand.Seed(time.Now().UnixNano())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate a random string of the specified length
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}

	return string(b)
}

func randPadString(input string) string {

	// Determine the number of characters to pad
	numPadding := requestedPasswordLength - len(input)

	// If the input string is already longer than the desired length, return the input string
	if numPadding <= 0 {
		return input
	}

	// Generate random strings for the left and right padding
	leftPadding := randomStringNumbers(numPadding / 2)
	rightPadding := randomStringNumbers(numPadding - len(leftPadding))

	// Concatenate the left and right padding with the input string
	output := leftPadding + input + rightPadding

	// Generate a random position to insert the input string
	//rand.Seed(time.Now().UnixNano())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	insertPos := r.Intn(len(output) - len(input) + 1)

	// Insert the input string at the random position and return the result
	output = output[:insertPos] + input + output[insertPos+len(input):]
	return output
}
