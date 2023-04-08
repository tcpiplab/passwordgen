package main

import (
	"github.com/fatih/color"
	"log"
	"math"
	"os"
	"os/exec"
	"runtime"
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

// Set the Windows terminal dimensions manually
func consoleSizeWindows() (int, int) {

	// Set to Windows default values for now
	return 30, 120
}

func funcName(consoleHeight int) int {
	if OS == "darwin" || OS == "linux" || OS == "unix" {

		consoleHeight, _ = consoleSizeUnix()

	} else if OS == "windows" {

		consoleHeight, _ = consoleSizeWindows()

	}
	return consoleHeight
}

func detectOS() string {
	//fmt.Printf("Running on %s\n", runtime.GOOS)

	//OS = runtime.GOOS

	return runtime.GOOS
}

// checkPasswordLength() checks the length of a password and returns a boolean indicating
// whether the password length is valid or not. If the requested password length is less
// than 8 characters, it will print a red-colored error message and return true. If an
// error is passed in as the second argument, it will print a red-colored error message
// indicating that the password length argument is invalid and return true. Otherwise,
// it will return false to indicate that the password length is valid.
//
//	Parameters:
//	  - requestedPasswordLength: the length of the password to be checked
//	  - err: an error indicating if the password length argument is invalid
//
//	Returns:
//	  - a boolean value indicating if the password length is valid or not
func checkPasswordLength(requestedPasswordLength int, randomHex *bool) bool {

	if *randomHex {

		if int(requestedPasswordLength) < 4 {

			color.HiRed("\nHex PIN length must be 4 or longer.\n\n")
			return true
		}

	} else if !*randomHex {

		if int(requestedPasswordLength) < 8 {

			color.HiRed("\nPassword length must be 8 or longer.\n\n")
			return true
		}
	}

	//else if int(requestedPasswordLength) > 255 {
	//	// Use passwordLength as an integer
	//	requestedPasswordLength = 10
	//}

	return false
}

func isHighEntropy(s string) bool {
	entropy := 0.0
	counts := make(map[rune]int)

	// Count the number of occurrences of each character
	for _, r := range s {
		counts[r]++
	}

	// Calculate the entropy of the string
	for _, count := range counts {
		p := float64(count) / float64(len(s))
		entropy -= p * math.Log2(p)
	}

	// Check if the entropy is high enough
	return entropy >= math.Log2(float64(len(s)))-1
}

func testForEntropy() {
	s := "123123123" // Replace with the string you want to test
	if isHighEntropy(s) {
		println("The string has high entropy")
	} else {
		println("The string does not have high entropy")
	}
}
