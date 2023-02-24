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

// printPasswordTable prints a table of randomized passwords with index numbers to the terminal screen.
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
func printPasswordTable(rows int, requestedPasswordLength int, arrayPasswords []string) {

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

		// Fetch a new randomized password string of the specified length
		password := randString(requestedPasswordLength)

		red := color.New(color.FgRed).SprintFunc()

		rowNumberString := fmt.Sprintf("%02d", rowNumber)

		// Print an index number for each printed password
		fmt.Printf("%s %s %s ", grey("│"), red(rowNumberString), grey("│"))

		arrayPasswords[rowNumber] = password

		// Colorize and print the password
		colorizeCharacters(requestedPasswordLength, password)

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

// copyToClipboard copies the selected password to the system clipboard and
// optionally erases the clipboard contents based on the value of the erase parameter.
// The function takes two input parameters:
//
//	erase - a pointer to a boolean value indicating whether the clipboard should be cleared
//	arrayPasswords - a string array containing the available passwords to choose from
//
// The function returns a boolean value indicating whether an error occurred during the process:
//   - if an error occurs, the function returns true.
//   - if the operation is successful, the function returns false.
func copyToClipboard(erase *bool, arrayPasswords []string) (copyErroredOut bool) {

	// Copy the selected password to the clipboard
	err := clipboard.WriteAll(arrayPasswords[selectedPasswordNumber])

	if err != nil {

		fmt.Println("Error: copying password to clipboard:", err)

		return true
	}

	fmt.Println("Input has been copied to clipboard.")

	if *erase {

		clipboardData, clipboardCleared := eraseClipboard(true, err)

		if clipboardCleared {

			return clipboardData

		}
	}

	return false
}

// eraseClipboard clears the contents of the clipboard if the erase parameter is true.
// The function takes two input parameters:
//
//	erase - a boolean value indicating whether the clipboard should be cleared
//	err - an error value that will be updated during the clearing process
//
// The function returns two boolean values:
//
//	success - indicating whether the clipboard was cleared successfully
//	hasError - indicating whether an error occurred during the clearing process
func eraseClipboard(erase bool, err error) (success bool, hasError bool) {

	// If the value of the erase parameter is true
	if erase {

		fmt.Println("Waiting for 60 seconds before clearing the clipboard.")

		// Start the progress bar goroutine
		progressBarStartStopChannel := make(chan bool)
		go progressBar(progressBarStartStopChannel)

		time.Sleep(60 * time.Second)

		// Clear the contents of the clipboard
		err = clipboard.WriteAll("")

		// Send a value to the channel to stop the progress bar
		progressBarStartStopChannel <- true

		if err != nil {

			fmt.Println("Error: Unable to clear the clipboard:", err)

			// If there is an error during the clearing process, the function returns false
			// and true to indicate that the operation was not successful and that an error
			// occurred.
			return false, true
		}

		fmt.Println("Clipboard has been cleared.")

	}

	// If the erase parameter is false, the function simply returns
	// false and false, indicating that no action was taken.
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

// progressBar displays a progress bar in the terminal by printing a period
// every 500 milliseconds until it receives a value on the given channel. The
// function runs in a separate goroutine, so it can be executed concurrently
// with other parts of a Go program. The progress bar can be stopped by sending
// a value on the channel.
//
// Parameters:
// progressBarChannel - the channel to listen for stop signal on.
//
// Example:
//
//	progressBarChannel := make(chan bool)
//	go progressBar(progressBarChannel)
//
//	// Do some work
//	time.Sleep(60 * time.Second)
//
//	// Send a value to the channel to stop the progress bar
//	progressBarChannel <- true
func progressBar(progressBarChannel chan bool) {
	for {
		select {
		case <-progressBarChannel:
			// Stop the progress bar when the channel receives a value
			fmt.Printf("\n")
			return
		default:

			// Display a progress bar with 60 steps, each step taking 1 second.
			for i := 0; i <= 60; i++ {

				// For each step,
				//   1. Print a solid block character █ (Unicode \u2588).
				//   2. Then print the remaining-seconds countdown number (zero-padded
				//      to two digits with %02d).
				//   3. Then move the cursor back two spaces using the ASCII control sequence
				//      \u001B[2D which moves the cursor two characters to the left to overwrite
				//      the number that was printed in the previous step.
				// We stay on one line the whole time.
				fmt.Printf("\u2588%02d\u001B[2D", i)

				time.Sleep(1 * time.Second)
			}
		}
	}
}
