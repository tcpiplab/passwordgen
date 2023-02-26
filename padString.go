package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//inputStr := "Hello World"
	inputStr := callWordApi()
	paddedStr := padString(inputStr)
	fmt.Println(paddedStr) // Output: "[Hello World]"
}

func padString(s string) string {
	//leftChar := '['
	var rightChar rune

	// Initialize a slice containing the characters to choose from
	leftMostChars := []rune{'{', '[', '(', '<'}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Randomly select a character from the slice
	leftChar := leftMostChars[rand.Intn(len(leftMostChars))]

	// Check if the left and right characters are a valid pair
	//if !matchCharacterPair(leftChar, rightChar) {
	//	panic("Invalid character pair")
	//}

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
