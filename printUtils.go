package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"math/rand"
	"os"
	"strings"
	"time"
)

// printPasswordTableWindows prints a table of randomized passwords with index numbers to the terminal screen.
//
// The function takes in the number of rows to print, the requested length of each password,
// and an array of passwords to populate. The function prints an index number for each password
// and colors each character of the password string. The table is bordered with a horizontal
// line and cross line characters, and each row is separated with a vertical line. The password
// array is populated with the generated passwords.
//
// Parameters:
// - rows: an int specifying the number of rows to print
// - requestedPasswordLength: an int specifying the length of each password to generate
// - arrayPasswords: a slice of strings representing the passwords to be populated
// Returns: nothing
func printPasswordTableWindows(
	rows int,
	requestedPasswordLength int,
	arrayPasswords []string,
	randomPasswords bool,
	wordChains bool,
	mixedPasswords bool,
	passPhrases bool) []string {

	grey := color.New(color.FgCyan, color.Faint).SprintfFunc()

	underline := grey("─")

	if !passPhrases && !wordChains && !mixedPasswords && !randomPasswords {
		fmt.Printf(
			"%s%s%s\n",
			grey("+────+"),
			strings.Repeat(underline, requestedPasswordLength+2),
			grey("+"),
		)
	}

	// Loop to print rows of index numbers and passwords to the terminal screen
	for rowNumber := 0; rowNumber < ((rows / 2) - 1); rowNumber++ {

		if !passPhrases && !wordChains && !mixedPasswords && !randomPasswords {

			// TODO: Get colors working on Windows
			//red := color.New(color.FgHiRed).SprintFunc()

			rowNumberString := fmt.Sprintf("%02d", rowNumber)

			fmt.Printf("%s ", grey("│"))

			// Print index number in HiRed
			color.NoColor = false
			redIndexNumberWindows := color.New(color.FgHiRed).PrintfFunc()
			redIndexNumberWindows(rowNumberString)
			color.NoColor = true

			fmt.Printf(" %s ", grey("│"))

		}

		if randomPasswords {

			break

			//// Fetch a new randomized password string of the specified length
			//password := randStringPassword(requestedPasswordLength)
			//
			//arrayPasswords[rowNumber] = password
			//
			//// Colorize and print the password
			//colorizeCharactersWindows(requestedPasswordLength, password)

		} else if wordChains {

			//password := createWordChain(requestedPasswordLength)
			//
			//arrayPasswords[rowNumber] = password

			break

		} else if mixedPasswords {

			password := createMixedPassword(mixedPasswords, randomPasswords, rows)

			arrayPasswords[rowNumber] = password

			break

			//password := createMixedPassword(mixedPasswords, randomPasswords, rows)
			//
			//arrayPasswords[rowNumber] = password

			// Colorize and print the password
			//colorizeCharactersWindows(requestedPasswordLength, password)

		} else if passPhrases {

			password := createPassphrase()

			arrayPasswords[rowNumber] = password

			break
		}

		// Vertical line after the password
		fmt.Printf(" %s", grey("│"))

		// Newline at end of row
		fmt.Printf("\n")

		//fmt.Printf("%s of %s %s\n", rowNumber, rows, len(arrayPasswords))

		// If it's the final line we're printing
		if rowNumber == (len(arrayPasswords) - 9) {

			// └
			fmt.Print(grey("+"))
		} else if rowNumber >= 0 {

			// Beginning of row line, middle of table ├
			fmt.Print(grey("+"))
		}

		// Line under password index number, then cross line character ┼
		fmt.Printf("%s%s", strings.Repeat(underline, 4), grey("+"))

		// Line between rows
		fmt.Printf("%s", strings.Repeat(underline, requestedPasswordLength+2))

		// End of row line ┤
		fmt.Printf("%s", grey("+"))

		// Newline at end of row line
		fmt.Printf("\n")

	}
	if passPhrases {

		arrayPasswords = printPassphraseTable()

	} else if wordChains {

		arrayPasswords = printWordChainsTable()

	} else if mixedPasswords {

		arrayPasswords = printMixedPasswordsTable(mixedPasswords, randomPasswords)

	} else if randomPasswords {

		arrayPasswords = printRandomPasswordsTable()

	}

	return arrayPasswords
}

// printPasswordTableUnix prints a table of randomized passwords with index numbers to the terminal screen.
//
// The function takes in the number of rows to print, the requested length of each password,
// and an array of passwords to populate. The function prints an index number for each password
// and colors each character of the password string. The table is bordered with a horizontal
// line and cross line characters, and each row is separated with a vertical line. The password
// array is populated with the generated passwords.
//
// Parameters:
// - rows: an int specifying the number of rows to print
// - requestedPasswordLength: an int specifying the length of each password to generate
// - arrayPasswords: a slice of strings representing the passwords to be populated
// Returns: nothing
func printPasswordTableUnix(arrayPasswords []string, randomPasswords bool, wordChains bool, mixedPasswords bool, passPhrases bool, memorable bool, randomHex bool, grammatical bool, grammaticalAI bool) []string {

	if passPhrases {

		arrayPasswords = printPassphraseTable()

	} else if wordChains {

		arrayPasswords = printWordChainsTable()

	} else if mixedPasswords {

		arrayPasswords = printMixedPasswordsTable(mixedPasswords, randomPasswords)

	} else if randomPasswords {

		arrayPasswords = printRandomPasswordsTable()

	} else if memorable {

		arrayPasswords = printMemorableTable()

	} else if randomHex {

		arrayPasswords = printRandomHexTable()

	} else if grammatical {

		arrayPasswords = printGrammaticalTable(false)

	} else if grammaticalAI {

		arrayPasswords = printGrammaticalTable(true)
	}

	return arrayPasswords
}

func colorizeCharactersUnix(password string, print bool) string {

	var coloredCharsString string

	// TODO: Create flag to trim the password down to the requestedPasswordLength
	//password = trimPassword(password, requestedPasswordLength)

	// Check each character's ascii value and colorize according to category
	//for i := 0; i < requestedPasswordLength; i++ {
	for i := 0; i < len(password); i++ {

		// Convert the character back to ascii value for the color assignment
		character := int32(password[i])

		if character >= 65 && character <= 90 {

			// Assign a color to uppercase characters
			coloredCharsString += color.WhiteString(string(character))

		} else if character >= 97 && character <= 122 {

			// Assign a color to lowercase characters
			coloredCharsString += color.HiGreenString(string(character))

		} else if character >= 48 && character <= 57 {

			// Assign a color to number characters
			coloredCharsString += color.CyanString(string(character))

		} else if character >= 33 && character <= 47 {

			// Assign a color to special characters, first range
			coloredCharsString += color.HiBlueString(string(character))

		} else if character >= 58 && character <= 64 {

			// Assign a color to special characters, second range
			coloredCharsString += color.HiBlueString(string(character))
		} else if character >= 91 && character <= 96 {

			// Assign a color to special characters, third range
			coloredCharsString += color.HiBlueString(string(character))

		} else if character >= 123 && character <= 126 {

			// Assign a color to special characters, fourth range
			coloredCharsString += color.HiBlueString(string(character))

		} else {

			// Assign a color to any character not represented above
			coloredCharsString += color.HiYellowString(string(character))
		}
	}

	if print == true {

		fmt.Print(coloredCharsString)
	}

	return coloredCharsString
}

// printPassphraseTable() generates an array of random passphrases and prints them
// in a table format to the console output.
//
// printPassphraseTable generates and prints a table of passphrases with
// colorized characters. The function sets the console height, instantiates a new
// table writer object, creates a new empty array with the same length as the
// original array, and loops through the console screen height and prints a table
// of passphrases. Each row of the table contains an index number and the
// passphrase. The function returns an array of passphrases because it's needed
// for the clipboard functions if the program is in interactive mode. Note: The
// colorization of characters works on all platforms but no color renders on
// Windows.
//
//	Returns:
//	 - An array of strings
func printPassphraseTable() []string {

	var consoleHeight int

	// Set the console height
	consoleHeight = funcName(consoleHeight)

	// Instantiate a new table writer object
	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(os.Stdout)

	// Create a new empty array with the same length as the original array
	// This avoids leftover empty array elements causing clipboard copy
	// failures later on.
	arrayOfPassphrases := make([]string, (consoleHeight/2)-1)

	// Loop through the console screen height and print a table of passphrases
	for i := 0; i < (consoleHeight/2)-1; i++ {

		passphrase := createPassphrase()

		// Colorize the passphrase that we're saving to the array
		// The following works on all platforms but no color renders on Windows
		passphrase = colorizeCharactersUnix(passphrase, false)

		// Append the passphrase to the array
		arrayOfPassphrases[i] = passphrase

		// Prepare color for the index number
		red := color.New(color.FgHiRed).SprintfFunc()

		// Print the index number and current element of the array
		tableWriter.AppendRow([]interface{}{red("%d", i), passphrase})

		tableWriter.AppendSeparator()
	}
	tableWriter.SetStyle(table.StyleLight)
	tableWriter.Render()

	// Return the array because it's needed for the
	// clipboard functions if we're in interactive mode.
	return arrayOfPassphrases
}

// printWordChainsTable() prints a table of word chains to the console screen and returns an array of word chains.
// The console screen height is determined by the funcName function, which is called to set the consoleHeight variable.
// An array of word chains is created with the same length as the original array to avoid leftover empty array elements
// causing clipboard copy failures later on.
//
// The function loops through the console screen height and prints a table of word chains by appending each word chain
// to the arrayOfWordChains array and displaying it with its corresponding index number in a table using the table writer object.
// The function also colorizes the word chain and its index number for better readability and uses the colorizeCharactersUnix function
// to colorize the word chain, which works on all platforms except Windows. The function then sets the table style to Light and
// renders the table.
//
// The function finally returns the arrayOfWordChains array because it's needed for the clipboard functions
// if the program is in interactive mode.
//
//	Returns:
//	  - An array of strings
func printWordChainsTable() []string {

	var consoleHeight int

	// Set the console height
	consoleHeight = funcName(consoleHeight)

	// Instantiate a new table writer object
	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(os.Stdout)

	// Create a new empty array with the same length as the original array
	// This avoids leftover empty array elements causing clipboard copy
	// failures later on.
	arrayOfWordChains := make([]string, consoleHeight/2)

	// Loop through the console screen height and print a table of word chains
	for i := 0; i < (consoleHeight/2)-1; i++ {

		wordChainNoColor := createWordChain(requestedPasswordLength)

		// Colorize the word chain that we're saving to the array
		// The following works on all platforms but no color renders on Windows
		wordChainColorized := colorizeCharactersUnix(wordChainNoColor, false)

		// Append the word chain to the array to be used by the clipboard if in interactive mode
		arrayOfWordChains[i] = wordChainNoColor

		// Prepare color for the index number
		red := color.New(color.FgHiRed).SprintfFunc()

		// Print the index number and current element of the array
		tableWriter.AppendRow([]interface{}{red("%d", i), wordChainColorized})

		tableWriter.AppendSeparator()
	}
	tableWriter.SetStyle(table.StyleLight)
	tableWriter.Render()

	// Return the array because it's needed for the
	// clipboard functions if we're in interactive mode.
	return arrayOfWordChains
}

func printMixedPasswordsTable(mixedPasswords bool, randomPasswords bool) []string {

	var consoleHeight int

	// Set the console height
	consoleHeight = funcName(consoleHeight)

	// Instantiate a new table writer object
	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(os.Stdout)

	// Create a new empty array with the same length as the original array
	// This avoids leftover empty array elements causing clipboard copy
	// failures later on.
	arrayOfMixedPasswords := make([]string, consoleHeight/2)

	// Loop through the console screen height and print a table of mixed passwords
	for i := 0; i < (consoleHeight/2)-1; i++ {

		mixedPasswordNoColor := createMixedPassword(mixedPasswords, randomPasswords, consoleHeight)

		// Colorize the mixed password that we're saving to the array
		// The following works on all platforms but no color renders on Windows
		mixedPasswordColorized := colorizeCharactersUnix(mixedPasswordNoColor, false)

		// Append the mixed password to the array to be used by the clipboard if in interactive mode
		arrayOfMixedPasswords[i] = mixedPasswordNoColor

		// Prepare color for the index number
		red := color.New(color.FgHiRed).SprintfFunc()

		// Print the index number and current element of the array
		tableWriter.AppendRow([]interface{}{red("%d", i), mixedPasswordColorized})

		tableWriter.AppendSeparator()
	}
	tableWriter.SetStyle(table.StyleLight)
	tableWriter.Render()

	// Return the array because it's needed for the
	// clipboard functions if we're in interactive mode.
	return arrayOfMixedPasswords
}

func printRandomPasswordsTable() []string {

	var consoleHeight int

	// Set the console height
	consoleHeight = funcName(consoleHeight)

	// Instantiate a new table writer object
	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(os.Stdout)

	// Create a new empty array with the same length as the original array
	// This avoids leftover empty array elements causing clipboard copy
	// failures later on.
	arrayOfRandomPasswords := make([]string, consoleHeight/2)

	// Loop through the console screen height and print a table of random passwords
	for i := 0; i < (consoleHeight/2)-1; i++ {

		randomPasswordNoColor := randStringPassword(requestedPasswordLength, false)

		// Colorize the random password that we're saving to the array
		// The following works on all platforms but no color renders on Windows
		randomPasswordColorized := colorizeCharactersUnix(randomPasswordNoColor, false)

		// Append the random password to the array to be used by the clipboard if in interactive mode
		arrayOfRandomPasswords[i] = randomPasswordNoColor

		// Prepare color for the index number
		red := color.New(color.FgHiRed).SprintfFunc()

		// Print the index number and current element of the array
		tableWriter.AppendRow([]interface{}{red("%d", i), randomPasswordColorized})

		tableWriter.AppendSeparator()
	}
	tableWriter.SetStyle(table.StyleLight)
	tableWriter.Render()

	// Return the array because it's needed for the
	// clipboard functions if we're in interactive mode.
	return arrayOfRandomPasswords
}

func printMemorableTable() []string {

	var consoleHeight int

	// Set the console height
	consoleHeight = funcName(consoleHeight)

	// Instantiate a new table writer object
	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(os.Stdout)

	// Create a new empty array with the same length as the original array
	// This avoids leftover empty array elements causing clipboard copy
	// failures later on.
	arrayOfMemorablePasswords := make([]string, consoleHeight/2)

	// Loop through the console screen height and print a table of memorable passwords
	for i := 0; i < (consoleHeight/2)-1; i++ {

		memorablePasswordNoColor := createMemorablePassword(requestedPasswordLength)

		// Colorize the word chain that we're saving to the array
		// The following works on all platforms but no color renders on Windows
		memorablePasswordColorized := colorizeCharactersUnix(memorablePasswordNoColor, false)

		// Append the memorable password to the array to be used by the clipboard if in interactive mode
		arrayOfMemorablePasswords[i] = memorablePasswordNoColor

		// Prepare color for the index number
		red := color.New(color.FgHiRed).SprintfFunc()

		// Print the index number and current element of the array
		tableWriter.AppendRow([]interface{}{red("%d", i), memorablePasswordColorized})

		tableWriter.AppendSeparator()
	}
	tableWriter.SetStyle(table.StyleLight)
	tableWriter.Render()

	// Return the array because it's needed for the
	// clipboard functions if we're in interactive mode.
	return arrayOfMemorablePasswords
}

func printRandomHexTable() []string {
	// TODO this function could be combined with printRandomPasswordTable() in some way

	var consoleHeight int

	// Set the console height
	consoleHeight = funcName(consoleHeight)

	// Instantiate a new table writer object
	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(os.Stdout)

	// Create a new empty array with the same length as the original array
	// This avoids leftover empty array elements causing clipboard copy
	// failures later on.
	arrayOfRandomHex := make([]string, consoleHeight/2)

	// Loop through the console screen height and print a table of random hex passwords
	for i := 0; i < (consoleHeight/2)-1; i++ {

		// TODO: Call a new function for randStringHex() here
		randomHexNoColor := randStringPassword(requestedPasswordLength, true)

		// Colorize the random hex passwords that we're saving to the array
		// The following works on all platforms but no color renders on Windows
		randomHexColorized := colorizeCharactersUnix(randomHexNoColor, false)

		// Append the random hex password to the array to be used by the clipboard if in interactive mode
		arrayOfRandomHex[i] = randomHexNoColor

		// Prepare color for the index number
		red := color.New(color.FgHiRed).SprintfFunc()

		// Print the index number and current element of the array
		tableWriter.AppendRow([]interface{}{red("%d", i), randomHexColorized})

		tableWriter.AppendSeparator()
	}
	tableWriter.SetStyle(table.StyleLight)
	tableWriter.Render()

	// Return the array because it's needed for the
	// clipboard functions if we're in interactive mode.
	return arrayOfRandomHex
}

func printPasswordTypesTable() []string {

	// Define a struct to hold a string pair
	type PasswordAndCommandFlag struct {
		PasswordExample string
		CommandFlag     string
	}

	// Instantiate a new table writer object
	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(os.Stdout)

	// Create a slice of PasswordAndCommandFlag structs
	var arrayOfPasswordTypes []PasswordAndCommandFlag

	// Random non-hex example password
	randStringPasswordExample := randStringPassword(requestedPasswordLength, false)
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: randStringPasswordExample,
			CommandFlag:     "--random",
		})

	// Random hex example password
	randHexPasswordExample := randStringPassword(requestedPasswordLength, true)
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: randHexPasswordExample,
			CommandFlag:     "--hex",
		})

	// Word chain example password
	wordChainPasswordExample := createWordChain(requestedPasswordLength)
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: wordChainPasswordExample,
			CommandFlag:     "--word-chains",
		})

	// TODO: This errors out if rows is < 8
	// Mixed password example password
	mixedPasswordExample := createMixedPassword(true, false, 8)
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: mixedPasswordExample,
			CommandFlag:     "--mixed",
		})

	// Passphrase example password
	passphraseExample := createPassphrase()
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: passphraseExample,
			CommandFlag:     "--passphrases",
		})

	// Memorable password example password
	var memorablePassword string
	//memorableExample := memorableTransformOne(memorablePassword, requestedPasswordLength)
	memorableExample := chooseMemorableTransform(memorablePassword, requestedPasswordLength)
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: memorableExample,
			CommandFlag:     "--memorable",
		})

	// Grammatical example password
	//memorablePassword = ""
	grammaticalExample := createGrammaticalPassword()
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: grammaticalExample,
			CommandFlag:     "--grammatical",
		})

	// Grammatical-AI example password
	nonSensicalSentence := createGrammaticalPassword()
	grammaticalExampleAI := createGrammaticalPasswordAI(nonSensicalSentence)
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: grammaticalExampleAI,
			CommandFlag:     "--grammatical-ai",
		})

	// Print the slice of string pairs
	for _, pair := range arrayOfPasswordTypes {

		exampleColorized := colorizeCharactersUnix(pair.PasswordExample, false)

		// Prepare color for the command color
		red := color.New(color.FgHiRed).SprintfFunc()

		// Print the command option and example password
		tableWriter.AppendRow([]interface{}{red("%s", pair.CommandFlag), exampleColorized})

		tableWriter.AppendSeparator()
	}

	tableWriter.SetStyle(table.StyleLight)
	tableWriter.Render()

	//Return the array because it's needed for the
	//clipboard functions if we're in interactive mode.
	//return arrayOfRandomHex
	return nil
}

func createGrammaticalPassword() string {

	var verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition string

	// The new way to seed randomness each time a function is called
	// Otherwise randomness is only seeded at the start of runtime
	randomnessObject := rand.New(rand.NewSource(time.Now().UnixNano()))

	/* SENTENCE ONE ---------------------------------------------
	Tightly reanimate one roof.#1a
	Bewitch his steel yesterday.#1b
	-------------------------------------------------------------*/
	var sentenceOne string

	randomChoice := randomnessObject.Intn(2)

	if randomChoice == 0 {

		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		sentenceOne = capitalizeFirstLetter(verb) + " " + article + " " + noun + " " + adverb + "."

	} else {

		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(adjective, article)

		// Include adjective
		sentenceOne = capitalizeFirstLetter(verb) + " " + article + " " + adjective + " " + noun + "."
	}

	/* SENTENCE TWO ---------------------------------------------
	That is aware.#2a
	Those are my pay.#2b
	Are those my pay?#2c
	-------------------------------------------------------------*/

	var sentenceTwo string

	randomnessObject = rand.New(rand.NewSource(time.Now().UnixNano()))

	// Randomly choose between 0 and 1
	randomChoice = randomnessObject.Intn(3)

	if randomChoice == 0 {

		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		//sentenceTwo = capitalizeFirstLetter(pronounAndVerbPresent) + " " + adjective + ".#2a"

		// Build the sentence
		sentenceTwo = pronounAndVerbPresent + " " + adjective + "."

		// 50% chance that it will be prepended with something like, "And then,"
		// 50% chance it will be unchanged
		sentenceTwo = maybePrependConjAdvPhrase(sentenceTwo)

		sentenceTwo = capitalizeFirstLetter(sentenceTwo)

	} else if randomChoice == 1 {

		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()
		// Change "a" to "an" if the following word begins with a vowel

		article = modifyArticle(noun, article)

		sentenceTwo = capitalizeFirstLetter(pronounAndVerbPresent) + " " + article + " " + noun + "."

	} else {

		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()
		// Change "a" to "an" if the following word begins with a vowel

		article = modifyArticle(noun, article)

		// Reverse the subject and verb: "I am" becomes "Am I"
		subjectVerbPhrase := strings.Split(pronounAndVerbPresent, " ")
		verbAndPronounPresent := subjectVerbPhrase[1] + " " + subjectVerbPhrase[0]

		sentenceTwo = capitalizeFirstLetter(verbAndPronounPresent) + " " + article + " " + noun + "?"

	}

	/* SENTENCE THREE -------------------------------------------
	Progress at their spring.#3a
	He is between someone's breath.#3b
	Is he between someone's breath?#3c
	-------------------------------------------------------------*/
	var sentenceThree string

	randomnessObject = rand.New(rand.NewSource(time.Now().UnixNano()))

	randomChoice = randomnessObject.Intn(3)

	if randomChoice == 0 {

		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		sentenceThree = capitalizeFirstLetter(verb) + " " + preposition + " " + article + " " + noun + "."

	} else if randomChoice == 1 {

		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		sentenceThree = capitalizeFirstLetter(pronounAndVerbPresent) + " " + preposition + " " + article + " " + noun + "."

	} else {

		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()
		// Change "a" to "an" if the following word begins with a vowel

		article = modifyArticle(noun, article)

		// Reverse the subject and verb: "I am" becomes "Am I"
		subjectVerbPhrase := strings.Split(pronounAndVerbPresent, " ")
		verbAndPronounPresent := subjectVerbPhrase[1] + " " + subjectVerbPhrase[0]

		sentenceThree = capitalizeFirstLetter(verbAndPronounPresent) + " " + preposition + " " + article + " " + noun + "?"

	}

	/* SENTENCE FOUR --------------------------------------------
	Didn't you yesterday criminalise the skill?#4a
	Don't intoxicate one boat.#4b
	-------------------------------------------------------------*/
	var sentenceFour string
	randomnessObject = rand.New(rand.NewSource(time.Now().UnixNano()))
	randomChoice = randomnessObject.Intn(2)
	if randomChoice == 0 {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()
		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		// TODO: Implement randomized pronouns
		pronoun := getRandomPronoun()
		sentenceFour = capitalizeFirstLetter("Didn't") + " " + pronoun + " " + adverb + " " + verb + " " + article + " " + noun + "?"
	} else {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()
		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		sentenceFour = capitalizeFirstLetter("Don't") + " " + verb + " " + article + " " + noun + "."
	}

	/* SENTENCE FIVE --------------------------------------------
	Someone's road is the dog's.#5a
	Their outlying heavy shall accredit.#5b
	-------------------------------------------------------------*/
	var sentenceFive string

	randomnessObject = rand.New(rand.NewSource(time.Now().UnixNano()))

	randomChoice = randomnessObject.Intn(2)

	if randomChoice == 0 {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()
		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		sentenceFive = capitalizeFirstLetter(article) + " " + noun + " is " + possessivePronoun + "."
	} else {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(adjective, article)

		// Check if it is an irregular verb and change verb tense if auxiliary verb requires it
		verb = convertIrregularVerb(auxVerb, verb)

		sentenceFive = capitalizeFirstLetter(article) + " " + adjective + " " + noun + " " + auxVerb + " " + verb + "."
	}

	/* SENTENCE SIX ---------------------------------------------
	Wasn't one mood cleverly operated?#6a
	Wouldn't his fine suit displace?#6b
	-------------------------------------------------------------*/
	var sentenceSix string
	randomnessObject = rand.New(rand.NewSource(time.Now().UnixNano()))
	randomChoice = randomnessObject.Intn(2)
	if randomChoice == 0 {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		// Check if it is an irregular verb and change verb tense if auxiliary verb requires it
		verb = convertIrregularVerb(auxVerb, verb)

		sentenceSix = capitalizeFirstLetter(auxVerb) + " " + article + " " + noun + " " + adverb + " " + verb + "?"
	} else {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(adjective, article)

		// Check if it is an irregular verb and change verb tense if auxiliary verb requires it
		verb = convertIrregularVerb(auxVerb, verb)

		// include adjective
		sentenceSix = capitalizeFirstLetter(auxVerb) + " " + article + " " + adjective + " " + noun + " " + verb + "?"
	}

	/* SENTENCE SEVEN -------------------------------------------
	Memorize my devil.#7a
	Any relief shall surrender.#7b
	-------------------------------------------------------------*/
	var sentenceSeven string
	randomnessObject = rand.New(rand.NewSource(time.Now().UnixNano()))
	randomChoice = randomnessObject.Intn(2)
	if randomChoice == 0 {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		sentenceSeven = capitalizeFirstLetter(verb) + " " + article + " " + noun + "."
	} else {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		// Check if it is an irregular verb and change verb tense if auxiliary verb requires it
		verb = convertIrregularVerb(auxVerb, verb)

		sentenceSeven = capitalizeFirstLetter(article) + " " + noun + " " + auxVerb + " " + verb + "."
	}

	/* SENTENCE EIGHT -------------------------------------------
	Rarely typecast one leave.#8a
	The quarter won't discriminate.#8b
	-------------------------------------------------------------*/
	verbModifier := getVerbModifier(randomnessObject)
	var sentenceEight string
	randomnessObject = rand.New(rand.NewSource(time.Now().UnixNano()))
	randomChoice = randomnessObject.Intn(2)
	if randomChoice == 0 {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		sentenceEight = capitalizeFirstLetter(verbModifier) + " " + verb + " " + article + " " + noun + "."
	} else {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		// Check if it is an irregular verb and change verb tense if auxiliary verb requires it
		verb = convertIrregularVerb(auxVerb, verb)

		//sentenceEight = capitalizeFirstLetter(article) + " " + noun + " " + auxVerb + " " + verb + ".#8b"

		// Build the sentence
		sentenceEight = article + " " + noun + " " + auxVerb + " " + verb + "."

		// 50% chance that it will be prepended with something like, "And then,"
		// 50% chance it will be unchanged
		sentenceEight = maybePrependConjAdvPhrase(sentenceEight)

		sentenceEight = capitalizeFirstLetter(sentenceEight)
	}

	/* SENTENCE NINE --------------------------------------------
	I denied any frightening wonder.#9a
	It carefully desensitized your pleasure.#9b
	-------------------------------------------------------------*/
	var sentenceNine string
	randomnessObject = rand.New(rand.NewSource(time.Now().UnixNano()))
	randomChoice = randomnessObject.Intn(2)
	if randomChoice == 0 {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()
		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		pronoun := getRandomPronoun()
		sentenceNine = capitalizeFirstLetter(pronoun) + " " + convertVerbToPastTense(verb) + " " + article + " " + adjective + " " + noun + "."
	} else {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()
		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		pronoun := getRandomPronoun()
		sentenceNine = capitalizeFirstLetter(pronoun) + " " + adverb + " " + convertVerbToPastTense(verb) + " " + article + " " + noun + "."
	}

	// TODO: Pluralize noun if auxVerb is were or weren't
	// TODO: Should "hasn't" be replaced with "hasn't been" when followed by a past-tense verb?
	// TODO: Detect double negatives and handle them somehow
	// TODO: Add interrogative sentences with modal auxiliary verbs, ending in a question mark.
	// TODO: Get better vocab lists

	randomSentenceIndex := rand.Intn(9)

	var randomSentenceStructure string

	switch randomSentenceIndex {
	case 0:
		randomSentenceStructure = sentenceOne
	case 1:
		randomSentenceStructure = sentenceTwo
	case 2:
		randomSentenceStructure = sentenceThree
	case 3:
		randomSentenceStructure = sentenceFour
	case 4:
		randomSentenceStructure = sentenceFive
	case 5:
		randomSentenceStructure = sentenceSix
	case 6:
		randomSentenceStructure = sentenceSeven
	case 7:
		randomSentenceStructure = sentenceEight
	case 8:
		randomSentenceStructure = sentenceNine
	}

	return randomSentenceStructure
}

func printGrammaticalTable(grammaticalAI bool) []string {

	var consoleHeight int

	// Set the console height
	consoleHeight = funcName(consoleHeight)

	// Instantiate a new table writer object
	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(os.Stdout)

	// Create a new empty array with the same length as the original array
	// This avoids leftover empty array elements causing clipboard copy
	// failures later on.
	arrayOfRandomHex := make([]string, consoleHeight/2)

	var randomSentenceNoColor string

	// Loop through the console screen height and print a table of random sentences
	for i := 0; i < (consoleHeight/2)-1; i++ {

		if grammaticalAI == false {

			randomSentenceNoColor = createGrammaticalPassword()

		} else {

			nonSensicalSentence := createGrammaticalPassword()
			// Use AI to improve the sentence we generated
			randomSentenceNoColor = createGrammaticalPasswordAI(nonSensicalSentence)
		}

		// Colorize the random sentences that we're saving to the array
		// The following works on all platforms but no color renders on Windows
		randomSentenceColorized := colorizeCharactersUnix(randomSentenceNoColor, false)

		// Append the random sentence to the array to be used by the clipboard if in interactive mode
		arrayOfRandomHex[i] = randomSentenceNoColor

		// Prepare color for the index number
		red := color.New(color.FgHiRed).SprintfFunc()

		// Print the index number and current element of the array
		tableWriter.AppendRow([]interface{}{red("%d", i), randomSentenceColorized})

		tableWriter.AppendSeparator()
	}
	tableWriter.SetStyle(table.StyleLight)
	tableWriter.Render()

	// Return the array because it's needed for the
	// clipboard functions if we're in interactive mode.
	return arrayOfRandomHex
}
