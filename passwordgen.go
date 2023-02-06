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
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var selectedPasswordNumber int

func main() {

	help := flag.Bool("help", false, "./passwordgen n\n\nWhere n is the length of the password.")
	//flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	interactive := flag.Bool("interactive", false, "./passwordgen -interactive\n\n")
	flag.Parse()

	if len(os.Args) != 3 {

		color.HiRed("\nPlease provide a password length as an argument\nOr -h for help.\n\n")
		return
	}

	// Convert the requested length from string to int
	requestedPasswordLength, err := strconv.Atoi(os.Args[2])

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

	//var arrayPasswords [10]string

	arrayPasswords := make([]string, rows)

	// Fill the screen with passwords
	for rowNumber := 0; rowNumber < rows-1; rowNumber++ {

		// Fetch a new randomized password string of the specified length
		password := randString(requestedPasswordLength)

		// Print an index number for each printed password
		fmt.Printf("%02d ", rowNumber)

		arrayPasswords[rowNumber] = password

		//passwordColorized = colorizeCharacters(requestedPasswordLength, password)

		fmt.Printf("%s", arrayPasswords[rowNumber])

		colorizeCharacters(requestedPasswordLength, password)

		fmt.Printf("\n")

	}

	if ifInteractive(interactive, selectedPasswordNumber) {

		fmt.Print(arrayPasswords[selectedPasswordNumber])

		return
	}

}

func checkPasswordLength(requestedPasswordLength int, err error) bool {
	if int(requestedPasswordLength) < 10 {

		color.HiRed("\nPassword length must be 10 or longer.\n\n")
		return true
	}

	if err != nil {
		color.HiRed("Invalid password length argument")
		return true
	}
	return false
}

func colorizeCharacters(requestedPasswordLength int, password string) {

	// Check each character's ascii value and colorize according to category
	for i := 0; i < requestedPasswordLength; i++ {

		// Convert the character back to ascii value for the color assignment
		character := int32(password[i])

		if character >= 65 && character <= 90 {

			// Assign a color to uppercase characters
			fmt.Printf(strings.TrimRight(color.WhiteString(string(character)), "\n"))
			//character = strings.TrimRight(color.WhiteString(string(character)), "\n")
			//return color.WhiteString(string(character))

		} else if character >= 97 && character <= 122 {

			// Assign a color to lowercase characters
			fmt.Printf(strings.TrimRight(color.HiWhiteString(string(character)), "\n"))

		} else if character >= 48 && character <= 57 {

			// Assign a color to number characters
			fmt.Printf(strings.TrimRight(color.CyanString(string(character)), "\n"))

		} else if character >= 33 && character <= 47 {

			if character == 37 {

				// Double the % sign or printf thinks it is a formatting symbol
				fmt.Printf(strings.TrimRight(color.HiBlueString("%%"), "\n"))

			} else {

				// Assign a color to special characters, first range
				fmt.Printf(strings.TrimRight(color.HiBlueString(string(character)), "\n"))
			}

		} else if character >= 58 && character <= 64 {

			// Assign a color to special characters, second range
			fmt.Printf(strings.TrimRight(color.HiBlueString(string(character)), "\n"))

		} else if character >= 91 && character <= 96 {

			// Assign a color to special characters, third range
			fmt.Printf(strings.TrimRight(color.HiBlueString(string(character)), "\n"))

		} else if character >= 123 && character <= 126 {

			// Assign a color to special characters, fourth range
			fmt.Printf(strings.TrimRight(color.HiBlueString(string(character)), "\n"))

		} else {

			// Assign a color to any character not represented above
			fmt.Printf(strings.TrimRight(color.HiYellowString(string(character)), "\n"))
		}
	}
}

func ifInteractive(interactive *bool, selectedPasswordInt int) bool {
	if *interactive {

		// Declare a variable to store the user's choice of which password they select
		var passwordNumber int

		// Prompt the user to choose a password from the list
		fmt.Print("Enter an integer: ")

		// Accept user input and save it to passwordNumber
		_, err := fmt.Scan(&passwordNumber)

		// Print error and exit
		if err != nil {

			fmt.Printf("Error is %d", err)
			return true
		}

		// Print the user's chosen number
		//fmt.Printf("You entered number: %d", passwordNumber)

		// Set the global var to the entered number
		selectedPasswordNumber = passwordNumber

		return true
	}
	return false
}

func consoleSize() (int, int) {

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

		log.Fatal(err)
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

		log.Fatal(err)
	}

	// Convert width to an integer
	width, err := strconv.Atoi(heightAndWidthArray[1])

	// If the conversion to int errored out, log to the screen and exit
	if err != nil {

		log.Fatal(err)
	}

	return height, width
}

func randString(lengthOfRandString int) string {

	// Set allowed characters
	var allowedCharacters = []int32("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#^&*()[]{}%")

	// Make a list of type int32 of the length the user requested their passwords should be
	listOfInt32Characters := make([]int32, lengthOfRandString)

	for i := range listOfInt32Characters {

		// Grab random chars and put them in the list. But only from the set of allowed characters
		listOfInt32Characters[i] = allowedCharacters[rand.Intn(len(allowedCharacters))]
	}

	// Return a new random password string
	return string(listOfInt32Characters)
}
