package main

import "github.com/fatih/color"

// checkPasswordLength checks the length of a password and returns a boolean indicating
// whether the password length is valid or not. If the requested password length is less
// than 10 characters, it will print a red-colored error message and return true. If an
// error is passed in as the second argument, it will print a red-colored error message
// indicating that the password length argument is invalid and return true. Otherwise,
// it will return false to indicate that the password length is valid.
//
// Parameters:
// - requestedPasswordLength: the length of the password to be checked
// - err: an error indicating if the password length argument is invalid
//
// Returns:
// - a boolean value indicating if the password length is valid or not
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
