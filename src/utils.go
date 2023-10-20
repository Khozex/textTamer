package main

import (
	"unicode"
)

type Caracter struct {
	char         string
	typeChar     string
	identiFyRune int
}

type BoldArgs struct {
	typeBold string
	porcent  int
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
	halfWordLength := wordLength / 2

	if halfWordLength == 0 {
		return text
	}

	boldPortion := string(runes[:halfWordLength])
	remainder := string(runes[halfWordLength:])

	if boldArgs.typeBold == "html" {
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
