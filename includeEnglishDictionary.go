package main

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"math/rand"
	"strings"
	"time"
)

// getWordFromCompressedDictionary() decodes, decompresses, and selects a random word
// from a compressed dictionary data string that contains words separated by newline
// characters and tab-separated indices and words. It returns a string with the selected
// random word. If there is an error decoding, decompressing, or selecting a random word,
// the function panics with the corresponding error message.
//
//	Arguments:
//	  - dictionaryData: a string with base64-encoded and gzip-compressed data of the dictionary.
//
//	Returns:
//	  - a string with the selected random word from the dictionary, without leading or trailing spaces.
func getWordFromCompressedDictionary(dictionaryData string) string {

	// Decode the base64 data.
	decoded, err := base64.StdEncoding.DecodeString(dictionaryData)
	if err != nil {
		panic(err)
	}

	// Decompress the data.
	readerGzip, err := gzip.NewReader(bytes.NewReader(decoded))
	if err != nil {
		panic(err)
	}
	defer func(readerGzip *gzip.Reader) {
		err := readerGzip.Close()
		if err != nil {
			fmt.Printf("Error closing the gzip reader")
		}
	}(readerGzip)

	uncompressed, err := io.ReadAll(readerGzip)
	if err != nil {
		panic(err)
	}

	// Split the uncompressed data into lines.
	lines := strings.Split(string(uncompressed), "\n")

	// Seed the random number generator.
	rand.Seed(time.Now().UnixNano())

	// Select a random word.
	index := rand.Intn(len(lines) - 1)

	randomWordIndexed := lines[index]

	randomWordAlone := strings.Split(randomWordIndexed, "\t")[1]

	output := strings.TrimSpace(randomWordAlone)

	return output
}

// getArrayFromCompressedDictionary() generates an array of random strings of at
// least three characters in length by selecting random words from a compressed
// dictionary.
//
//	Parameters:
//	  - numPasswordRows: The number of random strings to generate
//
//	Returns:
//	  - An array of random strings
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
