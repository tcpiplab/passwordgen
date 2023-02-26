package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	//inputStr := "Hello World"
	inputStr := callWordApi()

	paddedStr := padString(inputStr)
	fmt.Println(paddedStr) // Output: "[Hello World]"

	surroundedStr := surroundString(inputStr)
	fmt.Println(surroundedStr) // Output: "_Hello World_"

	randomCaseStr := randomCase(inputStr)
	fmt.Println(randomCaseStr)

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

func surroundString(input string) string {

	// Initialize a slice containing the valid surrounding characters
	chars := []rune{'_', '-', '*', '&', '^', '%', '$', '#', '@', '!', '~', '?', '.'}

	// Call the surroundString function with an input string
	//input := "hello"

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Randomly select a character from the slice of valid characters
	randChar := chars[rand.Intn(len(chars))]

	// Surround the input string with the randomly selected character
	output := string(randChar) + input + string(randChar)

	//fmt.Printf("Input string: %s\nOutput string: %s\n", input, output)

	return output
}

func randomCase(input string) string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Convert the input string to a slice of runes
	inputRunes := []rune(input)

	// Iterate over the slice of runes and randomly change the case of each letter
	for i := 0; i < len(inputRunes); i++ {
		if rand.Intn(2) == 0 {
			inputRunes[i] = []rune(strings.ToUpper(string(inputRunes[i])))[0]
		} else {
			inputRunes[i] = []rune(strings.ToLower(string(inputRunes[i])))[0]
		}
	}

	// Convert the slice of runes back to a string and return it
	output := string(inputRunes)
	return output
}
