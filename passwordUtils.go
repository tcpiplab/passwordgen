package main

import (
	"bytes"
	crand "crypto/rand"
	"encoding/binary"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func randStringPassword(lengthOfRandString int, hexOnly bool) string {

	var allowedCharacters []int32

	// Set allowed characters
	if !hexOnly {

		allowedCharacters = []int32("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#^&*()[]{}%")

	} else {

		allowedCharacters = []int32("ABCDEF0123456789")

	}

	// Make a list of type int32 of the length the user requested their passwords should be
	listOfInt32Characters := make([]int32, lengthOfRandString)

	// Seed the randomness
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// TODO: Can we delete this if/else here? Test on Windows.
	if OS == "linux" || OS == "darwin" || OS == "unix" {

		//rand.Seed(time.Now().UnixNano())
		//r := rand.New(rand.NewSource(time.Now().UnixNano()))

	} else if OS == "windows" {

		var seed int64
		err := binary.Read(crand.Reader, binary.LittleEndian, &seed)
		if err != nil {
			panic(err)
		}
		//fmt.Println(seed)
		rand.Seed(seed)
	}

	for i := range listOfInt32Characters {

		// Grab random chars and put them in the list. But only from the set of allowed characters
		listOfInt32Characters[i] = allowedCharacters[r.Intn(len(allowedCharacters))]
	}

	// Return a new random password string
	return string(listOfInt32Characters)
}

func trimPassword(password string, requestedPasswordLength int) string {
	if requestedPasswordLength >= len(password) {
		return password
	}

	//rand.Seed(time.Now().UnixNano())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	trimPosition := r.Intn(len(password) - requestedPasswordLength + 1)

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

// createWordChain() generates a random word-chain password of the specified length.
//
//	Parameters:
//	requestedPasswordLength - the length of the password to generate
//
//	Returns:
//	A string representing the generated password
func createWordChain(requestedPasswordLength int) string {

	var buffer bytes.Buffer

	// Choose a single delimiter to place between the words
	delimiters := "-_=+/\\|~^$#@&*:."
	delimiter := string(delimiters[rand.Intn(len(delimiters))])

	var word string

	//for i := 0; i < requestedPasswordLength; i += len(word) {
	for i := 0; i < requestedPasswordLength; i++ {

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
	wordChain := strings.ReplaceAll(buffer.String(), " ", "_")

	// Truncate the resulting word-chain password to the specified length
	// by removing characters from the right side
	//if len(wordChain) > requestedPasswordLength {
	//
	//	wordChain = strings.TrimSpace(wordChain[:requestedPasswordLength])
	//}

	// TODO: this is commented out but will it fail on Windows?
	// Colorize word-chain wordChain
	//colorizeCharactersWindows(requestedPasswordLength, wordChain)

	// Fix bug with trailing delimiter in word chains
	wordChain = removeTrailingSpecialChar(wordChain)

	return wordChain
}

// removeTrailingSpecialChar removes the trailing special character from a string,
// if it exists. If the last character of the input string is a lowercase letter,
// the function returns the original string unchanged.
//
//	Args:
//	  s (string): The input string to remove the trailing special character from.
//
//	Returns:
//	  string: The updated string with the trailing special character removed, or
//	  the original string if the last character is a lowercase letter.
func removeTrailingSpecialChar(s string) string {
	lastChar := s[len(s)-1:]
	if lastChar >= "a" && lastChar <= "z" {
		return s
	}
	return s[:len(s)-1]
}

func shuffleStringTransforms(str string) string {
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

// createMemorable2Password() generates a memorable password if memorable2 is true, and random passwords otherwise.
//
//	Parameters:
//	  memorable2: A boolean indicating whether memorable2 passwords are requested.
//	  randomPasswords: A boolean indicating whether random passwords are requested.
//	  rows: An integer specifying the number of rows in the output.
//
//	Returns:
//	  A string containing the generated password.
func createMemorable2Password(memorable2 bool, randomPasswords bool, rows int) string {

	var memorable2Password string

	if memorable2 {

		// Need to do this for memorable2 passwords to work
		randomPasswords = false

		arrWords := getArrayFromCompressedDictionary(rows / 2)

		adjective := getEnglishVocabWord("adjective")
		noun := getEnglishVocabWord("noun")

		randomDelimiter := RandomDelimiter()

		var inputStr string

		if requestedPasswordLength < 12 {

			// For now just grab the first word in the array
			inputStr = capitalizeFirstLetter(arrWords[0])
			inputStr = randomDelimiterAppendOrPrepend(inputStr)
			inputStr = randomDigitAppendOrPrepend(inputStr)

		} else if requestedPasswordLength <= 20 {

			inputStr = capitalizeFirstLetter(adjective) + randomDelimiter + capitalizeFirstLetter(noun)
			inputStr = randomDigitAppendOrPrepend(inputStr)
			inputStr = surroundString(inputStr)

		} else if requestedPasswordLength > 20 {

			inputStr = surroundString(
				capitalizeFirstLetter(arrWords[0]) + randomDelimiter + capitalizeFirstLetter(arrWords[1]) + randomDelimiter + capitalizeFirstLetter(arrWords[2]))
		}

		//memorable2Password = shuffleStringTransforms(inputStr)
		memorable2Password = inputStr

	}
	return memorable2Password
}

func randomDigitAppendOrPrepend(inputStr string) string {

	// Select a random number between 0 and 9
	digit := rand.Intn(10)
	// Convert the number to a string
	strDigit := strconv.Itoa(digit)

	// The new way to seed randomness each time a function is called
	// Otherwise randomness is only seeded at the start of runtime
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Randomly choose to prepend or append
	//if rand.Float32() < 0.5 {
	if r.Float32() < 0.5 {
		inputStr = strings.Join([]string{strDigit, inputStr}, "")
	} else {
		inputStr = strings.Join([]string{inputStr, strDigit}, "")
	}

	return inputStr
}

func randomDelimiterAppendOrPrepend(inputStr string) string {

	// Select a random delimiter
	randomDelimiter := RandomDelimiter()

	// The new way to seed randomness each time a function is called
	// Otherwise randomness is only seeded at the start of runtime
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Randomly choose to prepend or append
	//if rand.Float32() < 0.5 {
	if r.Float32() < 0.5 {
		inputStr = strings.Join([]string{randomDelimiter, inputStr}, "")
	} else {
		inputStr = strings.Join([]string{inputStr, randomDelimiter}, "")
	}

	return inputStr
}

func createPassphrase(requestedPasswordLength int) string {

	// fmt.Printf("requestedPasswordLength == '%d'", requestedPasswordLength)

	//if requestedPasswordLength == 0 {
	//
	//	// Hardcode this default if none is supplied at the command line
	//	requestedPasswordLength = 5
	//}

	arrOfRandomWords := getArrayFromCompressedDictionary(requestedPasswordLength)

	// Join the array into a single string with a comma separator
	passphrase := strings.Join(arrOfRandomWords[:], " ")

	return passphrase
}
