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

	//// Execute the shell command `stty size` which returns two integers:
	//// height and width of the user's heightAndWidthString terminal
	//modeCommand := exec.Command("mode")
	//
	//// Specify the shell's heightAndWidthString STDIN so that executing
	//// I don't know if this is necessary for the Windows "mode" command
	//modeCommand.Stdin = os.Stdin
	//
	//// Execute the `mode` command and save the output and any resulting error.
	//heightAndWidthBytes, err := modeCommand.Output()
	//
	//// If it errored heightAndWidthBytes, log to the screen and exit
	//if err != nil {
	//
	//	log.Fatal("Error trying to get the size of the terminal:", err)
	//}
	//
	//// Save the height and width values as a string
	//heightAndWidthString := string(heightAndWidthBytes)
	//
	////fmt.Printf(heightAndWidthString)
	//
	//// Split the string into multiple lines based on the newline character
	//arrStringsFromModeCommand := strings.Split(heightAndWidthString, "\n")
	//
	//// Get the third and fourth rows from the array
	//rowsLineString := arrStringsFromModeCommand[3]
	//colsLineString := arrStringsFromModeCommand[4]
	//
	//fmt.Printf("%s\n", rowsLineString)
	//fmt.Printf("%s\n", colsLineString)
	//
	//// Remove extra whitespace
	//heightAndWidthString = strings.TrimSpace(heightAndWidthString)
	//
	//// Split height and width into an array of two strings
	//heightAndWidthArray := strings.Split(heightAndWidthString, " ")
	//
	//// Convert height to an integer
	//// Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
	//height, err := strconv.Atoi(heightAndWidthArray[0])
	//
	//// If the conversion to int errored out, log to the screen and exit
	//if err != nil {
	//
	//	log.Fatal("Error trying to convert terminal height to an integer:", err)
	//}
	//
	//// Convert width to an integer
	//width, err := strconv.Atoi(heightAndWidthArray[1])
	//
	//// If the conversion to int errored out, log to the screen and exit
	//if err != nil {
	//
	//	log.Fatal("Error trying to convert terminal width to an integer:", err)
	//}

	// Set to Windows default values for now
	return 30, 120
}

//func consoleSizeWindows() (int, int) {
//
//	cmd := exec.Command("mode")
//	stdout, err := cmd.StdoutPipe()
//
//	if err != nil {
//
//		fmt.Printf("Error creating StdoutPipe: %s", err.Error())
//		// If error, return default size of a windows cmd window
//
//		return 25, 80
//	}
//	//if err := cmd.Start(); err != nil {
//	//
//	//	fmt.Printf("Error starting command: %s", err.Error())
//	//
//	//	// If error, return default size of a windows cmd window
//	//	return 25, 80
//	//}
//
//	err = cmd.Start()
//	if err != nil {
//		log.Fatal(err)
//	}
//	log.Printf("Waiting for command to finish...")
//	err = cmd.Wait()
//	log.Printf("Command finished with error: %v", err)
//
//	//defer func(cmd *exec.Cmd) {
//	//	err := cmd.Wait()
//	//	if err != nil {
//	//
//	//		fmt.Printf("Problem deferring cmd.Wait(): %s", err)
//	//	}
//	//}(cmd)
//
//	scanner := bufio.NewScanner(stdout)
//	var width, height int
//
//	for scanner.Scan() {
//
//		fmt.Printf("About to read line.")
//		line := scanner.Text()
//		fmt.Printf("Done reading line.")
//
//		fmt.Printf("line is: %s", line)
//
//		//os.Exit(0)
//
//		//input := "Status for device CON:----------------------    Lines:          3000    Columns:        91    Keyboard rate:  31    Keyboard delay: 1    Code page:      437"
//		pairs := make(map[string]string)
//
//		// Split the input string into substrings based on whitespace
//		fields := strings.Fields(line)
//
//		// Loop over the substrings and split each one into a key-value pair
//		for i := 1; i < len(fields); i += 2 {
//			key := strings.Trim(fields[i], ":")
//			value := fields[i+1]
//			pairs[key] = value
//		}
//
//		// Print the key-value pairs
//		for key, value := range pairs {
//			fmt.Printf("%s: %s\n", key, value)
//		}
//	}
//
//	fmt.Printf("Terminal size: %d rows x %d columns\n", height, width)
//
//	if err != nil {
//
//		// Handle the error
//		fmt.Println("Error getting terminal size:", err)
//
//		// If error, return default size of a windows cmd window
//		return 25, 80
//
//	} else {
//
//		return height, width
//	}
//
//}
