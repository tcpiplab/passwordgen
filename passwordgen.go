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
	if len(os.Args) != 2 {
		color.Red("Please provide a password length as an argument")
		return
	}

	// Convert the requested length from string to int
	requested_password_length, err := strconv.Atoi(os.Args[1])

	if err != nil {
		color.Red("Invalid password length argument")
		return
	}

	rand.Seed(time.Now().UnixNano())

	var rows_columns [2]int
	rows_columns[0], rows_columns[1] = consoleSize()

	var rows int
	rows = rows_columns[0]

	for each_row := 0; each_row < rows-1; each_row++ {
		//fmt.Println(randString(requested_password_length))
		//color.Red(randString(requested_password_length))

		password := randString(requested_password_length)

		for i := 0; i < requested_password_length-1; i++ {
			character := password[i]

			//cyan_char := color.New(color.FgCyan, color.Bold)
			//red_char := color.New(color.FgRed, color.Bold)
			//green_char := color.New(color.FgGreen, color.Bold)
			//blue_char := color.New(color.FgBlue, color.Bold)

			//cyan_char.Printf("%s", character)
			//cyan_char.Printf("%s", string(character))

			// Assign a color to uppercase characters
			if character >= 65 && character <= 90 {
				color.Red(string(character))
				break
			}
			// Assign a color to lowercase characters
			if character >= 97 && character <= 122 {
				color.Blue(string(character))
				break
			}
			// Assign a color to number characters
			if character >= 48 && character <= 57 {
				color.Green(string(character))
				break
			}
			// Assign a color to special characters, first range
			if character >= 33 && character <= 47 {
				color.Magenta(string(character))
				break
			}
			// Assign a color to special characters, second range
			if character >= 58 && character <= 64 {
				color.Magenta(string(character))
				break
			}
			// Assign a color to special characters, third range
			if character >= 91 && character <= 96 {
				color.Magenta(string(character))
				break
			}
			// Assign a color to special characters, fourth range
			if character >= 123 && character <= 126 {
				color.Magenta(string(character))
				break
			} else {
				// Assign a color to any character not represented above
				color.Yellow(string(character))
			}

			//fmt.Printf("\n")
		}
		//fmt.Printf("\n")
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
	var allowed_characters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#%^&*()")
	b := make([]rune, length_of_rand_string)
	for i := range b {
		b[i] = allowed_characters[rand.Intn(len(allowed_characters))]
	}
	return string(b)
}
