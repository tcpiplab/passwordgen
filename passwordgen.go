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
	_ "github.com/fatih/color"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var selectedPasswordNumber int

var OS string

func main() {

	OS = detectOS()

	//fmt.Printf(OS)

	interactive, erase, done := argsHandler()
	if done {
		return
	}

	// Convert the requested length from string to int
	// Length must be the last argument
	requestedPasswordLength, err := strconv.Atoi(os.Args[len(os.Args)-1])

	// Check for password length and return errors if needed
	if checkPasswordLength(requestedPasswordLength, err) {
		return
	}

	// Seed the randomness
	rand.Seed(time.Now().UnixNano())

	// Get the height and width of the console
	var rowsColumns [2]int
	rowsColumns[0], rowsColumns[1] = consoleSize()

	// We only need the number of rows
	var rows int
	rows = rowsColumns[0]

	arrayPasswords := make([]string, rows)

	// Fill the screen with passwords
	printPasswordTable(rows, requestedPasswordLength, arrayPasswords)

	if ifInteractive(interactive, rows) {

		// Copy the selected password to the clipboard
		if copyToClipboard(erase, arrayPasswords) {

			return
		}

		return
	}

}
