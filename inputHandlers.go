package main

import (
	"flag"
	"fmt"
)

func argsHandler() (
	interactive *bool,
	erase *bool,
	randomPasswords *bool,
	wordChains *bool,
	memorable2 *bool,
	help *bool,
	passPhrases *bool,
	memorable *bool,
	randomHex *bool,
	examples *bool,
	grammatical *bool,
	grammaticalAI *bool,
	grammaticalAIWithNumbers *bool,
	mnemonic *bool,
	memorable3 *bool,
) {
	// FIXME: Only -h works, not -help or --help
	help = flag.Bool(
		"help",
		false,
		"./passwordgen -h",
	)

	if *help {
		flag.Usage()
		return nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil
	}

	// Interactive mode is the default
	interactive = flag.Bool(
		"interactive",
		true,
		"./passwordgen --interactive[=false]")

	erase = flag.Bool(
		"erase",
		true,
		"./passwordgen --erase[=false]")

	randomPasswords = flag.Bool(
		"random",
		true,
		"./passwordgen --random")

	wordChains = flag.Bool(
		"word-chains",
		false,
		"./passwordgen --word-chains")

	passPhrases = flag.Bool(
		"passphrases",
		false,
		"./passwordgen --passphrases")

	memorable = flag.Bool(
		"memorable",
		false,
		"./passwordgen --memorable")

	memorable2 = flag.Bool(
		"memorable-2",
		false,
		"./passwordgen --memorable-2")

	memorable3 = flag.Bool(
		"memorable-3",
		false,
		"./passwordgen --memorable-3")

	randomHex = flag.Bool(
		"hex",
		false,
		"./passwordgen --hex")

	examples = flag.Bool(
		"examples",
		false,
		"./passwordgen --examples")

	grammatical = flag.Bool(
		"grammatical",
		false,
		"./passwordgen --grammatical")

	grammaticalAI = flag.Bool(
		"grammatical-ai",
		false,
		"./passwordgen --grammatical-ai  (Requires an openai.com GPT-4 API key)")

	grammaticalAIWithNumbers = flag.Bool(
		"grammatical-ai-with-numbers",
		false,
		"./passwordgen --grammatical-ai-with-numbers  (Requires an openai.com GPT-4 API key)")

	mnemonic = flag.Bool(
		"mnemonic",
		false,
		"./passwordgen --mnemonic  (Requires an openai.com GPT-4 API key)")

	flag.Parse()

	return interactive, erase, randomPasswords, wordChains, memorable2, nil, passPhrases, memorable, randomHex, examples, grammatical, grammaticalAI, grammaticalAIWithNumbers, mnemonic, memorable3
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
