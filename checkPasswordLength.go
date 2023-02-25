package main

import "github.com/fatih/color"

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
