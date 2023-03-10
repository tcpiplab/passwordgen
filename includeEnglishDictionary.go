package main

import (
	"compress/gzip"
	"io"
	"os"
	"strings"
)

// In this example, inputFile is the path to the dictionary file that you want to
// compress, and outputFile is the path to the compressed output file. The
// compressDictionary function opens the input file, creates the output file,
// creates a gzip writer that writes to the output file, and copies the contents
// of the input file to the gzip writer. Finally, it closes the gzip writer and
// returns any errors.
//
// You can then use the compressed dictionary file in your password generator by
// decompressing it at runtime when it needs to be used.
func compressDictionary(inputFile, outputFile string) error {
	// Open the input file
	inputFileHandle, err := os.Open(inputFile)
	if err != nil {
		return err
	}
	defer func(inputFileHandle *os.File) {
		err := inputFileHandle.Close()
		if err != nil {
			// TODO: Handle this error.
		}
	}(inputFileHandle)

	// Create the output file
	outputFileHandle, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer func(outputFileHandle *os.File) {
		err := outputFileHandle.Close()
		if err != nil {
			// TODO: Handle this error.
		}
	}(outputFileHandle)

	// Create a gzip writer that writes to the output file
	gzipWriter := gzip.NewWriter(outputFileHandle)
	defer func(gzipWriter *gzip.Writer) {
		err := gzipWriter.Close()
		if err != nil {
			// TODO: Handle this error.
		}
	}(gzipWriter)

	// Copy the contents of the input file to the gzip writer
	_, err = io.Copy(gzipWriter, inputFileHandle)
	if err != nil {
		return err
	}

	return nil
}

// In this example, dictionaryFile is the path to the compressed dictionary file.
// The getDictionaryWords function opens the compressed file, creates a gzip
// reader that reads from the compressed file, reads the contents of the gzip
// reader into a byte slice, converts the byte slice to a string, and splits the
// string into words. Finally, it returns the words as a slice of strings and any
// errors.
func getDictionaryWords(dictionaryFile string) ([]string, error) {
	// Open the compressed dictionary file
	compressedFile, err := os.Open(dictionaryFile)
	if err != nil {
		return nil, err
	}
	defer func(compressedFile *os.File) {
		err := compressedFile.Close()
		if err != nil {
			// TODO: Handle this error.
		}
	}(compressedFile)

	// Create a gzip reader that reads from the compressed file
	gzipReader, err := gzip.NewReader(compressedFile)
	if err != nil {
		return nil, err
	}
	defer func(gzipReader *gzip.Reader) {
		err := gzipReader.Close()
		if err != nil {
			// TODO: Handle this error.
		}
	}(gzipReader)

	// Read the contents of the gzip reader into a byte slice
	uncompressedData, err := io.ReadAll(gzipReader)
	if err != nil {
		return nil, err
	}

	// Convert the byte slice to a string and split it into words
	words := strings.Split(string(uncompressedData), "\n")

	return words, nil
}
