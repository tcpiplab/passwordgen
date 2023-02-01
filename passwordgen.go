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

func main() {

	help := flag.Bool("help", false, "./passwordgen n\n\nWhere n is the length of the password.")
	//flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	interactive := flag.Bool("interactive", false, "./passwordgen -interactive\n\n")
	flag.Parse()

	if *interactive {

		// fmt.Printf("Interactive mode is %T", interactive)
		//var input int

		// This is broken. Change it all per this URL:
		// https://stackoverflow.com/questions/55193141/how-can-i-take-input-from-user-in-golang-fmt-scan

		fmt.Print("Enter an integer: ")
		number, err := fmt.Scanln()
		if err != nil {
			fmt.Printf("Error is %d", err)
			return
		}
		fmt.Printf("You entered: %d", number)

		return
	}

	if len(os.Args) != 2 {
		color.HiRed("\nPlease provide a password length as an argument\nOr -h for help.\n\n")
		return
	}

	// Convert the requested length from string to int
	requestedPasswordLength, err := strconv.Atoi(os.Args[1])

	if int(requestedPasswordLength) < 10 {

		color.HiRed("\nPassword length must be 10 or longer.\n\n")
		return
	}

	if err != nil {
		color.HiRed("Invalid password length argument")
		return
	}

	rand.Seed(time.Now().UnixNano())

	var rowsColumns [2]int
	rowsColumns[0], rowsColumns[1] = consoleSize()

	var rows int
	rows = rowsColumns[0]

	for rowNumber := 0; rowNumber < rows-1; rowNumber++ {

		password := randString(requestedPasswordLength)

		fmt.Printf("%02d ", rowNumber)

		for i := 0; i < requestedPasswordLength-1; i++ {

			character := int32(password[i])

			if character >= 65 && character <= 90 {

				// Assign a color to uppercase characters
				fmt.Printf(strings.TrimRight(color.WhiteString(string(character)), "\n"))

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

		fmt.Printf("\n")
	}
}

func consoleSize() (int, int) {

	// Originally from https://gist.github.com/steinelu/aa9a5f402b584bc967eb216e054ceefb

	// Execute the shell command `stty size` which returns two integers: height and width of the user'heightAndWidthString terminal
	sttyCommand := exec.Command("stty", "size")

	// Specify the shell'heightAndWidthString STDIN so that Executing `stty size` will work
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
