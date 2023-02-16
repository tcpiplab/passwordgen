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
	"github.com/atotto/clipboard"
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

	//var arrayPasswords [10]string

	arrayPasswords := make([]string, rows)

	// Fill the screen with passwords
	for rowNumber := 0; rowNumber < rows-1; rowNumber++ {

		// Fetch a new randomized password string of the specified length
		password := randString(requestedPasswordLength)

		// Print an index number for each printed password
		fmt.Printf("%02d ", rowNumber)

		arrayPasswords[rowNumber] = password

		colorizeCharacters(requestedPasswordLength, password)

		fmt.Printf("\n")

	}

	if ifInteractive(interactive) {

		// TEMP print out the selected password
		//fmt.Print(arrayPasswords[selectedPasswordNumber])

		// Copy the selected password to the clipboard
		if copyToClipboard(erase, arrayPasswords) {

			return
		}

		return
	}

}

func copyToClipboard(erase *bool, arrayPasswords []string) bool {

	// Copy the selected password to the clipboard
	err := clipboard.WriteAll(arrayPasswords[selectedPasswordNumber])

	if err != nil {

		fmt.Println("Error:", err)

		return true
	}

	fmt.Println("Input has been copied to clipboard.")

	if *erase {

		b, done := eraseClipboard(true, err)
		if done {
			return b
		}
	}

	return false
}

func eraseClipboard(erase bool, err error) (bool, bool) {

	if erase {

		fmt.Println("Waiting for 60 seconds before clearing the clipboard.")

		// TODO: make this optional with a command line flag
		time.Sleep(60 * time.Second)

		// Clear the contents of the clipboard
		err = clipboard.WriteAll("")

		if err != nil {

			fmt.Println("Error:", err)

			return false, true
		}

		fmt.Println("Clipboard has been cleared.")

	}
	return false, false
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

	var coloredCharsString string

	// Check each character's ascii value and colorize according to category
	for i := 0; i < requestedPasswordLength; i++ {

		// Convert the character back to ascii value for the color assignment
		character := int32(password[i])

		if character >= 65 && character <= 90 {

			// Assign a color to uppercase characters
			coloredCharsString += color.WhiteString(string(character))

		} else if character >= 97 && character <= 122 {

			// Assign a color to lowercase characters
			coloredCharsString += color.HiWhiteString(string(character))

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

func ifInteractive(interactive *bool) bool {

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
