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

// randomWordChain() generates a random word-chain password of the specified length.
//
//	Parameters:
//	requestedPasswordLength - the length of the password to generate
//
//	Returns:
//	A string representing the generated password
func randomWordChain(requestedPasswordLength int) string {

	var buffer bytes.Buffer

	// Choose a single delimiter to place between the words
	delimiters := "-_=+/\\|~^$#@&*:."
	delimiter := string(delimiters[rand.Intn(len(delimiters))])

	var word string

	for i := 0; i < requestedPasswordLength; i += len(word) {

		// Grab a word from the compressed dictionary
		word = getWordFromCompressedDictionary(dictionaryData)

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
	colorizeCharactersWindows(requestedPasswordLength, output)

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

// checkForWordList() checks whether a word list file exists on the system.
//
//	Returns:
//	A boolean value indicating whether the word list file exists.
func checkForWordList() bool {

	if OS == "darwin" || OS == "linux" || OS == "unix" {

		// Check if wordlist exists
		wordlistExists, err := fileExists("/usr/share/dict/words")
		if err != nil {
			// handle error
			fmt.Println("Error:", err)
			return false
		}
		if wordlistExists {

			return true

		} else {
			// file does not exist
			fmt.Println("Wordlist file does not exist.")

			return false
		}
	} else if OS == "windows" {

		fmt.Println("wordlist file does not exist.")
		fmt.Println("word-chains and mixed passwords are not yet implemented for Windows.")

		return false
	}

	// We shouldn't ever get here
	return false
}

// selectSeedWords() opens a file containing a dictionary of words, reads all its lines into
// memory, and selects a given number of random words of at least three characters in length.
// It returns a slice of strings with the selected random words.
//
//	Arguments:
//	  - numPasswordRows: an integer with the number of random words to select.
//
//	Returns:
//	  - a slice of strings with numPasswordRows random words of at least three characters in length.
//
//	Example:
//	  seedWords := selectSeedWords(4)
//
//	fmt.Println(seedWords)
//	Output: [passage freedom border galaxy]
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

	return arrSeedWords
}

func getArrayFromCompressedDictionary(numPasswordRows int) []string {

	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// select numPasswordRows random words of at least three characters in length
	var arrSeedWords []string
	for i := 0; i < numPasswordRows; i++ {
		password := ""
		for len(password) < 3 || len(password) > 7 {
			//password = lines[rand.Intn(len(lines))]
			password = getWordFromCompressedDictionary(dictionaryData)
		}
		arrSeedWords = append(arrSeedWords, password)
	}

	return arrSeedWords
}

// getBetterWord() solves a problem that is caused by grabbing words from
// out of copyright dictionaries: most of the words are odd, obscure, or
// outdated.
func getBetterWord(oldTimeyWord string) (string, error) {

	// API call:
	//   - v: The vocabulary database to use. "enwiki" is the largest.
	//   - max: We only want one word.
	//   - ml: "means like". Specify the word we want near synonyms for.
	url := fmt.Sprintf("https://api.datamuse.com/words?v=enwiki&max=1&ml=%s", oldTimeyWord)

	httpResponse, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			// Don't know how we'd ever get here
			fmt.Printf("Error closing HTTP response body after reading: %s\n", err)
		}
	}(httpResponse.Body)

	httpResponseBodyJSON, err := io.ReadAll(httpResponse.Body)
	if err != nil {
		return "", err
	}

	// Typical JSON httpResponse from the API:
	//   [{"word":"disgraceful","score":30050217,"tags":["syn","adj","results_type:primary_rel"]}]
	// All we need is the word
	var words []struct {
		Word string `json:"word"`
	}

	err = json.Unmarshal(httpResponseBodyJSON, &words)
	if err != nil {
		return "", err
	}

	// Sometimes, if the oldTimeyWord is obscure enough, the API will respond with "[]".
	if len(words) > 0 {

		// Return the more normal word found by the API
		return words[0].Word, nil
	}

	return "", fmt.Errorf("No word found for '%s'", oldTimeyWord)
}
