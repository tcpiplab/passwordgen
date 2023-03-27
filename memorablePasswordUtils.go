package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"unicode"
)

func RandomYear() string {
	// Seed the random number generator with the current Unix timestamp
	rand.Seed(time.Now().UnixNano())
	minYear := 1800
	maxYear := 2000
	randomYear := rand.Intn(maxYear-minYear+1) + minYear

	return strconv.Itoa(randomYear)
}

func generateMemorablePassword(requestedPasswordLength int) string {

	var memorablePassword string
	//randomWord := getWordFromCompressedDictionary(dictionaryData)
	//randomYear := RandomYear()

	if requestedPasswordLength <= 15 {

		memorablePassword = chooseMemorableTransform(memorablePassword)

	} else if requestedPasswordLength <= 20 {

		memorablePassword = chooseMemorableTransform(memorablePassword)
		memorablePassword += chooseMemorableTransform(memorablePassword)

	} else if requestedPasswordLength <= 30 {

		memorablePassword = chooseMemorableTransform(memorablePassword)
		memorablePassword += chooseMemorableTransform(memorablePassword)
		memorablePassword += chooseMemorableTransform(memorablePassword)

	} else if requestedPasswordLength > 30 {

		memorablePassword = chooseMemorableTransform(memorablePassword)
		memorablePassword += chooseMemorableTransform(memorablePassword)
		memorablePassword += chooseMemorableTransform(memorablePassword)
		memorablePassword += chooseMemorableTransform(memorablePassword)
	}

	return memorablePassword
}

func chooseMemorableTransform(memorablePassword string) string {
	rand.Seed(time.Now().UnixNano())

	randomChoice := rand.Intn(4)
	switch randomChoice {
	case 0:
		memorablePassword = memorableTransformOne(memorablePassword)
	case 1:
		memorablePassword = memorableTransformTwo(memorablePassword)
	case 2:
		memorablePassword = memorableTransformThree(memorablePassword)
	case 3:
		memorablePassword = memorableTransformFour(memorablePassword)
	default:
		// This case should never be reached, but it's here for completeness
		fmt.Println("Invalid choice.")
	}

	return memorablePassword
}

func memorableTransformOne(memorablePassword string) string {

	randomWord := getWordFromCompressedDictionary(dictionaryData)
	randomYear := RandomYear()

	// Swordfish[1492]
	memorablePassword = capitalizeFirstLetter(randomWord)
	memorablePassword += padString(randomYear)

	return memorablePassword
}

func memorableTransformTwo(memorablePassword string) string {

	randomWord := getWordFromCompressedDictionary(dictionaryData)
	randomYear := RandomYear()

	// 1492[Swordfish]
	memorablePassword = capitalizeFirstLetter(randomYear)
	memorablePassword += padString(randomWord)

	return memorablePassword

}

func memorableTransformThree(memorablePassword string) string {

	randomWord := getWordFromCompressedDictionary(dictionaryData)
	randomYear := RandomYear()

	// [Swordfish]1492
	memorablePassword = padString(capitalizeFirstLetter(randomWord))
	memorablePassword += randomYear

	return memorablePassword
}

func memorableTransformFour(memorablePassword string) string {

	randomWord := getWordFromCompressedDictionary(dictionaryData)
	randomYear := RandomYear()

	// [1492]Swordfish
	memorablePassword = padString(randomYear)
	memorablePassword += capitalizeFirstLetter(randomWord)

	return memorablePassword
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}

	// Convert the first character to uppercase and concatenate with the rest of the string
	return string(unicode.ToUpper(rune(s[0]))) + s[1:]
}
