package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/jedib0t/go-pretty/v6/table"
	"os"
	"strings"
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
func printPasswordTableUnix(arrayPasswords []string, randomPasswords bool, wordChains bool, mixedPasswords bool, passPhrases bool, memorable bool, randomHex bool, grammatical bool, grammaticalAI bool, grammaticalAIWithNumbers bool) []string {

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

		arrayPasswords = printGrammaticalTable(false, false)

	} else if grammaticalAI {

		arrayPasswords = printGrammaticalTable(true, false)

	} else if grammaticalAIWithNumbers {

		arrayPasswords = printGrammaticalTable(true, true)
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
	grammaticalExampleAI := createGrammaticalPasswordAI(nonSensicalSentence, false)
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
	arrayOfRandomHex := make([]string, consoleHeight/2)

	var randomSentenceNoColor string

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
