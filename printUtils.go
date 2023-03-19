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
	passPhrases bool) {

	//underlineRed := "_"
	//
	//color.NoColor = false
	//red := color.New(color.FgHiRed).PrintfFunc()
	//red("Testing colors on Windows")
	//red(strings.Repeat(underlineRed, requestedPasswordLength+2))
	//color.NoColor = true

	grey := color.New(color.FgCyan, color.Faint).SprintfFunc()

	underline := grey("─")

	fmt.Printf(
		"%s%s%s\n",
		grey("+────+"),
		strings.Repeat(underline, requestedPasswordLength+2),
		grey("+"),
	)

	// Loop to print rows of index numbers and passwords to the terminal screen
	for rowNumber := 0; rowNumber < ((rows / 2) - 1); rowNumber++ {

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

		if randomPasswords {

			// Fetch a new randomized password string of the specified length
			password := randStringPassword(requestedPasswordLength)

			arrayPasswords[rowNumber] = password

			// Colorize and print the password
			colorizeCharactersWindows(requestedPasswordLength, password)

		} else if wordChains {

			password := randomWordChain(requestedPasswordLength)

			arrayPasswords[rowNumber] = password

		} else if mixedPasswords {

			password := ifMixedPasswords(mixedPasswords, randomPasswords, rows)

			arrayPasswords[rowNumber] = password

			// Colorize and print the password
			colorizeCharactersWindows(requestedPasswordLength, password)
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
}

func colorizeCharactersWindows(requestedPasswordLength int, password string) {

	var coloredCharsString string

	// TODO: Trim the password down to the requestedPasswordLength
	password = trimPassword(password, requestedPasswordLength)

	// Check each character's ascii value and colorize according to category
	for i := 0; i < requestedPasswordLength; i++ {

		// Convert the character back to ascii value for the color assignment
		character := int32(password[i])

		// TODO: Make uppercase and lowercase Hi colors and others not Hi.

		if character >= 65 && character <= 90 {

			// Assign a color to uppercase characters
			color.NoColor = false
			uppercaseColorPrint := color.New(color.FgCyan).PrintfFunc()
			uppercaseColorPrint(string(character))
			color.NoColor = true

		} else if character >= 97 && character <= 122 {

			// Assign a color to lowercase characters
			color.NoColor = false
			lowercaseColorPrint := color.New(color.FgGreen).PrintfFunc()
			lowercaseColorPrint(string(character))
			color.NoColor = true

		} else if character >= 48 && character <= 57 {

			// Assign a color to number characters
			color.NoColor = false
			numberColorPrint := color.New(color.FgHiCyan).PrintfFunc()
			numberColorPrint(string(character))
			color.NoColor = true

		} else if character >= 33 && character <= 47 {

			if character == 37 {

				// Double the % sign or printf thinks it is a formatting symbol
				color.NoColor = false
				specialCharColorPrint := color.New(color.FgHiGreen).PrintfFunc()
				specialCharColorPrint("%%")
				color.NoColor = true

			} else {

				// Assign a color to special characters, first range
				color.NoColor = false
				specialCharColorPrint := color.New(color.FgHiGreen).PrintfFunc()
				specialCharColorPrint(string(character))
				color.NoColor = true

			}
		} else if character >= 58 && character <= 64 {

			// Assign a color to special characters, second range
			color.NoColor = false
			specialCharColorPrint := color.New(color.FgHiGreen).PrintfFunc()
			specialCharColorPrint(string(character))
			color.NoColor = true

		} else if character >= 91 && character <= 96 {

			// Assign a color to special characters, third range
			color.NoColor = false
			specialCharColorPrint := color.New(color.FgHiGreen).PrintfFunc()
			specialCharColorPrint(string(character))
			color.NoColor = true

		} else if character >= 123 && character <= 126 {

			// Assign a color to special characters, fourth range
			color.NoColor = false
			specialCharColorPrint := color.New(color.FgHiGreen).PrintfFunc()
			specialCharColorPrint(string(character))
			color.NoColor = true

		} else {

			// Assign a color to any character not represented above
			color.NoColor = false
			specialCharColorPrint := color.New(color.FgHiYellow).PrintfFunc()
			specialCharColorPrint(string(character))
			color.NoColor = true

		}
		color.NoColor = true
	}

	fmt.Print(coloredCharsString)
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
func printPasswordTableUnix(
	rows int,
	requestedPasswordLength int,
	arrayPasswords []string,
	randomPasswords bool,
	wordChains bool,
	mixedPasswords bool,
	passPhrases bool) []string {

	grey := color.New(color.FgCyan, color.Faint).SprintfFunc()
	underline := grey("─")

	// Loop to print rows of index numbers and passwords to the terminal screen
	for rowNumber := 0; rowNumber < ((rows / 2) - 1); rowNumber++ {

		if !passPhrases {

			fmt.Printf(
				"%s%s%s\n",
				grey("+────+"),
				strings.Repeat(underline, requestedPasswordLength+2),
				grey("+"),
			)
			red := color.New(color.FgRed).SprintFunc()
			rowNumberString := fmt.Sprintf("%02d", rowNumber)

			// Print an index number for each printed password
			color.NoColor = false
			fmt.Printf("%s %s %s ", grey("│"), red(rowNumberString), grey("│"))
		}

		if randomPasswords {

			// Fetch a new randomized password string of the specified length
			password := randStringPassword(requestedPasswordLength)

			arrayPasswords[rowNumber] = password

			// Colorize and print the password
			colorizeCharactersUnix(requestedPasswordLength, password)

		} else if wordChains {

			password := randomWordChain(requestedPasswordLength)

			arrayPasswords[rowNumber] = password

		} else if mixedPasswords {

			password := ifMixedPasswords(mixedPasswords, randomPasswords, rows)

			arrayPasswords[rowNumber] = password

			// Colorize and print the password
			colorizeCharactersUnix(requestedPasswordLength, password)
		} else if passPhrases {

			password := createPassphrase()

			arrayPasswords[rowNumber] = password

			//printPassphraseTable(arrayPasswords)
			break
			// Colorize and print the password
			//colorizeCharactersUnix(requestedPasswordLength, password)
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
	}

	return arrayPasswords
}

func printTopOfTable(arrayPasswords []string) (func(format string, a ...interface{}) string, string) {
	// Initialize the maximum length to zero
	maxPasswordLength := 0

	// Loop through the array and update the maximum length if necessary
	for _, password := range arrayPasswords {
		length := len(password)
		if length > maxPasswordLength {
			maxPasswordLength = length
		}
	}

	grey := color.New(color.FgCyan, color.Faint).SprintfFunc()

	underline := grey("─")

	fmt.Printf(
		"%s%s%s\n",
		grey("+────+"),
		strings.Repeat(underline, maxPasswordLength+2),
		grey("+"),
	)
	return grey, underline
}

func colorizeCharactersUnix(requestedPasswordLength int, password string) {

	var coloredCharsString string

	// TODO: Default to making table expand for longer word-chains
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

	fmt.Print(coloredCharsString)

}

func printPassphraseTable() []string {

	consoleHeight, _ := consoleSizeUnix()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	// create a new empty array with the same length as the original array
	arrayOfPassphrases := make([]string, (consoleHeight/2)-1)

	for i := 0; i < (consoleHeight/2)-1; i++ {

		passphrase := createPassphrase()

		arrayOfPassphrases[i] = passphrase

		// print the current element of the array
		t.AppendRow([]interface{}{i, passphrase})
		t.AppendSeparator()
	}
	t.SetStyle(table.StyleLight)
	t.Render()

	return arrayOfPassphrases
}
