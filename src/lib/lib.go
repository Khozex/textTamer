package lib

import (
	"strings"
	"unicode"
)

type runeKind int

const (
	runeLetter runeKind = iota
	runeSpace
	runeSymbol
	runeOther
)

type BoldArgs struct {
	TypeBold string
	Percent  int
	Jumps    int
}

func classify(r rune) runeKind {
	switch {
	case unicode.IsLetter(r):
		return runeLetter
	case unicode.IsSpace(r):
		return runeSpace
	case unicode.IsDigit(r):
		return runeOther
	default:
		return runeSymbol
	}
}

func parseWordToBold(text string, args BoldArgs) string {
	runes := []rune(text)
	wordLength := len(runes)

	percent := args.Percent
	if percent < 50 {
		percent = 50
	}
	boldCount := int(float64(wordLength)*(float64(percent)/100.0) + 0.5)

	boldPortion := string(runes[:boldCount])
	remainder := string(runes[boldCount:])

	if args.TypeBold == "html" {
		return "<b>" + boldPortion + "</b>" + remainder
	}
	return "**" + boldPortion + "**" + remainder
}

func MakeTextToBold(text string, args BoldArgs) string {
	runes := []rune(text)
	var out strings.Builder
	out.Grow(len(text) + len(text)/4)

	var word strings.Builder
	jumpCount := args.Jumps

	flushWord := func(trailing string) {
		if word.Len() == 0 {
			out.WriteString(trailing)
			return
		}
		w := word.String()
		if jumpCount == args.Jumps {
			out.WriteString(parseWordToBold(w, args))
			jumpCount = 0
		} else {
			out.WriteString(w)
			jumpCount++
		}
		out.WriteString(trailing)
		word.Reset()
	}

	for i, r := range runes {
		isLast := i == len(runes)-1
		switch classify(r) {
		case runeLetter:
			word.WriteRune(r)
			if isLast {
				flushWord("")
			}
		case runeSpace:
			flushWord(string(r))
		case runeSymbol:
			flushWord(string(r))
		default:
			out.WriteRune(r)
		}
	}
	return out.String()
}
