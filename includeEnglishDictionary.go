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

	//// Use the uncompressed dictionary data as needed.
	//dictionaryString := string(uncompressed)
	//// ...
	//
	//dictionaryString, err := ioutil.ReadAll(readerGzip)
	//if err != nil {
	//	panic(err)
	//}

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
