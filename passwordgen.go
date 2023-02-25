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
	"flag"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/fatih/color"
	"math/rand"
	"os"
	"strconv"
	"time"
)

var selectedPasswordNumber int

func main() {

	help := flag.Bool(
		"help",
		false,
		"./passwordgen n\nWhere n is the length of the password.\nLength must be the last argument.",
	)
	//flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	// Interactive mode is the default
	interactive := flag.Bool(
		"interactive",
		true,
		"./passwordgen -interactive[=false]\n")

	erase := flag.Bool(
		"erase",
		true,
		"./passwordgen -erase[=false]\n")

	flag.Parse()

	// For now the length is mandatory and must be the last arg
	if len(os.Args) < 2 {

		color.HiRed("\nPlease provide a password length as an argument\nOr -h for help.\n")
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

func ifInteractive(interactive *bool, rows int) bool {

	if *interactive {

		// Declare a variable to store the user's choice of which password they select
		var passwordNumber int

		// Prompt the user to choose a password from the list
		fmt.Print("Enter an integer: ")

		for {
			// Accept user input and save it to passwordNumber
			// We don't need the number of args, which is the first returned value,
			// so just put that in '_'
			_, err := fmt.Scan(&passwordNumber)

			// Check if input is an integer. If not, re-prompt the user
			if err != nil {

				fmt.Printf("Error: Expected input to be an integer: %s", err)
				fmt.Printf("\nEnter an integer: ")
				continue
			}

			// Check if selected password number is in range
			if passwordNumber < 0 || passwordNumber >= (rows-1) {

				fmt.Printf("Error: Your selection is out of range")
				fmt.Printf("\nEnter an integer: ")
				continue
			}
			break

		}

		// Set the global var to the entered number
		selectedPasswordNumber = passwordNumber

		return true
	}
	return false
}
