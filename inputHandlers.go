package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
)

func argsHandler() (
	interactive *bool,
	erase *bool,
	randomPasswords *bool,
	wordChains *bool,
	mixedPasswords *bool,
	help *bool,
	passPhrases *bool,
	memorable *bool,
	randomHex *bool,
	examples *bool,
	grammatical *bool,
	grammaticalAI *bool,
) {
	help = flag.Bool(
		"help",
		false,
		"./passwordgen n\nWhere n is the length of the password.\nLength must be the last argument.\n",
	)

	if *help {
		flag.Usage()
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil
	}

	// Interactive mode is the default
	interactive = flag.Bool(
		"interactive",
		true,
		"./passwordgen --interactive[=false]\n")

	erase = flag.Bool(
		"erase",
		true,
		"./passwordgen --erase[=false]\n")

	randomPasswords = flag.Bool(
		"random",
		true,
		"./passwordgen --random\n")

	wordChains = flag.Bool(
		"word-chains",
		false,
		"./passwordgen --word-chains\n")

	mixedPasswords = flag.Bool(
		"mixed",
		false,
		"./passwordgen --mixed\n")

	passPhrases = flag.Bool(
		"passphrases",
		false,
		"./passwordgen --passphrases\n")

	memorable = flag.Bool(
		"memorable",
		false,
		"./passwordgen --memorable\n")

	randomHex = flag.Bool(
		"hex",
		false,
		"./passwordgen --hex\n")

	examples = flag.Bool(
		"examples",
		false,
		"./passwordgen --examples\n")

	grammatical = flag.Bool(
		"grammatical",
		false,
		"./passwordgen --grammatical\n")

	grammaticalAI = flag.Bool(
		"grammatical-ai",
		false,
		"./passwordgen --grammatical-ai\n(Requires an openai.com GPT-4 API key)")

	flag.Parse()

	// TODO: For now the length is mandatory and must be the last arg
	if len(os.Args) < 2 {

		color.HiRed("\nPlease provide a password length as the last argument\nOr -h for help.\n")
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil
	}
	return interactive, erase, randomPasswords, wordChains, mixedPasswords, nil, passPhrases, memorable, randomHex, examples, grammatical, grammaticalAI
}

func ifInteractive(interactive *bool, rows int) bool {

	// TODO: handle `--interactive=false` arg by returning one password and exiting
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
