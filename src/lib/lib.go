package lib

import (
	"fmt"
	"strings"
	"unicode"
)

type Caracter struct {
	char         string
	typeChar     string
	identiFyRune int
}

type BoldArgs struct {
	TypeBold string
	Porcent  int
}

func categorizeRune(r rune) string {
	switch {
	case unicode.IsLetter(r):
		return "string"
	case unicode.IsDigit(r):
		return "number"
	case unicode.IsSpace(r):
		return "space"
	default:
		return "symbol"
	}
}

func parseWordToBold(text string, boldArgs BoldArgs) string {
	runes := []rune(text)
	wordLength := len(runes)

	if boldArgs.Porcent < 50 {
		boldArgs.Porcent = 50
	}

	halfWordLength := int(float64(wordLength)*(float64(boldArgs.Porcent)/100.0) + 0.5)

	boldPortion := string(runes[:halfWordLength])
	remainder := string(runes[halfWordLength:])

	if boldArgs.TypeBold == "html" {
		return "<b>" + boldPortion + "</b>" + remainder
	}
	return "**" + boldPortion + "**" + remainder
}

func identifyRuneByType(runeChar rune) Caracter {
	var runeWithIdentify Caracter
	runeWithIdentify.char = string(rune(runeChar))
	runeWithIdentify.identiFyRune = int(runeChar)
	runeWithIdentify.typeChar = categorizeRune(runeChar)
	return runeWithIdentify
}

func MakeTextToBold(text string, args BoldArgs) string {
	runes := []rune(text)
	var words []string
	var word string
	var count int
	for _, runeChar := range runes {
		count++
		runeIdentify := identifyRuneByType(runeChar)
		if runeIdentify.typeChar == "string" {
			word += runeIdentify.char
			if count == len(runes) {
				words = append(words, parseWordToBold(word, args))
			}
		} else if runeIdentify.typeChar == "space" && word != "" {
			words = append(words, parseWordToBold(word, args))
			words = append(words, runeIdentify.char)
			word = ""
		} else if runeIdentify.typeChar == "symbol" && word != "" {
			words = append(words, parseWordToBold(word, args)+runeIdentify.char)
			word = ""
		} else {
			words = append(words, runeIdentify.char)
		}
	}
	fmt.Println(strings.Join(words, ""))
	return strings.Join(words, "")
}
