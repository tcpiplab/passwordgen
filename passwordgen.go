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
	"os"
	"strconv"
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

	interactive, erase, randomPasswords, wordChains, mixedPasswords, _, passPhrases, memorable, randomHex, examples, grammatical, grammaticalAI, grammaticalAIWithNumbers := argsHandler()

	//if *done {
	//	return
	//}

	if !*passPhrases {

		// Convert the requested length from string to int
		// Length must be the last argument
		requestedPasswordLength, _ = strconv.Atoi(os.Args[len(os.Args)-1])

		// Check for password length and return errors if needed
		// For now the length is mandatory and must be the last arg
		if checkPasswordLength(requestedPasswordLength, randomHex) {
			return
		}
	}

	// Get the height and width of the console
	var rowsColumns [2]int

	if OS == "darwin" || OS == "linux" || OS == "unix" {

		rowsColumns[0], rowsColumns[1] = consoleSizeUnix()

	} else if OS == "windows" {

		rowsColumns[0], rowsColumns[1] = consoleSizeWindows()

		// Temporarily disable color output on Windows
		color.NoColor = true // disables colorized output
	}

	// We only need the number of rows
	var rows int
	rows = rowsColumns[0]

	if *wordChains {

		*randomPasswords = false
	}
	if *mixedPasswords {

		*randomPasswords = false
	}
	if *passPhrases {

		*randomPasswords = false
	}
	if *memorable {

		*randomPasswords = false
	}
	if *randomHex {

		*randomPasswords = false
	}
	if *examples {

		*randomPasswords = false
		printPasswordTypesTable()
		os.Exit(0)
	}
	if *grammatical {
		*randomPasswords = false
	}
	if *grammaticalAI {
		*randomPasswords = false
	}
	if *grammaticalAIWithNumbers {
		*randomPasswords = false
	}

	arrayPasswords := make([]string, rows)

	if OS == "darwin" || OS == "linux" || OS == "unix" {

		// Fill the screen with passwords
		arrayPasswords = printPasswordTableUnix(arrayPasswords, *randomPasswords, *wordChains, *mixedPasswords, *passPhrases, *memorable, *randomHex, *grammatical, *grammaticalAI, *grammaticalAIWithNumbers)

	} else if OS == "windows" {

		// Fill the screen with passwords
		arrayPasswords = printPasswordTableUnix(arrayPasswords, *randomPasswords, *wordChains, *mixedPasswords, *passPhrases, *memorable, *randomHex, *grammatical, *grammaticalAI, *grammaticalAIWithNumbers)
	}

	if ifInteractive(interactive, rows) {

		// Copy the selected password to the clipboard
		if copyToClipboard(erase, arrayPasswords) {

			return
		}

		return
	}

}
