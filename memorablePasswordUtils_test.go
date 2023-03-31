package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
	"unicode"
)

func TestRandomYear(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	minYear := 0
	maxYear := 2000

	// Run the test multiple times to check for randomness
	for i := 0; i < 100; i++ {
		randomYearStr := RandomYear()
		randomYear, err := strconv.Atoi(randomYearStr)

		if err != nil {
			t.Errorf("RandomYear() returned an invalid number: %s", randomYearStr)
		}

		if randomYear < minYear || randomYear > maxYear {
			t.Errorf("RandomYear() returned a number out of range: %d (expected between %d and %d)", randomYear, minYear, maxYear)
		}
	}
}

func TestMemorableTransformFive(t *testing.T) {
	tests := []struct {
		name                    string
		requestedPasswordLength int
		expectRandomUnit        bool
	}{
		{"length 24", 24, false},
		{"length 25", 25, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			memorablePassword := memorableTransformFive("", tt.requestedPasswordLength)

			fmt.Println(memorablePassword)

			// Test if password matches the pattern: {Degrading-1729-Earshot}
			// https://regex101.com/r/vsmcXd/2
			assert.Regexp(t, "^[\\[\\{\\(\\<][A-Za-z]+-\\d{1,4}-[A-Za-z]+[\\]\\}\\)\\>]$", memorablePassword)

			splitPassword := strings.Split(memorablePassword, "-")

			// Test for word, dash, word
			if len(splitPassword) != 3 {
				t.Errorf("memorableTransformFive() returned password with incorrect format: got %s", memorablePassword)
			}

			// Create a different slice to test this so we don't modify the original password
			splitPasswordStripBracketsLeft := splitPassword

			// This for loop iterates through each element of the splitPassword slice, applying
			// the strings.TrimLeftFunc function to remove non-alphabetic characters from the
			// beginning of each string, and updates the slice with the modified strings.
			for i, passwordPart := range splitPassword {
				splitPasswordStripBracketsLeft[i] = strings.TrimLeftFunc(passwordPart, func(r rune) bool {
					return !unicode.IsLetter(r)
				})
			}

			if !isCapitalized(splitPasswordStripBracketsLeft[0]) {
				t.Errorf("memorableTransformFive() failed to capitalize the first word: got %s", splitPassword[0])
			}

			if !isCapitalized(splitPasswordStripBracketsLeft[2]) {
				t.Errorf("memorableTransformFive() failed to capitalize the second word: got %s", splitPassword[2])
			}

		})
	}
}

func isCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	firstRune := rune(s[0])
	return unicode.IsUpper(firstRune) && s[1:] == strings.ToLower(s[1:])
}

func TestCapitalizeFirstLetter(t *testing.T) {
	tests := []struct {
		input  string
		output string
	}{
		{"hello", "Hello"},
		{"world", "World"},
		{"capitalize", "Capitalize"},
		{"test123", "Test123"},
		{"", ""},
	}

	for _, test := range tests {
		capitalized := capitalizeFirstLetter(test.input)
		if capitalized != test.output {
			t.Errorf("capitalizeFirstLetter() failed for input %s: expected %s, got %s", test.input, test.output, capitalized)
		}
	}
}

func TestCapitalizeFirstLetterIdempotent(t *testing.T) {
	tests := []string{
		"Hello",
		"World",
		"Capitalize",
		"Test123",
		"",
	}

	for _, input := range tests {
		capitalized := capitalizeFirstLetter(input)
		if capitalized != input {
			t.Errorf("capitalizeFirstLetter() is not idempotent for input %s: expected %s, got %s", input, input, capitalized)
		}
	}
}

func TestCapitalizeFirstLetterOnlyFirstRune(t *testing.T) {
	input := "gödel"
	expected := "Gödel"
	capitalized := capitalizeFirstLetter(input)

	if capitalized != expected {
		t.Errorf("capitalizeFirstLetter() failed for input %s: expected %s, got %s", input, expected, capitalized)
	}
}

func TestAppendRandomUnit(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	units := []string{
		"Mhz", "Ghz", "Mbps", "Mph", "Gbps", "Kbps", "inches", "feet", "miles",
		"Hz", "kHz", "THz", "nm", "mm", "cm", "m", "km", "yd", "mi", "nmi",
		"liters", "gallons", "pints", "quarts", "milliliters", "cubic meters",
		"grams", "kilograms", "pounds", "ounces", "tons", "tonnes",
		"seconds", "minutes", "hours", "days", "weeks", "months", "years",
		"degrees Celsius", "degrees Fahrenheit", "Kelvin",
		"pascals", "bars", "atmospheres", "mmHg", "torr",
		"volts", "amperes", "watts", "ohms", "farads", "henrys", "teslas", "webers",
		"Joules", "calories", "BTU", "ergs", "electronvolts",
		"lumens", "candelas", "lux", "foot-candles",
		"bits", "bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB",
		"dpi", "ppi", "pt", "em", "rem", "px",
	}

	testNumbers := []string{"42", "123", "0", "1000"}

	// Run the test multiple times to check for randomness
	for i := 0; i < 100; i++ {
		for _, testNumber := range testNumbers {
			result := appendRandomUnit(testNumber)

			if !strings.HasPrefix(result, testNumber) {
				t.Errorf("appendRandomUnit() result does not start with the input number: got %s, expected prefix %s", result, testNumber)
			}

			appendedUnit := strings.TrimPrefix(result, testNumber)
			unitFound := false

			for _, unit := range units {
				if unit == appendedUnit {
					unitFound = true
					break
				}
			}

			if !unitFound {
				t.Errorf("appendRandomUnit() appended an invalid unit: got %s", appendedUnit)
			}
		}
	}
}
