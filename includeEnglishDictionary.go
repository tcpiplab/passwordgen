package main

import (
	"compress/gzip"
	"io"
	"os"
	"strings"
)

func compressDictionary(inputFile, outputFile string) error {
	// Open the input file
	inputFileHandle, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer inputFileHandle.Close()

	// Create the output file
	outputFileHandle, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outputFileHandle.Close()

	// Create a gzip writer that writes to the output file
	gzipWriter := gzip.NewWriter(outputFileHandle)
	defer gzipWriter.Close()

	// Copy the contents of the input file to the gzip writer
	_, err = io.Copy(gzipWriter, inputFileHandle)
	if err != nil {
		return err
	}

	return nil
}

func getDictionaryWords(dictionaryFile string) ([]string, error) {
	// Open the compressed dictionary file
	compressedFile, err := os.Open(dictionaryFile)
	if err != nil {
		return nil, err
	}
	defer compressedFile.Close()

	// Create a gzip reader that reads from the compressed file
	gzipReader, err := gzip.NewReader(compressedFile)
	if err != nil {
		return nil, err
	}
	defer gzipReader.Close()

	// Read the contents of the gzip reader into a byte slice
	uncompressedData, err := io.ReadAll(gzipReader)
	if err != nil {
		return nil, err
	}

	// Convert the byte slice to a string and split it into words
	words := strings.Split(string(uncompressedData), "\n")

	return words, nil
}
