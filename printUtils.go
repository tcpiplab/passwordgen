package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/vbauerster/mpb/v7"
	"os"
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
func printPasswordTableUnix(arrayPasswords []string, randomPasswords bool, wordChains bool, memorable2 bool, passPhrases bool, memorable bool, randomHex bool, grammatical bool, grammaticalAI bool, grammaticalAIWithNumbers bool, mnemonic bool, memorable3 bool, memorable4 bool) []string {

	if passPhrases {

		arrayPasswords = printPassphraseTable(requestedPasswordLength)

	} else if wordChains {

		arrayPasswords = printWordChainsTable()

	} else if memorable2 {

		arrayPasswords = printMemorable2Table(memorable2, randomPasswords)

	} else if randomPasswords {

		arrayPasswords = printRandomPasswordsTable()

	} else if memorable {

		arrayPasswords = printMemorableTable(1)

	} else if randomHex {

		arrayPasswords = printRandomHexTable()

	} else if grammatical {

		arrayPasswords = printGrammaticalTable(false, false)

	} else if grammaticalAI {

		arrayPasswords = printGrammaticalTable(true, false)

	} else if grammaticalAIWithNumbers {

		arrayPasswords = printGrammaticalTable(true, true)

	} else if mnemonic {

		arrayPasswords = printMnemonicTable()

	} else if memorable3 {

		arrayPasswords = printMemorableTable(3)

	} else if memorable4 {

		arrayPasswords = printMemorableTable(4)
	}

	return arrayPasswords
}

// colorizeCharactersUnix This function colorizes a given password string
// according to its characters' ASCII values.
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
func printPassphraseTable(requestedPasswordLength int) []string {

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

		passphrase := createPassphrase(requestedPasswordLength)

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

	//fmt.Printf("In printWordChainsTable() before loop.\nrequestedPasswordLength == '%d'\n", requestedPasswordLength)

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

// printMemorable2Table This function prints a table of memorable and random
// passwords of the specified console height with their indexes.
func printMemorable2Table(memorable2 bool, randomPasswords bool) []string {

	var consoleHeight int

	// Set the console height
	consoleHeight = funcName(consoleHeight)

	// Instantiate a new table writer object
	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(os.Stdout)

	// Create a new empty array with the same length as the original array
	// This avoids leftover empty array elements causing clipboard copy
	// failures later on.
	arrayOfMemorable2Passwords := make([]string, consoleHeight/2)

	// Loop through the console screen height and print a table of memorable2 passwords
	for i := 0; i < (consoleHeight/2)-1; i++ {

		memorable2PasswordNoColor := createMemorable2Password(memorable2, randomPasswords, consoleHeight)

		// Colorize the memorable2 password that we're saving to the array
		// The following works on all platforms but no color renders on Windows
		memorable2PasswordColorized := colorizeCharactersUnix(memorable2PasswordNoColor, false)

		// Append the memorable2 password to the array to be used by the clipboard if in interactive mode
		arrayOfMemorable2Passwords[i] = memorable2PasswordNoColor

		// Prepare color for the index number
		red := color.New(color.FgHiRed).SprintfFunc()

		// Print the index number and current element of the array
		tableWriter.AppendRow([]interface{}{red("%d", i), memorable2PasswordColorized})

		tableWriter.AppendSeparator()
	}
	tableWriter.SetStyle(table.StyleLight)
	tableWriter.Render()

	// Return the array because it's needed for the
	// clipboard functions if we're in interactive mode.
	return arrayOfMemorable2Passwords
}

// printRandomPasswordsTable This function prints a table of random passwords and
// returns an array of them to be used for clipboard functions if needed.
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

// printMemorableTable This function prints a table of memorable passwords
// of types 3 or 4, which are then stored in an array and returned for further use.
func printMemorableTable(memorableType int) []string {

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

	var memorablePasswordNoColor string

	// Define the total number of iterations. This will be used by the progress bar
	totalIterations := (consoleHeight / 2) - 1

	// Create a new progress bar container
	progressBarContainer := mpb.New()

	// Create a progress bar called progressBar
	progressBar := createProgressBar(progressBarContainer, totalIterations)

	// Loop through the console screen height and print a table of memorable passwords
	for i := 0; i < (consoleHeight/2)-1; i++ {

		if memorableType == 1 {

			memorablePasswordNoColor = createMemorablePassword(requestedPasswordLength)

		} else if memorableType == 3 {

			memorablePasswordNoColor = createMemorable3Password()

		} else if memorableType == 4 {

			memorablePasswordNoColor = createMemorablePasswordAI()
		}

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

		// Increment the progress progressBar
		progressBar.Increment()
	}

	// Wait for the progress progressBar to finish rendering
	progressBarContainer.Wait()

	tableWriter.SetStyle(table.StyleLight)
	tableWriter.Render()

	// Return the array because it's needed for the
	// clipboard functions if we're in interactive mode.
	return arrayOfMemorablePasswords
}

// printRandomHexTable This function provides an efficient way to generate an
// array of random hex passwords and print it as a table.
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

// printPasswordExamplesTable This function generates a table of example usages for
// different types of passwords, along with their corresponding command flag.
func printPasswordExamplesTable() []string {

	/*---------------------------------------------------------------------------
	Set up the progress bar
	---------------------------------------------------------------------------*/

	// Define the total number of iterations. This will be used by the progress bar
	// Here it is hardcoded for the number of password examples printed
	totalIterations := 10

	// Create a new progress bar container
	progressBarContainer := mpb.New()

	// Create a progress bar called progressBar
	progressBar := createProgressBar(progressBarContainer, totalIterations)

	/*---------------------------------------------------------------------------
	Set up the table to be printed and data structures to populate it
	---------------------------------------------------------------------------*/

	// Define a struct to hold a string pair that will populate each row of the table
	type PasswordAndCommandFlag struct {
		PasswordExample string
		CommandFlag     string
	}

	// Create a slice of PasswordAndCommandFlag structs
	var arrayOfPasswordTypes []PasswordAndCommandFlag

	// Instantiate a new table writer object
	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(os.Stdout)

	/*---------------------------------------------------------------------------
	Create pairs of each password type and its corresponding command flag.
	Increment the progress bar each time
	---------------------------------------------------------------------------*/

	// --grammatical -----------------------------------------------------------
	grammaticalExample := createGrammaticalPassword()
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: grammaticalExample,
			CommandFlag:     "--grammatical",
		})
	progressBar.Increment()
	// -------------------------------------------------------------------------

	// --grammatical-ai --------------------------------------------------------
	nonSensicalSentence := createGrammaticalPassword()
	grammaticalExampleAI := createGrammaticalPasswordAI(nonSensicalSentence, false)
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: grammaticalExampleAI,
			CommandFlag:     "--grammatical-ai",
		})
	progressBar.Increment()
	// ------------------------------------------------------------------------

	// --grammatical-ai-with-numbers ------------------------------------------
	nonSensicalSentence2 := createGrammaticalPassword()
	grammaticalExampleAIWithNumbers := createGrammaticalPasswordAI(nonSensicalSentence2, true)
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: grammaticalExampleAIWithNumbers,
			CommandFlag:     "--grammatical-ai-with-numbers",
		})
	progressBar.Increment()
	// -------------------------------------------------------------------------

	// -- hex ------------------------------------------------------------------
	randHexPasswordExample := randStringPassword(requestedPasswordLength, true)
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: randHexPasswordExample,
			CommandFlag:     "--hex",
		})
	progressBar.Increment()
	// --------------------------------------------------------------------------

	// --memorable --------------------------------------------------------------
	var memorablePassword string
	memorableExample := chooseMemorableTransform(memorablePassword, requestedPasswordLength)
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: memorableExample,
			CommandFlag:     "--memorable",
		})
	progressBar.Increment()
	// --------------------------------------------------------------------------

	// --memorable-2 ------------------------------------------------------------
	// TODO: This errors out if rows is < 8
	memorable2PasswordExample := createMemorable2Password(true, false, 8)
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: memorable2PasswordExample,
			CommandFlag:     "--memorable-2",
		})
	progressBar.Increment()
	// --------------------------------------------------------------------------

	// --memorable-3 ------------------------------------------------------------
	memorable3PasswordExample := createMemorable3Password()
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: memorable3PasswordExample,
			CommandFlag:     "--memorable-3",
		})
	progressBar.Increment()
	// --------------------------------------------------------------------------

	// --memorable-4 ------------------------------------------------------------
	memorable4PasswordExample := createMemorablePasswordAI()
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: memorable4PasswordExample,
			CommandFlag:     "--memorable-4",
		})
	progressBar.Increment()
	// --------------------------------------------------------------------------

	// --mnemonic ---------------------------------------------------------------
	var arrMnemonicPair [2]string
	arrMnemonicPair = printMnemonicExampleRow()
	// The --mnemonic example has an extra "\n" in the row
	mnemonicExample := arrMnemonicPair[0] + "\n" + arrMnemonicPair[1]
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: mnemonicExample,
			CommandFlag:     "--mnemonic",
		})
	progressBar.Increment()
	// --------------------------------------------------------------------------

	// --passphrases ------------------------------------------------------------
	// is hardcoded to 5 here to match its default value
	passphraseExample := createPassphrase(5)
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: passphraseExample,
			CommandFlag:     "--passphrases",
		})
	progressBar.Increment()
	// --------------------------------------------------------------------------

	// --random -----------------------------------------------------------------
	randStringPasswordExample := randStringPassword(requestedPasswordLength, false)
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: randStringPasswordExample,
			CommandFlag:     "--random",
		})
	progressBar.Increment()
	// --------------------------------------------------------------------------

	// --word-chains ------------------------------------------------------------
	wordChainPasswordExample := createWordChain(requestedPasswordLength)
	arrayOfPasswordTypes = append(arrayOfPasswordTypes,
		PasswordAndCommandFlag{
			PasswordExample: wordChainPasswordExample,
			CommandFlag:     "--word-chains",
		})
	progressBar.Increment()
	// --------------------------------------------------------------------------

	/*---------------------------------------------------------------------------
	Print the table
	---------------------------------------------------------------------------*/

	// Print the slice of string pairs
	for _, pair := range arrayOfPasswordTypes {

		exampleColorized := colorizeCharactersUnix(pair.PasswordExample, false)

		// Prepare color for the command color
		red := color.New(color.FgHiRed).SprintfFunc()

		// Print the command option and example password
		tableWriter.AppendRow([]interface{}{red("%s", pair.CommandFlag), exampleColorized})

		tableWriter.AppendSeparator()
	}

	// Wait for the progress progressBar to finish rendering
	progressBarContainer.Wait()

	// Render the table
	tableWriter.SetStyle(table.StyleLight)
	tableWriter.Render()

	//Return the array because it's needed for the
	//clipboard functions if we're in interactive mode.
	//return arrayOfRandomHex
	return nil
}

// printGrammaticalTable generates and prints a table of randomly generated
// sentences which have been colorized and can optionally be improved with AI
// and the AI sentences can optionally contain numbers. Currently, there is no
// option for non-AI sentences containing numbers. The function returns an array
// containing the generated sentences.
func printGrammaticalTable(grammaticalAI bool, grammaticalAIWithNumbers bool) []string {

	var consoleHeight int

	// Set the console height
	consoleHeight = funcName(consoleHeight)

	// Instantiate a new table writer object
	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(os.Stdout)

	// Create a new empty array with the same length as the original array
	// This avoids leftover empty array elements causing clipboard copy
	// failures later on.
	arrayOfGrammatical := make([]string, consoleHeight/2)

	var randomSentenceNoColor string

	// Define the total number of iterations. This will be used by the progress bar
	totalIterations := (consoleHeight / 2) - 1

	// Create a new progress bar container
	progressBarContainer := mpb.New()

	// Create a progress bar called progressBar
	progressBar := createProgressBar(progressBarContainer, totalIterations)

	// Loop through the console screen height and print a table of random sentences
	for i := 0; i < (consoleHeight/2)-1; i++ {

		if grammaticalAI == false {

			randomSentenceNoColor = createGrammaticalPassword()

		} else {

			if grammaticalAIWithNumbers == true {

				nonSensicalSentence := createGrammaticalPassword()

				// Have AI rewrite the sentence once
				nonSensicalSentence = createGrammaticalPasswordAI(nonSensicalSentence, false)

				// Then use AI again to improve the sentence we generated but also have it add numbers
				randomSentenceNoColor = createGrammaticalPasswordAI(nonSensicalSentence, grammaticalAIWithNumbers)

			} else if grammaticalAIWithNumbers == false {

				nonSensicalSentence := createGrammaticalPassword()

				// Use AI to improve the sentence we generated
				randomSentenceNoColor = createGrammaticalPasswordAI(nonSensicalSentence, grammaticalAIWithNumbers)
			}
		}

		// Colorize the random sentences that we're saving to the array
		// The following works on all platforms but no color renders on Windows
		randomSentenceColorized := colorizeCharactersUnix(randomSentenceNoColor, false)

		// Append the random sentence to the array to be used by the clipboard if in interactive mode
		arrayOfGrammatical[i] = randomSentenceNoColor

		// Prepare color for the index number
		red := color.New(color.FgHiRed).SprintfFunc()

		// Print the index number and current element of the array
		tableWriter.AppendRow([]interface{}{red("%d", i), randomSentenceColorized})

		tableWriter.AppendSeparator()

		// Increment the progress progressBar
		progressBar.Increment()
	}

	// Wait for the progress progressBar to finish rendering
	progressBarContainer.Wait()

	tableWriter.SetStyle(table.StyleLight)
	tableWriter.Render()

	// Return the array because it's needed for the
	// clipboard functions if we're in interactive mode.
	return arrayOfGrammatical
}

func printMnemonicTable() []string {

	var consoleHeight int

	// Set the console height
	consoleHeight = funcName(consoleHeight)

	// Instantiate a new table writer object
	tableWriter := table.NewWriter()
	tableWriter.SetOutputMirror(os.Stdout)

	// Create a new empty array with the same length as the original array
	// This avoids leftover empty array elements causing clipboard copy
	// failures later on.
	arrayOfMnemonics := make([]string, consoleHeight/2)

	var randomSentenceNoColor string
	var arrayMnemonicAndSentence [2]string

	// Define the total number of iterations
	totalIterations := (consoleHeight / 2) - 1

	// Create a new progress container
	progressBarContainer := mpb.New()

	// Create a progress bar called progressBar
	progressBar := createProgressBar(progressBarContainer, totalIterations)

	// Loop through the console screen height and print a table of random sentences
	for i := 0; i < (consoleHeight/2)-1; i++ {

		nonSensicalSentence := createGrammaticalPassword()

		// Have AI rewrite the sentence once
		nonSensicalSentence = createGrammaticalPasswordAI(nonSensicalSentence, false)

		// Then use AI again to improve the sentence we generated but also have it add numbers
		randomSentenceNoColor = createGrammaticalPasswordAI(nonSensicalSentence, true)

		// Populate a two-element array containing the mnemonic password and the corresponding sentence. Example:
		// "Ivd79tdI?" and "I verbally described 79 treats, didn't I?"
		//
		//arrayMnemonicAndSentence[0] = createMnemonicPasswordAI(randomSentenceNoColor)
		arrayMnemonicAndSentence[0] = createMnemonicFromSentence(randomSentenceNoColor)
		arrayMnemonicAndSentence[1] = randomSentenceNoColor

		// Colorize the random sentences that we're saving to the array
		// The following works on all platforms but no color renders on Windows
		randomSentenceColorized := colorizeCharactersUnix(randomSentenceNoColor, false)
		mnemonicPasswordColorized := colorizeCharactersUnix(arrayMnemonicAndSentence[0], false)

		// TODO: Not sure how to handle the clipboard with this pair of values
		// Append the random sentence to the array to be used by the clipboard if in interactive mode
		arrayOfMnemonics[i] = randomSentenceNoColor

		// Prepare color for the index number
		red := color.New(color.FgHiRed).SprintfFunc()

		// Print the index number, password, and sentence
		tableWriter.AppendRow([]interface{}{red("%d", i), mnemonicPasswordColorized, randomSentenceColorized})

		tableWriter.AppendSeparator()

		// Increment the progress progressBar
		progressBar.Increment()
	}

	// Wait for the progress progressBar to finish rendering
	progressBarContainer.Wait()

	tableWriter.SetStyle(table.StyleLight)
	tableWriter.Render()

	// Return the array because it's needed for the
	// clipboard functions if we're in interactive mode.
	return arrayOfMnemonics
}

// printMnemonicExampleRow This function creates a randomly generated mnemonic
// password and its corresponding sentence.
func printMnemonicExampleRow() [2]string {

	var randomSentenceNoColor string
	var arrayMnemonicAndSentence [2]string

	nonSensicalSentence := createGrammaticalPassword()

	// Have AI rewrite the sentence once
	nonSensicalSentence = createGrammaticalPasswordAI(nonSensicalSentence, false)

	// Then use AI again to improve the sentence we generated but also have it add numbers
	randomSentenceNoColor = createGrammaticalPasswordAI(nonSensicalSentence, true)

	// Populate a two-element array containing the mnemonic password and the corresponding sentence. Example:
	// "Ivd79tdI?" and "I verbally described 79 treats, didn't I?"
	arrayMnemonicAndSentence[0] = createMnemonicFromSentence(randomSentenceNoColor)
	arrayMnemonicAndSentence[1] = randomSentenceNoColor

	return arrayMnemonicAndSentence
}
