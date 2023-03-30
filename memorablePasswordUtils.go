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
	minYear := 0
	maxYear := 2000
	randomYear := rand.Intn(maxYear-minYear+1) + minYear

	return strconv.Itoa(randomYear)
}

func createMemorablePassword(requestedPasswordLength int) string {

	var memorablePassword string

	memorablePassword = chooseMemorableTransform(memorablePassword, requestedPasswordLength)

	//if requestedPasswordLength <= 15 {
	//
	//	memorablePassword = chooseMemorableTransform(memorablePassword, requestedPasswordLength)
	//
	//}  else if requestedPasswordLength <= 20 {
	//
	//	memorablePassword = chooseMemorableTransform(memorablePassword, requestedPasswordLength)
	//	memorablePassword += chooseMemorableTransform(memorablePassword, requestedPasswordLength)
	//
	//} else if requestedPasswordLength <= 30 {
	//
	//	memorablePassword = chooseMemorableTransform(memorablePassword, requestedPasswordLength)
	//	memorablePassword += chooseMemorableTransform(memorablePassword, requestedPasswordLength)
	//	memorablePassword += chooseMemorableTransform(memorablePassword, requestedPasswordLength)
	//
	//} else if requestedPasswordLength > 30 {
	//
	//	memorablePassword = chooseMemorableTransform(memorablePassword, requestedPasswordLength)
	//	memorablePassword += chooseMemorableTransform(memorablePassword, requestedPasswordLength)
	//	memorablePassword += chooseMemorableTransform(memorablePassword, requestedPasswordLength)
	//	memorablePassword += chooseMemorableTransform(memorablePassword, requestedPasswordLength)
	//}

	return memorablePassword
}

func chooseMemorableTransform(memorablePassword string, requestedPasswordLength int) string {
	rand.Seed(time.Now().UnixNano())

	randomChoice := rand.Intn(5)
	switch randomChoice {
	case 0:
		memorablePassword = memorableTransformOne(memorablePassword, requestedPasswordLength)
	case 1:
		memorablePassword = memorableTransformTwo(memorablePassword, requestedPasswordLength)
	case 2:
		memorablePassword = memorableTransformThree(memorablePassword, requestedPasswordLength)
	case 3:
		memorablePassword = memorableTransformFour(memorablePassword, requestedPasswordLength)
	case 4:
		memorablePassword = memorableTransformFive(memorablePassword, requestedPasswordLength)
	default:
		// This case should never be reached, but it's here for completeness
		fmt.Println("Invalid choice.")
	}

	return memorablePassword
}

func memorableTransformOne(memorablePassword string, requestedPasswordLength int) string {

	randomWord := getWordFromCompressedDictionary(dictionaryData)
	randomYear := RandomYear()

	if requestedPasswordLength >= 20 {

		// 1492Mhz
		randomYear = appendRandomUnit(randomYear)
	}

	// Swordfish[1492] or Swordfish[1492Mhz]
	memorablePassword = capitalizeFirstLetter(randomWord)
	memorablePassword += padString(randomYear)

	return memorablePassword

}

func memorableTransformTwo(memorablePassword string, requestedPasswordLength int) string {

	randomWord := getWordFromCompressedDictionary(dictionaryData)
	randomYear := RandomYear()

	if requestedPasswordLength >= 20 {

		// 1492Mhz
		randomYear = appendRandomUnit(randomYear)
	}

	// 1492[Swordfish] or 1492Mhz[Swordfish]
	memorablePassword = randomYear
	memorablePassword += padString(capitalizeFirstLetter(randomWord))

	return memorablePassword

}

func memorableTransformThree(memorablePassword string, requestedPasswordLength int) string {

	randomWord := getWordFromCompressedDictionary(dictionaryData)
	randomYear := RandomYear()

	if requestedPasswordLength >= 20 {

		// 1492Mhz
		randomYear = appendRandomUnit(randomYear)
	}

	// [Swordfish]1492 or [Swordfish]1492Mhz
	memorablePassword = padString(capitalizeFirstLetter(randomWord))
	memorablePassword += randomYear

	return memorablePassword
}

func memorableTransformFour(memorablePassword string, requestedPasswordLength int) string {

	randomWord := getWordFromCompressedDictionary(dictionaryData)
	randomYear := RandomYear()

	if requestedPasswordLength >= 20 {

		// 1492Mhz
		randomYear = appendRandomUnit(randomYear)
	}

	// [1492]Swordfish or [1492Mhz]Swordfish
	memorablePassword = padString(randomYear)
	memorablePassword += capitalizeFirstLetter(randomWord)

	return memorablePassword
}

func memorableTransformFive(memorablePassword string, requestedPasswordLength int) string {

	randomWordOne := getWordFromCompressedDictionary(dictionaryData)

	randomWordTwo := getWordFromCompressedDictionary(dictionaryData)
	randomYear := RandomYear()

	if requestedPasswordLength > 25 {

		// 1492Mhz
		randomYear = appendRandomUnit(randomYear)
	}

	// [Swordfish-1492-Bankrupt] or [Swordfish-1492Mhz-Bankrupt]
	wordPair := capitalizeFirstLetter(randomWordOne)
	wordPair += "-" + randomYear + "-"
	wordPair += capitalizeFirstLetter(randomWordTwo)
	memorablePassword += padString(wordPair)

	return memorablePassword
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}

	// Convert the first character to uppercase and concatenate with the rest of the string
	return string(unicode.ToUpper(rune(s[0]))) + s[1:]
}

// appendRandomUnit appends a random unit from a given list to the input number.
func appendRandomUnit(number string) string {
	units := []string{
		"Mhz", "Ghz", "Mbps", "Mph", "Gbps", "Kbps", "inches", "feet", "miles",
		"Hz", "kHz", "THz", "nm", "mm", "cm", "m", "km", "yd", "mi", "nmi",
		"liters", "gallons", "pints", "quarts", "milliliters", "cubic meters",
		"grams", "kilograms", "pounds", "ounces", "tons", "tonnes",
		"seconds", "minutes", "hours", "days", "weeks", "months", "years",
		"degrees Celsius", "degrees Fahrenheit", "Kelvin",
		"pascals", "bars", "atmospheres", "mmHg", "torr",
		"volts", "amperes", "watts", "ohms", "farads", "henrys", "teslas", "webers",
		"Joules", "calories", "BTU", "ergs", "electronvolts",
		"lumens", "candelas", "lux", "foot-candles",
		"bits", "bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB",
		"dpi", "ppi", "pt", "em", "rem", "px",
	}

	rand.Seed(time.Now().UnixNano())
	randomUnit := units[rand.Intn(len(units))]

	return number + randomUnit
}

// TODO: Write unit tests for all these functions

// TODO: Move unit tests to a separate package
// That will require making all functions in main and in
// main_test (the test dir name I should use) be exportable,
// meaning starting with a capital letter.
