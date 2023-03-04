package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

type Words struct {
	Word  string `json:"word"`
	Score int    `json:"score"`
}

func callWordApi() string {
	// Set up the HTTP client and request
	client := http.Client{}
	req, err := http.NewRequest("GET", "https://api.datamuse.com/sug", nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	searchString := randomWordStem()
	//searchString := "del"

	// Set any headers or query parameters as needed
	//req.Header.Add("Authorization", "Bearer <token>")
	q := req.URL.Query()
	q.Add("s", searchString)
	q.Add("max", "10")
	req.URL.RawQuery = q.Encode()

	// Send the request and read the response
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Parse the JSON response
	var words []Words
	err = json.NewDecoder(resp.Body).Decode(&words)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	longestWord := ""
	for _, word := range words {
		if len(word.Word) > len(longestWord) {
			longestWord = word.Word
		}
	}

	output := strings.ReplaceAll(longestWord, " ", "_")

	return output
}

func randomWordChain(requestedPasswordLength int) string {

	// Call callWordApi() and concatenate the returned words into a string
	var buffer bytes.Buffer

	// Choose a single delimiter to place between the words
	delimiters := "-_=+/\\|~^$#@&*:."
	delimiter := string(delimiters[rand.Intn(len(delimiters))])

	var word string

	for i := 0; i < requestedPasswordLength; i += len(word) {

		// TODO: Check if the word is obscure or not.
		// If it doesn't return anything from:
		// https://api.datamuse.com/words?v=enwiki&max=1&ml=ignominious
		// then reject it and grab another word.
		// TODO: If it returned anything used the word returned because it
		// will tend to be a more common word.
		word = callWordApi()

		word, _ = getBetterWord(word)

		if len(word) > 2 {

			buffer.WriteString(word)

			if i != requestedPasswordLength {
				// Add a delimiter between the words except for the last word
				if i != requestedPasswordLength-1 {

					buffer.WriteString(delimiter)
				}
			}
		}
	}

	// Replace spaces with an underscore character
	output := strings.ReplaceAll(buffer.String(), " ", "_")

	// Truncate the resulting word-chain password to the specified length
	// by removing characters from the right side
	if len(output) > requestedPasswordLength {

		output = strings.TrimSpace(output[:requestedPasswordLength])
	}

	// Colorize word-chain output
	colorizeCharacters(requestedPasswordLength, output)

	return output
}

func fileExists(filename string) (bool, error) {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func checkForWordList(rows int) {
	//OS := runtime.GOOS

	if OS == "darwin" || OS == "linux" || OS == "unix" {
		// Check if wordlist exists
		wordlistExists, err := fileExists("/usr/share/dict/words")
		if err != nil {
			// handle error
			fmt.Println("Error:", err)
			return
		}
		if wordlistExists {
			// file exists
			//fmt.Println("Wordlist file exists.")
			passwordRows := rows / 2

			// TODO: Grab the resulting array and call the api for words it triggers
			selectSeedWords(passwordRows)

		} else {
			// file does not exist
			fmt.Println("Wordlist file does not exist.")
		}
	} else if OS == "windows" {

		fmt.Println("Wordlist file does not exist.")
		fmt.Println("word-chains are not yet implemented for Windows.")
	}
}

// TODO: If I don't end up using this function then remove it.
func selectSeedWords(numPasswordRows int) []string {

	// open the file for reading
	file, err := os.Open("/usr/share/dict/words")

	// Handle any error from opening.
	if err != nil {
		panic(err)
	}

	// Wrap the file.Close() function in a closure to handle any error that might occur.
	defer func(file *os.File) {

		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	// read all the lines into memory
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// select numPasswordRows random words of at least three characters in length
	var arrSeedWords []string
	for i := 0; i < numPasswordRows; i++ {
		password := ""
		for len(password) < 3 || len(password) > 7 {
			password = lines[rand.Intn(len(lines))]
		}
		arrSeedWords = append(arrSeedWords, password)
	}

	//fmt.Printf("%s", arrSeedWords)

	return arrSeedWords
}

func getBetterWord(oldTimeyWord string) (string, error) {

	url := fmt.Sprintf("https://api.datamuse.com/words?v=enwiki&max=1&ml=%s", oldTimeyWord)
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			// TODO: Handle error
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var words []struct {
		Word string `json:"word"`
	}

	err = json.Unmarshal(body, &words)
	if err != nil {
		return "", err
	}

	if len(words) > 0 {

		// Return the more normal word found by the API
		return words[0].Word, nil
	}

	return "", fmt.Errorf("No word found for '%s'", oldTimeyWord)
}

// Here is how to call the above function:
//func main() {
//	ml := "ignominious"
//	word, err := getWord(ml)
//	if err != nil {
//		fmt.Printf("Error getting word for '%s': %s\n", ml, err)
//		return
//	}
//	fmt.Printf("Word for '%s': %s\n", ml, word)
//}
