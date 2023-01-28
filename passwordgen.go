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
	flag.Parse()

	if *help {
		flag.Usage()
		return
	}

	if len(os.Args) != 2 {
		color.HiRed("\nPlease provide a password length as an argument\nOr -h for help.\n\n")
		return
	}

	// Convert the requested length from string to int
	requested_password_length, err := strconv.Atoi(os.Args[1])

	if int(requested_password_length) < 10 {

		color.HiRed("\nPassword length must be 10 or longer.\n\n")
		return
	}

	if err != nil {
		color.HiRed("Invalid password length argument")
		return
	}

	rand.Seed(time.Now().UnixNano())

	var rows_columns [2]int
	rows_columns[0], rows_columns[1] = consoleSize()

	var rows int
	rows = rows_columns[0]

	for each_row := 0; each_row < rows-1; each_row++ {

		password := randString(requested_password_length)

		for i := 0; i < requested_password_length-1; i++ {

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

	// https://gist.github.com/steinelu/aa9a5f402b584bc967eb216e054ceefb

	cmd := exec.Command("stty", "size")

	cmd.Stdin = os.Stdin

	out, err := cmd.Output()

	if err != nil {

		log.Fatal(err)
	}

	s := string(out)

	s = strings.TrimSpace(s)

	sArr := strings.Split(s, " ")

	height, err := strconv.Atoi(sArr[0])

	if err != nil {

		log.Fatal(err)
	}

	width, err := strconv.Atoi(sArr[1])

	if err != nil {

		log.Fatal(err)
	}

	return height, width
}

func randString(length_of_rand_string int) string {

	var allowed_characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#^&*()[]{}%")

	b := make([]rune, length_of_rand_string)

	for i := range b {

		b[i] = allowed_characters[rand.Intn(len(allowed_characters))]
	}

	return string(b)
}
