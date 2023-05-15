package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// RandomYearOrFloat This function generates either a random year or float with a
// 33% chance of producing either result.
func RandomYearOrFloat() string {

	// If Seed is not called, the generator is seeded randomly at program startup.
	// Seed the random number generator with the current Unix timestamp
	//rand.Seed(time.Now().UnixNano())

	// Randomly decide between year (0) or float (1)
	yearOrFloat := rand.Intn(3)

	minYear := 0
	maxYear := 2000

	// Return a year 33% of the time
	if yearOrFloat == 0 || yearOrFloat == 1 {
		// Generate and return random year as a string
		randomYear := rand.Intn(maxYear-minYear+1) + minYear
		return strconv.Itoa(randomYear)
	} else {
		// Generate and return random float as a string
		minFloat := 0.0
		maxFloat := 99.99
		randomFloat := minFloat + rand.Float64()*(maxFloat-minFloat)
		return fmt.Sprintf("%.1f", randomFloat)
	}
}

// RandomYearOrInt This function randomly returns a year, int, the current year,
// or a double-digit from globally initialized pseudo-random generators.
func RandomYearOrInt() string {

	// initialize global pseudo random generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Randomly decide between year, int, current year, or double digit
	yearOrInt := r.Intn(4)

	minYear := 1900
	maxYear := time.Now().Year()

	// Return a year 25% of the time
	if yearOrInt == 0 {

		// initialize global pseudo random generator
		r0 := rand.New(rand.NewSource(time.Now().UnixNano()))

		// Generate and return random year as a string
		randomYear := r0.Intn(maxYear-minYear+1) + minYear
		return strconv.Itoa(randomYear)

	} else if yearOrInt == 1 {

		// initialize global pseudo random generator
		r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

		// generates a random number between 0 and 99
		randomInt := r1.Intn(100)

		return fmt.Sprintf("%d", randomInt)

	} else if yearOrInt == 2 {

		currentYear := time.Now().Year()

		// Return the current year
		return strconv.Itoa(currentYear)

	} else {

		return randomDoubledDigit()
	}
}

// randomDoubledDigit This function generates a random double-digit string using
// the pseudo random generator and formatting it for output.
func randomDoubledDigit() string {

	// initialize global pseudo random generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// generates a random number between 0 and 9
	intBetween0and9 := r.Intn(10)

	return fmt.Sprintf("%02d", intBetween0and9*11)
}

// createMemorable3Password This function creates a more memorable password
// by combining randomly chosen adjective/noun pairs with a special character and a year or integer.
// For example:
//
//	JellyDonut$2023
//	PurplePenguin#45
//	SunnyDay@Beach2023
//	MountainHiking!79
//	CrispyBacon&Toast
//	SnowyWinter$December
//	TravelWorld@2024
//	ChocolateCake#Yum
//	MusicLover@77Jazz
//	FitnessGoal!10kRun
func createMemorable3Password() string {

	//var (
	//	_, noun, _, adjective, _, _,
	//	_, _,
	//	_ = getVocabWords()
	//)

	//var (
	//	verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()
	//)

	var (
		verb, noun, adverb, adjective, _, _, _, _, _ = getVocabWords()
	)

	var memorable3Password string

	adjective = capitalizeFirstLetter(adjective)
	noun = capitalizeFirstLetter(noun)
	verb = capitalizeFirstLetter(verb)
	adverb = capitalizeFirstLetter(adverb)
	specialChar := getRandomSpecialChar(true)
	randomYear := RandomYearOrInt()

	// initialize global pseudo random generator
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))

	// generates a random number between 0 and 1
	lexicalChoice := r1.Intn(2)

	if lexicalChoice == 0 {

		memorable3Password = createNounAdjPassword(memorable3Password, adjective, noun, specialChar, randomYear)

	} else if lexicalChoice == 1 {

		// createVerbAdvPassword()
		memorable3Password = createVerbAdvPassword(memorable3Password, adverb, verb, specialChar, randomYear)

	} else if lexicalChoice == 2 {

		// createPronounNounPassword()

	} else if lexicalChoice == 3 {

		// createPronounVerbPassword()

	} else if lexicalChoice == 4 {

		// createPossessiveNounPassword()
	}

	return memorable3Password
}

// createNounAdjPassword This function creates a memorable noun-adjective
// password using random numbers, adjectives, nouns, special characters and
// years.
func createNounAdjPassword(memorable3Password string, adjective string, noun string, specialChar string, randomYear string) string {

	// initialize global pseudo random generator
	r2 := rand.New(rand.NewSource(time.Now().UnixNano()))

	// generates a random number between 0 and 3
	intBetween0and3 := r2.Intn(4)

	// Half the time, choose the pattern like "FeistySample|2023" because it reads easier
	if intBetween0and3 == 0 || intBetween0and3 == 3 {

		memorable3Password = adjective + noun + specialChar + randomYear

	} else if intBetween0and3 == 1 {

		memorable3Password = randomYear + specialChar + adjective + noun

	} else if intBetween0and3 == 2 {

		memorable3Password = adjective + randomYear + specialChar + noun
	}
	return memorable3Password
}

// createVerbAdvPassword This function creates a secure yet memorable password by
// randomly combining a verb, adverb, special character and a random year.
func createVerbAdvPassword(memorable3Password string, adverb string, verb string, specialChar string, randomYear string) string {

	// initialize global pseudo random generator
	r2 := rand.New(rand.NewSource(time.Now().UnixNano()))

	// generates a random number between 0 and 3
	intBetween0and3 := r2.Intn(4)

	// Half the time, choose the pattern like "FeistySample|2023" because it reads easier
	if intBetween0and3 == 0 || intBetween0and3 == 3 {

		memorable3Password = verb + adverb + specialChar + randomYear

	} else if intBetween0and3 == 1 {

		memorable3Password = randomYear + specialChar + verb + adverb

	} else if intBetween0and3 == 2 {

		memorable3Password = verb + randomYear + specialChar + adverb
	}
	return memorable3Password
}

// getRandomSpecialChar This function returns a random special character.
// If noBrackets is true it will not return any chars that are part of bracket pairs.
func getRandomSpecialChar(noBrackets bool) string {

	var specialCharacters string

	if noBrackets == false {

		specialCharacters = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"

	} else if noBrackets == true {

		specialCharacters = "!\"#$%&'*+,-./:;=?@^_`|~"
	}

	// initialize global pseudo random generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return string(specialCharacters[r.Intn(len(specialCharacters))])
}

// createMemorablePassword This function creates a secure and memorable password
// to help keep your data secure. This function is deprecated because the passwords
// are not very memorable.
func createMemorablePassword(requestedPasswordLength int) string {

	var memorablePassword string

	memorablePassword = chooseMemorableTransform(memorablePassword, requestedPasswordLength)

	return memorablePassword
}

// chooseMemorableTransform This function provides a selection of different
// transformations to choose from in order to create a memorable password.
func chooseMemorableTransform(memorablePassword string, requestedPasswordLength int) string {

	// If Seed is not called, the generator is seeded randomly at program startup.
	//rand.Seed(time.Now().UnixNano())

	randomChoice := rand.Intn(7)
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
	case 5:
		memorablePassword = memorableTransformSix(memorablePassword, requestedPasswordLength)
	case 6:
		memorablePassword = memorableTransformSeven(memorablePassword, requestedPasswordLength)
	default:
		// This case should never be reached, but it's here for completeness
		fmt.Println("Invalid choice.")
	}

	return memorablePassword
}

// memorableTransformOne This function creates a memorable password by combining
// a random word with a randomly generated year or float and optionally appending
// a unit.
func memorableTransformOne(memorablePassword string, requestedPasswordLength int) string {

	randomWord := getWordFromCompressedDictionary(dictionaryData)
	randomYear := RandomYearOrFloat()

	if requestedPasswordLength >= 20 {

		// 1492Mhz
		randomYear = appendRandomUnit(randomYear)
	}

	// Swordfish[1492] or Swordfish[1492Mhz]
	memorablePassword = capitalizeFirstLetter(randomWord)

	// If Seed is not called, the generator is seeded randomly at program startup.
	// Seed the random number generator with the current Unix timestamp
	//rand.Seed(time.Now().UnixNano())

	// Randomly decide between wrapping (0) or delimiting (1)
	wrapOrDelimit := rand.Intn(2)

	if wrapOrDelimit == 0 {

		// Swordfish[1492]
		memorablePassword += padString(randomYear)

	} else {

		randomDelimiter := RandomDelimiter()

		// Swordfish_1492
		memorablePassword += randomDelimiter + randomYear
	}

	return memorablePassword

}

// memorableTransformTwo This function generates a memorable password with a
// chosen length, by combining a random year or a year and a unit, with a random
// word and ensuring the first letter of the word is capitalised.
func memorableTransformTwo(memorablePassword string, requestedPasswordLength int) string {

	randomWord := getWordFromCompressedDictionary(dictionaryData)
	randomYear := RandomYearOrFloat()

	if requestedPasswordLength >= 20 {

		// 1492Mhz
		randomYear = appendRandomUnit(randomYear)
	}

	// If Seed is not called, the generator is seeded randomly at program startup.
	// Seed the random number generator with the current Unix timestamp
	//rand.Seed(time.Now().UnixNano())

	// Randomly decide between wrapping (0) or delimiting (1)
	wrapOrDelimit := rand.Intn(2)

	memorablePassword = randomYear

	if wrapOrDelimit == 0 {

		// 1492[Swordfish] or 1492Mhz[Swordfish]
		memorablePassword += padString(capitalizeFirstLetter(randomWord))

	} else {

		randomDelimiter := RandomDelimiter()

		// 1492_Swordfish or 1492Mhz_Swordfish
		memorablePassword += randomDelimiter + capitalizeFirstLetter(randomWord)

	}

	return memorablePassword

}

// memorableTransformThree This function creates strong, memorable passwords by
// combining random words with random years or years with units.
func memorableTransformThree(memorablePassword string, requestedPasswordLength int) string {

	randomWord := getWordFromCompressedDictionary(dictionaryData)
	randomYear := RandomYearOrFloat()

	if requestedPasswordLength >= 20 {

		// 1492Mhz
		randomYear = appendRandomUnit(randomYear)
	}

	// Randomly decide between wrapping (0) or delimiting (1)
	wrapOrDelimit := rand.Intn(2)

	if wrapOrDelimit == 0 {

		// [Swordfish]1492 or [Swordfish]1492Mhz
		memorablePassword = padString(capitalizeFirstLetter(randomWord))

	} else {

		randomDelimiter := RandomDelimiter()

		// [Swordfish_]1492 or [Swordfish_]1492Mhz
		memorablePassword = padString(capitalizeFirstLetter(randomWord) + randomDelimiter)
	}

	memorablePassword += randomYear

	return memorablePassword
}

// memorableTransformFour This function can be used to generate a memorable
// password through the use of a random year, word, and delimiter.
func memorableTransformFour(memorablePassword string, requestedPasswordLength int) string {

	randomWord := getWordFromCompressedDictionary(dictionaryData)
	randomYear := RandomYearOrFloat()

	if requestedPasswordLength >= 20 {

		// 1492Mhz
		randomYear = appendRandomUnit(randomYear)
	}

	// Randomly decide between wrapping (0) or delimiting (1)
	wrapOrDelimit := rand.Intn(2)

	if wrapOrDelimit == 0 {

		// [1492]Swordfish or [1492Mhz]Swordfish
		memorablePassword = padString(randomYear)

	} else {

		randomDelimiter := RandomDelimiter()

		// [_1492]Swordfish or [_1492Mhz]Swordfish
		memorablePassword = padString(randomDelimiter + randomYear)
	}

	memorablePassword += capitalizeFirstLetter(randomWord)

	return memorablePassword
}

// memorableTransformFive This function generates a memorable password that
// combines elements of the two random words, the random year (or year and unit)
// and the random delimiter.
func memorableTransformFive(memorablePassword string, requestedPasswordLength int) string {

	randomWordOne := getWordFromCompressedDictionary(dictionaryData)

	randomWordTwo := getWordFromCompressedDictionary(dictionaryData)
	randomYear := RandomYearOrFloat()

	if requestedPasswordLength > 25 {

		// 1492Mhz
		randomYear = appendRandomUnit(randomYear)
	}

	// [Swordfish-1492-Bankrupt] or [Swordfish-1492Mhz-Bankrupt]
	randomDelimiter := RandomDelimiter()
	wordPair := capitalizeFirstLetter(randomWordOne)
	wordPair += randomDelimiter + randomYear + randomDelimiter
	wordPair += capitalizeFirstLetter(randomWordTwo)
	memorablePassword += padString(wordPair)

	return memorablePassword
}

// memorableTransformSix This function creates a memorable password by appending
// a word pair and a year or year unit to a memorable password.
func memorableTransformSix(memorablePassword string, requestedPasswordLength int) string {

	randomAdjective := getEnglishVocabWord("adjective")

	randomNoun := getEnglishVocabWord("noun")
	randomYear := RandomYearOrFloat()

	if requestedPasswordLength > 25 {

		// 1492Mhz
		randomYear = appendRandomUnit(randomYear)
	}

	// [Swordfish-1492-Bankrupt] or [Swordfish-1492Mhz-Bankrupt]
	randomDelimiter := RandomDelimiter()
	wordPair := capitalizeFirstLetter(randomAdjective)
	wordPair += randomDelimiter + capitalizeFirstLetter(randomNoun) + randomDelimiter
	wordPair += randomYear
	memorablePassword += padString(wordPair)

	return memorablePassword
}

// memorableTransformSeven This function transforms a memorable password using an
// english vocabulary word, a random year or float, and optionally a random unit.
func memorableTransformSeven(memorablePassword string, requestedPasswordLength int) string {

	randomVerb := getEnglishVocabWord("verb")

	randomAdverb := getEnglishVocabWord("adverb")
	randomYear := RandomYearOrFloat()

	if requestedPasswordLength > 25 {

		// 1492Mhz
		randomYear = appendRandomUnit(randomYear)
	}

	// [Generate-Brightly-1492] or [Generate-Brightly-1492Mhz]
	randomDelimiter := RandomDelimiter()
	wordPair := capitalizeFirstLetter(randomVerb)
	wordPair += randomDelimiter + capitalizeFirstLetter(randomAdverb) + randomDelimiter
	wordPair += randomYear
	memorablePassword += padString(wordPair)

	return memorablePassword
}

// capitalizeFirstLetter This function capitalizes the first letter of a given string.
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

	// If Seed is not called, the generator is seeded randomly at program startup.
	//rand.Seed(time.Now().UnixNano())
	randomUnit := units[rand.Intn(len(units))]

	return number + randomUnit
}

// RandomDelimiter returns a random delimiter from a list of special characters.
func RandomDelimiter() string {

	// If Seed is not called, the generator is seeded randomly at program startup.
	// Seed the random number generator with the current Unix timestamp
	//rand.Seed(time.Now().UnixNano())

	delimiters := []string{"!", "@", "#", "$", "%", "^", "&", "*", "-", "_", "+", "=", "~", "`", ".", "|", ":", "/", "\\"}
	randomIndex := rand.Intn(len(delimiters))
	return delimiters[randomIndex]
}

// TODO: Write unit tests for all these functions

// TODO: Move unit tests to a separate package
// That will require making all functions in main and in
// main_test (the test dir name I should use) be exportable,
// meaning starting with a capital letter.

// createMnemonicFromSentence uses the strings.FieldsFunc function to split the input sentence
// into words, considering a space as a separator. Then, it iterates through the
// words and writes the first letter of each word to a strings.Builder. Finally,
// it returns the resulting string. In the example given, if the input sentence
// is "Hello 42 worlds!", the output will be "H42w!".
func createMnemonicFromSentence(sentence string) string {

	// Instantiate a builder object we'll use to build the mnemonic password
	var result strings.Builder

	// Begin by setting this boolean to true.
	// It will decide when a character should be added to the mnemonic password
	shouldAddRune := true

	// Create a placeholder for each char to be evaluated in the next iteration.
	// This will allow us to grab punctuation.
	//var previousChar int32

	// Loop through every character in the sentence
	for _, ch := range sentence {

		if shouldAddRune {

			// If the character is not a space
			if !unicode.IsSpace(ch) {

				// Add the character to the mnemonic password
				result.WriteString(string(ch))
			}

			//If the current char is a number or space, set the bool to true so that the
			//next iteration of the loop will add the char to the password.
			if strings.ContainsRune(",.;?!-0123456789", ch) || unicode.IsNumber(ch) || unicode.IsSpace(ch) {

				shouldAddRune = true

			} else {

				// Tell the next iteration of the loop to not add the next char to the password.
				// For example if the current char is just a letter, we don't want the next letter.
				// TODO: This means we will never get punctuation marks
				shouldAddRune = false

			}

		}

		if unicode.IsSpace(ch) {
			shouldAddRune = true
		}
	}

	return result.String()

}
