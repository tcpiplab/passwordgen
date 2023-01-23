package main

// My first Golang code. Generates a screen full of random char passwords of a 
// specified. ChatGPT wrote the stub, I grabbed a gist for the column size 
// stuff, but had to tweak both to get things working.

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

