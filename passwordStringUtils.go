package main

import (
	"math/rand"
	"strings"
	"time"
)

func randString(lengthOfRandString int) string {

	// Set allowed characters
	var allowedCharacters = []int32("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#^&*()[]{}%")

	// Make a list of type int32 of the length the user requested their passwords should be
	listOfInt32Characters := make([]int32, lengthOfRandString)

	for i := range listOfInt32Characters {

		// Grab random chars and put them in the list. But only from the set of allowed characters
		listOfInt32Characters[i] = allowedCharacters[rand.Intn(len(allowedCharacters))]
	}

	// Return a new random password string
	return string(listOfInt32Characters)
}

func randomWordStem() string {

	rand.Seed(time.Now().UnixNano())

	const vowels = "AEIOU"
	const consonants = "BCDFGHJKLMNPQRSTVWXYZ"
	var result string
	for {
		// Generate a random letter from the vowels or consonants string
		if rand.Intn(2) == 0 {
			result += string(vowels[rand.Intn(len(vowels))])
		} else {
			result += string(consonants[rand.Intn(len(consonants))])
		}
		// If the string contains at least one vowel and one consonant, return it
		if strings.ContainsAny(result, vowels) && strings.ContainsAny(result, consonants) {
			return result
		}
	}
}
