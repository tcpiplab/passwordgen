package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"unicode"
)

func RandomYearOrFloat() string {
	// Seed the random number generator with the current Unix timestamp
	rand.Seed(time.Now().UnixNano())

	// Randomly decide between year (0) or float (1)
	yearOrFloat := rand.Intn(2)

	minYear := 0
	maxYear := 2000

	if yearOrFloat == 0 {
		// Generate and return random year as a string
		randomYear := rand.Intn(maxYear-minYear+1) + minYear
		return strconv.Itoa(randomYear)
	} else {
		// Generate and return random float as a string
		minFloat := 0.0
		maxFloat := 99.99
		randomFloat := minFloat + rand.Float64()*(maxFloat-minFloat)
		return fmt.Sprintf("%.2f", randomFloat)
	}
}

func createMemorablePassword(requestedPasswordLength int) string {

	var memorablePassword string

	memorablePassword = chooseMemorableTransform(memorablePassword, requestedPasswordLength)

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
	randomYear := RandomYearOrFloat()

	if requestedPasswordLength >= 20 {

		// 1492Mhz
		randomYear = appendRandomUnit(randomYear)
	}

	// Swordfish[1492] or Swordfish[1492Mhz]
	memorablePassword = capitalizeFirstLetter(randomWord)

	// Seed the random number generator with the current Unix timestamp
	rand.Seed(time.Now().UnixNano())

	// Randomly decide between wrapping (0) or delimiting (1)
	wrapOrDelimit := rand.Intn(2)

	if wrapOrDelimit == 0 {

		// Swordfish[1492]
		memorablePassword += padString(randomYear)

	} else {

		// TODO: Randomly choose a delimiter from a list
		randomDelimiter := RandomDelimiter()
		// TODO: Modify unit test to accommodate this change
		// Swordfish_1492
		memorablePassword += randomDelimiter + randomYear
	}

	return memorablePassword

}

func memorableTransformTwo(memorablePassword string, requestedPasswordLength int) string {

	randomWord := getWordFromCompressedDictionary(dictionaryData)
	randomYear := RandomYearOrFloat()

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
	randomYear := RandomYearOrFloat()

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
	randomYear := RandomYearOrFloat()

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
	randomYear := RandomYearOrFloat()

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

// RandomDelimiter returns a random delimiter from a list of special characters.
func RandomDelimiter() string {
	// Seed the random number generator with the current Unix timestamp
	rand.Seed(time.Now().UnixNano())

	delimiters := []string{"!", "@", "#", "$", "%", "^", "&", "*", "(", ")"}
	randomIndex := rand.Intn(len(delimiters))
	return delimiters[randomIndex]
}

// TODO: Write unit tests for all these functions

// TODO: Move unit tests to a separate package
// That will require making all functions in main and in
// main_test (the test dir name I should use) be exportable,
// meaning starting with a capital letter.
