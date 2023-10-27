package lib

import (
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
	Percent  int
	Jumps    int
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

	if boldArgs.Percent < 50 {
		boldArgs.Percent = 50
	}

	halfWordLength := int(float64(wordLength)*(float64(boldArgs.Percent)/100.0) + 0.5)

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
	var runeCount int
	jumpCount := args.Jumps
	for _, runeChar := range runes {
		runeCount++
		runeIdentify := identifyRuneByType(runeChar)
		if runeIdentify.typeChar == "string" {
			word += runeIdentify.char
			if runeCount == len(runes) {
				if jumpCount == args.Jumps {
					words = append(words, parseWordToBold(word, args))
					jumpCount = 0
				} else {
					words = append(words, word)
					jumpCount++
				}
			}
		} else if runeIdentify.typeChar == "space" && word != "" {
			if jumpCount == args.Jumps {
				words = append(words, parseWordToBold(word, args))
				jumpCount = 0
			} else {
				words = append(words, word)
				jumpCount++
			}
			words = append(words, runeIdentify.char)
			word = ""
		} else if runeIdentify.typeChar == "symbol" && word != "" {
			if jumpCount == args.Jumps {
				words = append(words, parseWordToBold(word, args)+runeIdentify.char)
				jumpCount = 0
			} else {
				words = append(words, word+runeIdentify.char)
				jumpCount++
			}
			word = ""
		} else {
			words = append(words, runeIdentify.char)
		}
	}
	return strings.Join(words, "")
}
