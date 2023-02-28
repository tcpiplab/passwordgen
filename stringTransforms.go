package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var requestedPasswordLength = 20

func main() {
	//inputStr := "Hello World"
	inputStr := callWordApi()

	//paddedStr := padString(inputStr)
	//fmt.Println(paddedStr) // Output: "[Hello World]"
	//
	//surroundedStr := surroundString(inputStr)
	//fmt.Println(surroundedStr) // Output: "_Hello World_"
	//
	//randomCaseStr := randomCase(inputStr)
	//fmt.Println(randomCaseStr)
	//
	//// Call the randPadString function with an integer and a string
	//length := 32
	//randomPadStr := randPadString(length, inputStr)
	//println(randomPadStr)

	fmt.Printf("%s", processString(inputStr))

}

func processString(str string) string {
	// create a slice of functions
	listOfFunctions := []func(string) string{
		padString,
		surroundString,
		randomCase,
		randPadString,
	}

	// shuffle the slice using a random permutation
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(listOfFunctions), func(i, j int) {
		listOfFunctions[i], listOfFunctions[j] = listOfFunctions[j], listOfFunctions[i]
	})

	// apply each function to the string in the shuffled order
	for _, f := range listOfFunctions {
		str = f(str)
	}

	return str
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

	// Only run this transformation routine about 25% of the times it is called
	// Generate a random number between 0 and 3
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(4)

	// Only perform the routines if the random number is 0
	if randomNumber == 0 {

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

	} else {

		return input
	}

}

func randomCase(input string) string {

	// Generate a random number between 0 and 2
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(3)

	// Only perform the routines about 33% of the times that this function
	// is called, i.e., if the random number is 0
	if randomNumber == 0 {

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

	} else {

		return input
	}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func randomString(length int) string {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random string of the specified length
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}

// 	// Call the randPadString function with an integer and a string
//	length := 10
//	input := "hello"
//	output := randPadString(length, input)
//
//	println("Input string:", input)
//	println("Output string:", output)

//func randPadString(input string) string {
//	// Determine the number of characters to pad
//	//numPadding := length - len(input)
//
//	numPadding := requestedPasswordLength - len(input)
//
//	// If the input string is already longer than the desired length, return the input string
//	if numPadding <= 0 {
//		return input
//	}
//
//	// Generate random strings for the left and right padding
//	leftPadding := randomString(numPadding / 2)
//	rightPadding := randomString(numPadding - len(leftPadding))
//
//	// Concatenate the left and right padding with the input string and return the result
//	output := leftPadding + input + rightPadding
//	return output
//}

func randPadString(input string) string {
	// Determine the number of characters to pad
	numPadding := requestedPasswordLength - len(input)

	// If the input string is already longer than the desired length, return the input string
	if numPadding <= 0 {
		return input
	}

	// Generate random strings for the left and right padding
	leftPadding := randomString(numPadding / 2)
	rightPadding := randomString(numPadding - len(leftPadding))

	// Concatenate the left and right padding with the input string
	output := leftPadding + input + rightPadding

	// Generate a random position to insert the input string
	rand.Seed(time.Now().UnixNano())
	insertPos := rand.Intn(len(output) - len(input) + 1)

	// Insert the input string at the random position and return the result
	output = output[:insertPos] + input + output[insertPos+len(input):]
	return output
}
