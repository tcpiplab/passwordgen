package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"time"
)

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
