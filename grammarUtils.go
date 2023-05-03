package main

import (
	"math/rand"
	"strings"
	"time"
)

func getVocabWords() (string, string, string, string, string, string, string, string, string) {
	verb := getEnglishVocabWord("verb")
	//verb2 := getEnglishVocabWord("verb")
	noun := getEnglishVocabWord("noun")
	adverb := getEnglishVocabWord("adverb")
	adjective := getEnglishVocabWord("adjective")
	article := getRandomArticle()
	auxVerb := getRandomAuxVerb()
	pronounAndVerbPresent := getPronounAndVerbPresent()
	possessivePronoun := getPossessivePronoun()
	preposition := getPreposition()
	return verb, noun, adverb, adjective, article, auxVerb, pronounAndVerbPresent, possessivePronoun, preposition
}

func getPreposition() string {
	randomInt := rand.Intn(10) // Generate a random integer between 0 and 9

	var preposition string

	switch randomInt {
	case 0:
		preposition = "in"
	case 1:
		preposition = "on"
	case 2:
		preposition = "at"
	case 3:
		preposition = "over"
	case 4:
		preposition = "under"
	case 5:
		preposition = "between"
	case 6:
		preposition = "behind"
	case 7:
		preposition = "before"
	case 8:
		preposition = "after"
	case 9:
		preposition = "through"
	}
	return preposition
}

func getRandomPronoun() string {
	pronouns := []string{"he", "she", "they", "it", "I", "you", "we"}
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(pronouns))
	return pronouns[randomIndex]
}

func getPossessivePronoun() string {

	randomInt := rand.Intn(10) // Generate a random integer between 0 and 9

	var possessivePronoun string

	switch randomInt {
	case 0:
		possessivePronoun = "mine"
	case 1:
		possessivePronoun = "yours"
	case 2:
		possessivePronoun = "his"
	case 3:
		possessivePronoun = "hers"
	case 4:
		possessivePronoun = "its"
	case 5:
		possessivePronoun = "ours"
	case 6:
		possessivePronoun = "theirs"
	case 7:
		possessivePronoun = "someone's"
	case 8:
		possessivePronoun = "nobody's"
	case 9:
		possessivePronoun = "the dog's"
	}
	return possessivePronoun
}

func getPronounAndVerbPresent() string {
	// Choose a pronounAndVerbPresent at random
	randomInt := rand.Intn(10) // Generate a random integer between 0 and 9

	var pronounAndVerbPresent string

	switch randomInt {
	case 0:
		pronounAndVerbPresent = "it is"
	case 1:
		pronounAndVerbPresent = "that is"
	case 2:
		pronounAndVerbPresent = "those are"
	case 3:
		pronounAndVerbPresent = "this is"
	case 4:
		pronounAndVerbPresent = "he is"
	case 5:
		pronounAndVerbPresent = "she is"
	case 6:
		pronounAndVerbPresent = "they are"
	case 7:
		pronounAndVerbPresent = "we are"
	case 8:
		pronounAndVerbPresent = "you are"
	case 9:
		pronounAndVerbPresent = "I am"
	}
	return pronounAndVerbPresent
}

func getVerbModifier(r *rand.Rand) string {
	// Generate a random number between 0 and 4 (inclusive).
	randomNumber := r.Intn(5)

	// Randomly choose a verb modifier using a switch statement.
	var verbModifier string
	switch randomNumber {
	case 0:
		verbModifier = "never"
	case 1:
		verbModifier = "always"
	case 2:
		verbModifier = "rarely"
	case 3:
		verbModifier = "sometimes"
	case 4:
		verbModifier = "often"
	default:
		verbModifier = "unknown"
	}
	return verbModifier
}

func getRandomAuxVerb() string {
	randomAuxVerbIndex := rand.Intn(15)

	var auxVerb string

	switch randomAuxVerbIndex {
	case 0:
		auxVerb = "wasn't"
	case 1:
		auxVerb = "is"
	case 2:
		auxVerb = "isn't"
	case 3:
		auxVerb = "was"
	case 4:
		auxVerb = "were"
	case 5:
		auxVerb = "will"
	case 6:
		auxVerb = "shall"
	case 7:
		auxVerb = "shall not"
	case 8:
		auxVerb = "won't" // contraction of "will not"
	case 9:
		auxVerb = "hasn't" // contraction of "has not"
	case 10:
		auxVerb = "didn't"
	case 11:
		auxVerb = "can't"
	case 12:
		auxVerb = "wouldn't"
	case 13:
		auxVerb = "shouldn't"
	case 14:
		auxVerb = "won't"
	}
	return auxVerb
}

func getRandomArticle() string {

	var article string

	randomIndex := rand.Intn(10)

	switch randomIndex {
	case 0:
		article = "a"
	case 1:
		article = "the"
	case 2:
		article = "one"
	case 3:
		article = "my"
	case 4:
		article = "your"
	case 5:
		article = "his"
	case 6:
		article = "her"
	case 7:
		article = "their"
	case 8:
		article = "someone's"
	case 9:
		article = "any"
	}
	return article
}

// modifyArticle checks if the firstLetter variable is present in the vowels string.
// If it is and the article is "a", the function returns "an".
// In all other cases, the function returns the article unchanged.
func modifyArticle(followingWord, article string) string {

	firstLetter := followingWord[:1]

	// TODO: this function isn't working. Try print statements.

	if isVowel(firstLetter) && article == "a" {

		//fmt.Printf("----------\n%s, %s\n", article, followingWord)
		//fmt.Printf("firstLetter: %s\n", firstLetter)
		//fmt.Printf("article: %s\n----------\n", article)

		return "an"

	} else {

		return article
	}

	//return article
}

func isVowel(char string) bool {
	vowels := "aeiouAEIOU"

	if len(char) != 1 {
		return false
	}

	// Return true if the char is a vowel
	return strings.Contains(vowels, char)
}

func convertVerbToPastTense(verb string) string {
	// If the verb ends with 'e', just add 'd' to the end.
	if strings.HasSuffix(verb, "e") {
		return verb + "d"
	}

	// If the verb ends with a consonant followed by 'y', replace 'y' with 'ied'.
	if len(verb) >= 2 && strings.Contains("bcdfghjklmnpqrstvwxyz", string(verb[len(verb)-2])) && strings.HasSuffix(verb, "y") {
		return verb[:len(verb)-1] + "ied"
	}

	// For other verbs, just add 'ed' to the end.
	return verb + "ed"
}

func applyAuxiliaryVerb(auxVerb string, verbPresentTense string) string {
	auxVerb = strings.ToLower(auxVerb)
	verbPresentTense = strings.ToLower(verbPresentTense)

	switch auxVerb {
	case "had", "has", "was", "is", "were", "hadn't", "weren't", "hasn't", "wasn't", "isn't":
		return convertVerbToPastTense(verbPresentTense)
	default:
		return verbPresentTense
	}
}

func convertIrregularVerb(auxVerb string, verb string) string {
	switch strings.ToLower(verb) {
	case "be":
		verb = "was"
	case "begin":
		verb = "began"
	case "bite":
		verb = "bit"
	case "blow":
		verb = "blew"
	case "break":
		verb = "broke"
	case "bring":
		verb = "brought"
	case "build":
		verb = "built"
	case "buy":
		verb = "bought"
	case "catch":
		verb = "caught"
	case "choose":
		verb = "chose"
	case "come":
		verb = "came"
	case "cost":
		verb = "cost"
	case "cut":
		verb = "cut"
	case "do":
		verb = "did"
	case "draw":
		verb = "drew"
	case "drink":
		verb = "drank"
	case "drive":
		verb = "drove"
	case "eat":
		verb = "ate"
	case "fall":
		verb = "fell"
	case "feel":
		verb = "felt"
	case "fight":
		verb = "fought"
	case "find":
		verb = "found"
	case "fly":
		verb = "flew"
	case "forget":
		verb = "forgot"
	case "freeze":
		verb = "froze"
	case "get":
		verb = "got"
	case "give":
		verb = "gave"
	case "go":
		verb = "went"
	case "grow":
		verb = "grew"
	case "hang":
		verb = "hung"
	case "have":
		verb = "had"
	case "hear":
		verb = "heard"
	case "hide":
		verb = "hid"
	case "hit":
		verb = "hit"
	case "hold":
		verb = "held"
	case "hurt":
		verb = "hurt"
	case "keep":
		verb = "kept"
	case "know":
		verb = "knew"
	case "lead":
		verb = "led"
	case "leave":
		verb = "left"
	case "lend":
		verb = "lent"
	case "let":
		verb = "let"
	case "lie":
		verb = "-"
	case "light":
		verb = "lit"
	case "lose":
		verb = "lost"
	case "make":
		verb = "made"
	case "mean":
		verb = "meant"
	case "meet":
		verb = "met"
	case "pay":
		verb = "paid"
	case "put":
		verb = "put"
	case "read":
		verb = "read"
	case "ride":
		verb = "rode"
	case "ring":
		verb = "rang"
	case "rise":
		verb = "rose"
	case "run":
		verb = "ran"
	case "say":
		verb = "said"
	case "see":
		verb = "saw"
	case "sell":
		verb = "sold"
	case "send":
		verb = "sent"
	case "set":
		verb = "set"
	case "shake":
		verb = "shook"
	case "shine":
		verb = "shone"
	case "shoot":
		verb = "shot"
	case "show":
		verb = "showed"
	case "shut":
		verb = "shut"
	case "sing":
		verb = "sang"
	case "sink":
		verb = "sank"
	case "sit":
		verb = "sat"
	case "sleep":
		verb = "slept"
	case "slide":
		verb = "slid"
	case "speak":
		verb = "spoke"
	case "spend":
		verb = "spent"
	case "spin":
		verb = "spun"
	case "spread":
		verb = "spread"
	case "stand":
		verb = "stood"
	case "steal":
		verb = "stole"
	case "stick":
		verb = "stuck"
	case "sting":
		verb = "stung"
	case "strike":
		verb = "struck"
	case "swear":
		verb = "swore"
	case "sweep":
		verb = "swept"
	case "swim":
		verb = "swam"
	case "take":
		verb = "took"
	case "teach":
		verb = "taught"
	case "tear":
		verb = "tore"
	case "tell":
		verb = "told"
	case "think":
		verb = "thought"
	case "throw":
		verb = "threw"
	case "understand":
		verb = "understood"
	case "wake":
		verb = "woke"
	case "wear":
		verb = "wore"
	case "win":
		verb = "won"
	case "write":
		verb = "wrote"
	case "grind":
		verb = "ground"
	case "stop":
		verb = "stopped"
	default:
		// If not an irregular verb, do the standard conversion to past tense
		// if auxiliary verb requires it
		return applyAuxiliaryVerb(auxVerb, verb)
	}
	// return past tense version of irregular verb
	return verb
}

func getConjunctiveAdverbialPhrase() string {

	var conjunctiveAdverbialPhrase string

	phrases := []string{
		"And then,",
		"In addition,",
		"Therefore,",
		"However,",
		"Conversely,",
		"Meanwhile,",
		"Moreover,",
		"Nonetheless,",
		"Furthermore,",
		"On the other hand,",
		"For example,",
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(phrases))

	switch randomIndex {
	case 0:
		conjunctiveAdverbialPhrase = phrases[0]
	case 1:
		conjunctiveAdverbialPhrase = phrases[1]
	case 2:
		conjunctiveAdverbialPhrase = phrases[2]
	case 3:
		conjunctiveAdverbialPhrase = phrases[3]
	case 4:
		conjunctiveAdverbialPhrase = phrases[4]
	case 5:
		conjunctiveAdverbialPhrase = phrases[5]
	case 6:
		conjunctiveAdverbialPhrase = phrases[6]
	case 7:
		conjunctiveAdverbialPhrase = phrases[7]
	case 8:
		conjunctiveAdverbialPhrase = phrases[8]
	case 9:
		conjunctiveAdverbialPhrase = phrases[9]
	case 10:
		conjunctiveAdverbialPhrase = phrases[10]
	default:
		conjunctiveAdverbialPhrase = "Wait. Um,"
	}
	return conjunctiveAdverbialPhrase
}

// The modifySentence function generates a random float between 0 and 1. If the
// value is less than 0.5, it calls getConjunctiveAdverbialPhrase() and prepends
// the resulting phrase to the input sentence. Otherwise, it returns the original
// sentence.
func maybePrependConjAdvPhrase(sentence string) string {
	rand.Seed(time.Now().UnixNano())
	shouldModify := rand.Float64()

	if shouldModify < 0.5 {
		return getConjunctiveAdverbialPhrase() + " " + sentence
	}
	return sentence
}
