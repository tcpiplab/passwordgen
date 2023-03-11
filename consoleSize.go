package main

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// consoleSizeUnix returns the height and width of the user's terminal console in number of characters.
// It executes the shell command `stty size` and reads its output to get the console size.
// If any errors occur during this process, it will log an error message to the screen and exit.
//
// Returns:
// - an integer representing the height of the console in number of characters
// - an integer representing the width of the console in number of characters
func consoleSizeUnix() (int, int) {

	// Originally from https://gist.github.com/steinelu/aa9a5f402b584bc967eb216e054ceefb

	// Execute the shell command `stty size` which returns two integers:
	// height and width of the user's heightAndWidthString terminal
	sttyCommand := exec.Command("stty", "size")

	// Specify the shell's heightAndWidthString STDIN so that executing
	// `stty size` will work
	sttyCommand.Stdin = os.Stdin

	// Execute the `stty size` command and save the output and any resulting error.
	heightAndWidthBytes, err := sttyCommand.Output()

	// If it errored heightAndWidthBytes, log to the screen and exit
	if err != nil {

		log.Fatal("Error trying to get the size of the terminal:", err)
	}

	// Save the height and width values as a string
	heightAndWidthString := string(heightAndWidthBytes)

	// Remove extra whitespace
	heightAndWidthString = strings.TrimSpace(heightAndWidthString)

	// Split height and width into an array of two strings
	heightAndWidthArray := strings.Split(heightAndWidthString, " ")

	// Convert height to an integer
	// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
	height, err := strconv.Atoi(heightAndWidthArray[0])

	// If the conversion to int errored out, log to the screen and exit
	if err != nil {

		log.Fatal("Error trying to convert terminal height to an integer:", err)
	}

	// Convert width to an integer
	width, err := strconv.Atoi(heightAndWidthArray[1])

	// If the conversion to int errored out, log to the screen and exit
	if err != nil {

		log.Fatal("Error trying to convert terminal width to an integer:", err)
	}

	return height, width
}

// Get the Windows terminal dimensions
func consoleSizeWindows() (int, int) {

	// Set to Windows default values for now
	return 30, 120
}
