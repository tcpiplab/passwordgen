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
	"fmt"
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
		fmt.Println("Please provide a password length as an argument")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Invalid password length argument")
		return
	}

	rand.Seed(time.Now().UnixNano())

	// rows, _ := strconv.Atoi(os.Getenv("LINES"))
	// fmt.Println("LINES == ", rows)

	var rows_columns [2]int
	rows_columns[0], rows_columns[1] = consoleSize()

	var rows int
	rows = rows_columns[0]

	for i := 0; i < rows-1; i++ {
		fmt.Println(randString(n))
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

	heigth, err := strconv.Atoi(sArr[0])
	if err != nil {
		log.Fatal(err)
	}

	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}
	return heigth, width
}

func randString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#%^&*()")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
