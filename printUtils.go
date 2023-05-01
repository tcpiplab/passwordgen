package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"io/ioutil"
	"math/rand"
	"net/http"
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

	//// Memorable Three example password
	//memorablePassword = ""
	//memorableThreeExample := memorableTransformThree(memorablePassword, requestedPasswordLength)
	//arrayOfPasswordTypes = append(arrayOfPasswordTypes,
	//	PasswordAndCommandFlag{
	//		PasswordExample: memorableThreeExample,
	//		CommandFlag:     "--mem3",
	//	})
	//
	//// Memorable Four example password
	//memorablePassword = ""
	//memorableFourExample := memorableTransformFour(memorablePassword, requestedPasswordLength)
	//arrayOfPasswordTypes = append(arrayOfPasswordTypes,
	//	PasswordAndCommandFlag{
	//		PasswordExample: memorableFourExample,
	//		CommandFlag:     "--mem4",
	//	})
	//
	//// Memorable Five example password
	//memorablePassword = ""
	//memorableFiveExample := memorableTransformFive(memorablePassword, requestedPasswordLength)
	//arrayOfPasswordTypes = append(arrayOfPasswordTypes,
	//	PasswordAndCommandFlag{
	//		PasswordExample: memorableFiveExample,
	//		CommandFlag:     "--mem5",
	//	})
	//
	//// Memorable Six example password
	//memorablePassword = ""
	//memorableSixExample := memorableTransformSix(memorablePassword, requestedPasswordLength)
	//arrayOfPasswordTypes = append(arrayOfPasswordTypes,
	//	PasswordAndCommandFlag{
	//		PasswordExample: memorableSixExample,
	//		CommandFlag:     "--mem6",
	//	})
	//
	//// Memorable Seven example password
	//memorablePassword = ""
	//memorableSevenExample := memorableTransformSeven(memorablePassword, requestedPasswordLength)
	//arrayOfPasswordTypes = append(arrayOfPasswordTypes,
	//	PasswordAndCommandFlag{
	//		PasswordExample: memorableSevenExample,
	//		CommandFlag:     "--mem7",
	//	})

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

		sentenceOne = capitalizeFirstLetter(verb) + " " + article + " " + noun + " " + adverb + ".#1a"

	} else {

		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(adjective, article)

		// Include adjective
		sentenceOne = capitalizeFirstLetter(verb) + " " + article + " " + adjective + " " + noun + ".#1b"
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
		sentenceTwo = pronounAndVerbPresent + " " + adjective + ".#2a"

		// 50% chance that it will be prepended with something like, "And then,"
		// 50% chance it will be unchanged
		sentenceTwo = maybePrependConjAdvPhrase(sentenceTwo)

		sentenceTwo = capitalizeFirstLetter(sentenceTwo)

	} else if randomChoice == 1 {

		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()
		// Change "a" to "an" if the following word begins with a vowel

		article = modifyArticle(noun, article)

		sentenceTwo = capitalizeFirstLetter(pronounAndVerbPresent) + " " + article + " " + noun + ".#2b"

	} else {

		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()
		// Change "a" to "an" if the following word begins with a vowel

		article = modifyArticle(noun, article)

		// Reverse the subject and verb: "I am" becomes "Am I"
		subjectVerbPhrase := strings.Split(pronounAndVerbPresent, " ")
		verbAndPronounPresent := subjectVerbPhrase[1] + " " + subjectVerbPhrase[0]

		sentenceTwo = capitalizeFirstLetter(verbAndPronounPresent) + " " + article + " " + noun + "?#2c"

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

		sentenceThree = capitalizeFirstLetter(verb) + " " + preposition + " " + article + " " + noun + ".#3a"

	} else if randomChoice == 1 {

		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		sentenceThree = capitalizeFirstLetter(pronounAndVerbPresent) + " " + preposition + " " + article + " " + noun + ".#3b"

	} else {

		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()
		// Change "a" to "an" if the following word begins with a vowel

		article = modifyArticle(noun, article)

		// Reverse the subject and verb: "I am" becomes "Am I"
		subjectVerbPhrase := strings.Split(pronounAndVerbPresent, " ")
		verbAndPronounPresent := subjectVerbPhrase[1] + " " + subjectVerbPhrase[0]

		sentenceThree = capitalizeFirstLetter(verbAndPronounPresent) + " " + preposition + " " + article + " " + noun + "?#3c"

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
		sentenceFour = capitalizeFirstLetter("Didn't") + " " + pronoun + " " + adverb + " " + verb + " " + article + " " + noun + "?#4a"
	} else {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()
		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		sentenceFour = capitalizeFirstLetter("Don't") + " " + verb + " " + article + " " + noun + ".#4b"
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

		sentenceFive = capitalizeFirstLetter(article) + " " + noun + " is " + possessivePronoun + ".#5a"
	} else {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(adjective, article)

		// Check if it is an irregular verb and change verb tense if auxiliary verb requires it
		verb = convertIrregularVerb(auxVerb, verb)

		sentenceFive = capitalizeFirstLetter(article) + " " + adjective + " " + noun + " " + auxVerb + " " + verb + ".#5b"
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

		sentenceSix = capitalizeFirstLetter(auxVerb) + " " + article + " " + noun + " " + adverb + " " + verb + "?#6a"
	} else {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(adjective, article)

		// Check if it is an irregular verb and change verb tense if auxiliary verb requires it
		verb = convertIrregularVerb(auxVerb, verb)

		// include adjective
		sentenceSix = capitalizeFirstLetter(auxVerb) + " " + article + " " + adjective + " " + noun + " " + verb + "?#6b"
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

		sentenceSeven = capitalizeFirstLetter(verb) + " " + article + " " + noun + ".#7a"
	} else {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		// Check if it is an irregular verb and change verb tense if auxiliary verb requires it
		verb = convertIrregularVerb(auxVerb, verb)

		sentenceSeven = capitalizeFirstLetter(article) + " " + noun + " " + auxVerb + " " + verb + ".#7b"
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

		sentenceEight = capitalizeFirstLetter(verbModifier) + " " + verb + " " + article + " " + noun + ".#8a"
	} else {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()

		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		// Check if it is an irregular verb and change verb tense if auxiliary verb requires it
		verb = convertIrregularVerb(auxVerb, verb)

		//sentenceEight = capitalizeFirstLetter(article) + " " + noun + " " + auxVerb + " " + verb + ".#8b"

		// Build the sentence
		sentenceEight = article + " " + noun + " " + auxVerb + " " + verb + ".#8b"

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
		sentenceNine = capitalizeFirstLetter(pronoun) + " " + convertVerbToPastTense(verb) + " " + article + " " + adjective + " " + noun + ".#9a"
	} else {
		verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition = getVocabWords()
		// Change "a" to "an" if the following word begins with a vowel
		article = modifyArticle(noun, article)

		pronoun := getRandomPronoun()
		sentenceNine = capitalizeFirstLetter(pronoun) + " " + adverb + " " + convertVerbToPastTense(verb) + " " + article + " " + noun + ".#9b"
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

func getVocabWords() (string, string, string, string, string, string, string, string, string) {
	verb := getEnglishVocabWord("verb")
	//verb2 := getEnglishVocabWord("verb")
	noun := getEnglishVocabWord("noun")
	adverb := getEnglishVocabWord("adverb")
	adjective := getEnglishVocabWord("adjective")
	article := getRandomArticle()
	auxVerb := getRandomAuxVerb()
	pronounAndVerbPresent := getPronounAndVerbPresent()
	possessivePronoun := getPossessivePronoun()
	preposition := getPreposition()
	return verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition
}

func getPreposition() string {
	randomInt := rand.Intn(10) // Generate a random integer between 0 and 9

	var preposition string

	switch randomInt {
	case 0:
		preposition = "in"
	case 1:
		preposition = "on"
	case 2:
		preposition = "at"
	case 3:
		preposition = "over"
	case 4:
		preposition = "under"
	case 5:
		preposition = "between"
	case 6:
		preposition = "behind"
	case 7:
		preposition = "before"
	case 8:
		preposition = "after"
	case 9:
		preposition = "through"
	}
	return preposition
}

func getRandomPronoun() string {
	pronouns := []string{"he", "she", "they", "it", "I", "you", "we"}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(pronouns))
	return pronouns[randomIndex]
}

func getPossessivePronoun() string {

	randomInt := rand.Intn(10) // Generate a random integer between 0 and 9

	var possessivePronoun string

	switch randomInt {
	case 0:
		possessivePronoun = "mine"
	case 1:
		possessivePronoun = "yours"
	case 2:
		possessivePronoun = "his"
	case 3:
		possessivePronoun = "hers"
	case 4:
		possessivePronoun = "its"
	case 5:
		possessivePronoun = "ours"
	case 6:
		possessivePronoun = "theirs"
	case 7:
		possessivePronoun = "someone's"
	case 8:
		possessivePronoun = "nobody's"
	case 9:
		possessivePronoun = "the dog's"
	}
	return possessivePronoun
}

func getPronounAndVerbPresent() string {
	// Choose a pronounAndVerbPresent at random
	randomInt := rand.Intn(10) // Generate a random integer between 0 and 9

	var pronounAndVerbPresent string

	switch randomInt {
	case 0:
		pronounAndVerbPresent = "it is"
	case 1:
		pronounAndVerbPresent = "that is"
	case 2:
		pronounAndVerbPresent = "those are"
	case 3:
		pronounAndVerbPresent = "this is"
	case 4:
		pronounAndVerbPresent = "he is"
	case 5:
		pronounAndVerbPresent = "she is"
	case 6:
		pronounAndVerbPresent = "they are"
	case 7:
		pronounAndVerbPresent = "we are"
	case 8:
		pronounAndVerbPresent = "you are"
	case 9:
		pronounAndVerbPresent = "I am"
	}
	return pronounAndVerbPresent
}

func getVerbModifier(r *rand.Rand) string {
	// Generate a random number between 0 and 4 (inclusive).
	randomNumber := r.Intn(5)

	// Randomly choose a verb modifier using a switch statement.
	var verbModifier string
	switch randomNumber {
	case 0:
		verbModifier = "never"
	case 1:
		verbModifier = "always"
	case 2:
		verbModifier = "rarely"
	case 3:
		verbModifier = "sometimes"
	case 4:
		verbModifier = "often"
	default:
		verbModifier = "unknown"
	}
	return verbModifier
}

func getRandomAuxVerb() string {
	randomAuxVerbIndex := rand.Intn(15)

	var auxVerb string

	switch randomAuxVerbIndex {
	case 0:
		auxVerb = "wasn't"
	case 1:
		auxVerb = "is"
	case 2:
		auxVerb = "isn't"
	case 3:
		auxVerb = "was"
	case 4:
		auxVerb = "were"
	case 5:
		auxVerb = "will"
	case 6:
		auxVerb = "shall"
	case 7:
		auxVerb = "shall not"
	case 8:
		auxVerb = "won't" // contraction of "will not"
	case 9:
		auxVerb = "hasn't" // contraction of "has not"
	case 10:
		auxVerb = "didn't"
	case 11:
		auxVerb = "can't"
	case 12:
		auxVerb = "wouldn't"
	case 13:
		auxVerb = "shouldn't"
	case 14:
		auxVerb = "won't"
	}
	return auxVerb
}

func getRandomArticle() string {

	var article string

	randomIndex := rand.Intn(10)

	switch randomIndex {
	case 0:
		article = "a"
	case 1:
		article = "the"
	case 2:
		article = "one"
	case 3:
		article = "my"
	case 4:
		article = "your"
	case 5:
		article = "his"
	case 6:
		article = "her"
	case 7:
		article = "their"
	case 8:
		article = "someone's"
	case 9:
		article = "any"
	}
	return article
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

func createGrammaticalPasswordAI(nonSensicalSentence string) string {

	openaiAPIURL := "https://api.openai.com/v1/completions"

	type CompletionCreateArgs struct {
		Model       string  `json:"model"`
		Prompt      string  `json:"prompt"`
		MaxTokens   int     `json:"max_tokens"`
		Temperature float64 `json:"temperature"`
	}

	apiKey := os.Getenv("GPT_API_KEY")

	//promptSentence := "Change the subject in the following nonsensical sentence so that the subject and verb sound like they belong together: '" + nonSensicalSentence + "'"

	promptSentence := "Change the subject in the following nonsensical sentence so that it makes more sense. Change the adverb, adjective, noun, or verb if they don't sound like they belong together: '" + nonSensicalSentence + "'"

	data := CompletionCreateArgs{
		Model:       "text-davinci-003",
		Prompt:      promptSentence,
		MaxTokens:   12,
		Temperature: 0,
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return "Error marshaling JSON"
	}

	req, err := http.NewRequest("POST", openaiAPIURL, bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println("Error creating request:", err)
		return "Error creating request"
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return "Error making request"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "Error reading response body"
	}

	//fmt.Println("Response:", string(body))

	rewrittenSentence := extractGPTJson(string(body))

	fmt.Println(nonSensicalSentence)
	fmt.Println(rewrittenSentence)

	return rewrittenSentence
}

func extractGPTJson(jsonData string) string {
	//jsonData := `{"id":"cmpl-7BE7N12xt4CoVvr18Hl3vboDQvsSp","object":"text_completion","created":1682910233,"model":"text-davinci-003","choices":[{"text":"\n\nSteal someone's jacket.","index":0,"logprobs":null,"finish_reason":"stop"}],"usage":{"prompt_tokens":31,"completion_tokens":8,"total_tokens":39}}`

	var sentence string

	type Response struct {
		ID      string `json:"id"`
		Object  string `json:"object"`
		Created int64  `json:"created"`
		Model   string `json:"model"`
		Choices []struct {
			Text         string      `json:"text"`
			Index        int         `json:"index"`
			Logprobs     interface{} `json:"logprobs"`
			FinishReason string      `json:"finish_reason"`
		} `json:"choices"`
		Usage struct {
			PromptTokens     int `json:"prompt_tokens"`
			CompletionTokens int `json:"completion_tokens"`
			TotalTokens      int `json:"total_tokens"`
		} `json:"usage"`
	}

	var response Response
	err := json.Unmarshal([]byte(jsonData), &response)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return "Error unmarshaling JSON"
	}

	if len(response.Choices) > 0 {
		sentence = response.Choices[0].Text
		//fmt.Println("Extracted sentence:", sentence)

		// Remove two leading newline characters
		sentence = strings.TrimPrefix(sentence, "\n\n")

	} else {
		fmt.Println("No choices found in the JSON")
	}
	return sentence
}

// modifyArticle checks if the firstLetter variable is present in the vowels string.
// If it is and the article is "a", the function returns "an".
// In all other cases, the function returns the article unchanged.
func modifyArticle(followingWord, article string) string {

	firstLetter := followingWord[:1]

	// TODO: this function isn't working. Try print statements.

	if isVowel(firstLetter) && article == "a" {

		//fmt.Printf("----------\n%s, %s\n", article, followingWord)
		//fmt.Printf("firstLetter: %s\n", firstLetter)
		//fmt.Printf("article: %s\n----------\n", article)

		return "an"

	} else {

		return article
	}

	//return article
}

func isVowel(char string) bool {
	vowels := "aeiouAEIOU"

	if len(char) != 1 {
		return false
	}

	// Return true if the char is a vowel
	return strings.Contains(vowels, char)
}

func convertVerbToPastTense(verb string) string {
	// If the verb ends with 'e', just add 'd' to the end.
	if strings.HasSuffix(verb, "e") {
		return verb + "d"
	}

	// If the verb ends with a consonant followed by 'y', replace 'y' with 'ied'.
	if len(verb) >= 2 && strings.Contains("bcdfghjklmnpqrstvwxyz", string(verb[len(verb)-2])) && strings.HasSuffix(verb, "y") {
		return verb[:len(verb)-1] + "ied"
	}

	// For other verbs, just add 'ed' to the end.
	return verb + "ed"
}

func applyAuxiliaryVerb(auxVerb string, verbPresentTense string) string {
	auxVerb = strings.ToLower(auxVerb)
	verbPresentTense = strings.ToLower(verbPresentTense)

	switch auxVerb {
	case "had", "has", "was", "is", "were", "hadn't", "weren't", "hasn't", "wasn't", "isn't":
		return convertVerbToPastTense(verbPresentTense)
	default:
		return verbPresentTense
	}
}

func convertIrregularVerb(auxVerb string, verb string) string {
	switch strings.ToLower(verb) {
	case "be":
		verb = "was"
	case "begin":
		verb = "began"
	case "bite":
		verb = "bit"
	case "blow":
		verb = "blew"
	case "break":
		verb = "broke"
	case "bring":
		verb = "brought"
	case "build":
		verb = "built"
	case "buy":
		verb = "bought"
	case "catch":
		verb = "caught"
	case "choose":
		verb = "chose"
	case "come":
		verb = "came"
	case "cost":
		verb = "cost"
	case "cut":
		verb = "cut"
	case "do":
		verb = "did"
	case "draw":
		verb = "drew"
	case "drink":
		verb = "drank"
	case "drive":
		verb = "drove"
	case "eat":
		verb = "ate"
	case "fall":
		verb = "fell"
	case "feel":
		verb = "felt"
	case "fight":
		verb = "fought"
	case "find":
		verb = "found"
	case "fly":
		verb = "flew"
	case "forget":
		verb = "forgot"
	case "freeze":
		verb = "froze"
	case "get":
		verb = "got"
	case "give":
		verb = "gave"
	case "go":
		verb = "went"
	case "grow":
		verb = "grew"
	case "hang":
		verb = "hung"
	case "have":
		verb = "had"
	case "hear":
		verb = "heard"
	case "hide":
		verb = "hid"
	case "hit":
		verb = "hit"
	case "hold":
		verb = "held"
	case "hurt":
		verb = "hurt"
	case "keep":
		verb = "kept"
	case "know":
		verb = "knew"
	case "lead":
		verb = "led"
	case "leave":
		verb = "left"
	case "lend":
		verb = "lent"
	case "let":
		verb = "let"
	case "lie":
		verb = "-"
	case "light":
		verb = "lit"
	case "lose":
		verb = "lost"
	case "make":
		verb = "made"
	case "mean":
		verb = "meant"
	case "meet":
		verb = "met"
	case "pay":
		verb = "paid"
	case "put":
		verb = "put"
	case "read":
		verb = "read"
	case "ride":
		verb = "rode"
	case "ring":
		verb = "rang"
	case "rise":
		verb = "rose"
	case "run":
		verb = "ran"
	case "say":
		verb = "said"
	case "see":
		verb = "saw"
	case "sell":
		verb = "sold"
	case "send":
		verb = "sent"
	case "set":
		verb = "set"
	case "shake":
		verb = "shook"
	case "shine":
		verb = "shone"
	case "shoot":
		verb = "shot"
	case "show":
		verb = "showed"
	case "shut":
		verb = "shut"
	case "sing":
		verb = "sang"
	case "sink":
		verb = "sank"
	case "sit":
		verb = "sat"
	case "sleep":
		verb = "slept"
	case "slide":
		verb = "slid"
	case "speak":
		verb = "spoke"
	case "spend":
		verb = "spent"
	case "spin":
		verb = "spun"
	case "spread":
		verb = "spread"
	case "stand":
		verb = "stood"
	case "steal":
		verb = "stole"
	case "stick":
		verb = "stuck"
	case "sting":
		verb = "stung"
	case "strike":
		verb = "struck"
	case "swear":
		verb = "swore"
	case "sweep":
		verb = "swept"
	case "swim":
		verb = "swam"
	case "take":
		verb = "took"
	case "teach":
		verb = "taught"
	case "tear":
		verb = "tore"
	case "tell":
		verb = "told"
	case "think":
		verb = "thought"
	case "throw":
		verb = "threw"
	case "understand":
		verb = "understood"
	case "wake":
		verb = "woke"
	case "wear":
		verb = "wore"
	case "win":
		verb = "won"
	case "write":
		verb = "wrote"
	case "grind":
		verb = "ground"
	case "stop":
		verb = "stopped"
	default:
		// If not an irregular verb, do the standard conversion to past tense
		// if auxiliary verb requires it
		return applyAuxiliaryVerb(auxVerb, verb)
	}
	// return past tense version of irregular verb
	return verb
}

func getConjunctiveAdverbialPhrase() string {

	var conjunctiveAdverbialPhrase string

	phrases := []string{
		"And then,",
		"In addition,",
		"Therefore,",
		"However,",
		"Conversely,",
		"Meanwhile,",
		"Moreover,",
		"Nonetheless,",
		"Furthermore,",
		"On the other hand,",
		"For example,",
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(phrases))

	switch randomIndex {
	case 0:
		conjunctiveAdverbialPhrase = phrases[0]
	case 1:
		conjunctiveAdverbialPhrase = phrases[1]
	case 2:
		conjunctiveAdverbialPhrase = phrases[2]
	case 3:
		conjunctiveAdverbialPhrase = phrases[3]
	case 4:
		conjunctiveAdverbialPhrase = phrases[4]
	case 5:
		conjunctiveAdverbialPhrase = phrases[5]
	case 6:
		conjunctiveAdverbialPhrase = phrases[6]
	case 7:
		conjunctiveAdverbialPhrase = phrases[7]
	case 8:
		conjunctiveAdverbialPhrase = phrases[8]
	case 9:
		conjunctiveAdverbialPhrase = phrases[9]
	case 10:
		conjunctiveAdverbialPhrase = phrases[10]
	default:
		conjunctiveAdverbialPhrase = "Wait. Um,"
	}
	return conjunctiveAdverbialPhrase
}

// The modifySentence function generates a random float between 0 and 1. If the
// value is less than 0.5, it calls getConjunctiveAdverbialPhrase() and prepends
// the resulting phrase to the input sentence. Otherwise, it returns the original
// sentence.
func maybePrependConjAdvPhrase(sentence string) string {
	rand.Seed(time.Now().UnixNano())
	shouldModify := rand.Float64()

	if shouldModify < 0.5 {
		return getConjunctiveAdverbialPhrase() + " " + sentence
	}
	return sentence
}
