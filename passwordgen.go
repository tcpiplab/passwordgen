package main

// My first Golang code. It generates a screen full of random char passwords
// of a specified length. ChatGPT wrote the stub for me after this input:

//    please write me a command line tool, written in golang, that generates
//    passwords of random characters. The command line tool should allow the
//    user to request a specific length of the generated passwords. The
//    command line tool's output should generate exactly enough passwords to
//    fill the screen but not any further. For example, if the command line
//    tool is run in a terminal screen that is 30 rows high then the command
//    line tool should generate 29 passwords.

// I then grabbed a gist for the column size stuff. The url is inline, below.
// But I had to tweak both the ChatGPT code and the gist to get things working.

import (
	"github.com/fatih/color"
	_ "github.com/fatih/color"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// Declare global variables
var selectedPasswordNumber int
var requestedPasswordLength int
var OS string

// main() is the entry point of the application. It handles the command-line
// arguments, detects the operating system, gets the console size, generates
// passwords, prints them to the screen, and optionally copies the selected
// password to the clipboard.
func main() {

	OS = detectOS()

	interactive, erase, randomPasswords, wordChains, mixedPasswords, _ := argsHandler()
	//if *done {
	//	return
	//}

	// Convert the requested length from string to int
	// Length must be the last argument
	requestedPasswordLength, _ = strconv.Atoi(os.Args[len(os.Args)-1])

	// Check for password length and return errors if needed
	if checkPasswordLength(requestedPasswordLength) {
		return
	}

	// Seed the randomness
	rand.Seed(time.Now().UnixNano())

	// Get the height and width of the console
	var rowsColumns [2]int

	if OS == "darwin" || OS == "linux" || OS == "unix" {

		rowsColumns[0], rowsColumns[1] = consoleSizeUnix()

	} else if OS == "windows" {

		rowsColumns[0], rowsColumns[1] = consoleSizeWindows()

		// Disable color output on windows
		//NO_COLOR = "true"
		color.NoColor = true // disables colorized output
	}

	// We only need the number of rows
	var rows int
	rows = rowsColumns[0]

	// If the user wants word-chain passwords, check to see if we have
	// an available wordlist on their OS for seeding the API queries
	if *wordChains {

		// Need to do this for word-chains to work
		*randomPasswords = false
	}
	if *mixedPasswords {

		// Need to do this for mixed passwords to work
		*randomPasswords = false

		if OS == "windows" {

			color.NoColor = false
			color.HiRed("Mixed passwords are not yet implemented on Windows.")
			os.Exit(1)
		}
	}

	arrayPasswords := make([]string, rows)

	if OS == "darwin" || OS == "linux" || OS == "unix" {

		// Fill the screen with passwords
		printPasswordTableUnix(
			rows,
			requestedPasswordLength,
			arrayPasswords,
			*randomPasswords,
			*wordChains,
			*mixedPasswords)

	} else if OS == "windows" {

		// Fill the screen with passwords
		printPasswordTableWindows(
			rows,
			requestedPasswordLength,
			arrayPasswords,
			*randomPasswords,
			*wordChains,
			*mixedPasswords)
	}

	if ifInteractive(interactive, rows) {

		// Copy the selected password to the clipboard
		if copyToClipboard(erase, arrayPasswords) {

			return
		}

		return
	}

}

// ifMixedPasswords() generates a mixed password if mixedPasswords is true, and random passwords otherwise.
//
//	Parameters:
//	  mixedPasswords: A boolean indicating whether mixed passwords are requested.
//	  randomPasswords: A boolean indicating whether random passwords are requested.
//	  rows: An integer specifying the number of rows in the output.
//
//	Returns:
//	  A string containing the generated password.
func ifMixedPasswords(mixedPasswords bool, randomPasswords bool, rows int) string {

	var outputStr string

	if mixedPasswords {

		// Need to do this for mixed passwords to work
		randomPasswords = false

		if checkForWordList() {

			arrWords := selectSeedWords(rows / 2)

			var inputStr string

			if requestedPasswordLength < 12 {

				// For now just grab the first word in the array
				inputStr = randomCase(arrWords[0])

			} else if requestedPasswordLength <= 20 {

				inputStr = surroundString(
					surroundString(
						surroundString(
							arrWords[0]) + "-" + arrWords[1]))

			} else if requestedPasswordLength > 20 {

				inputStr = surroundString(
					surroundString(
						surroundString(
							arrWords[0])+"-"+arrWords[1]) + "-" + arrWords[2])
			}

			outputStr = createMixedPassword(inputStr)

		}

	}
	return outputStr
}
