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
	default:
		return runeSymbol
	}
}

func bionicBoldCount(wordLen int) int {
	switch {
	case wordLen <= 1:
		return 0
	case wordLen <= 3:
		return 1
	case wordLen == 4:
		return 2
	case wordLen <= 6:
		return 3
	case wordLen <= 9:
		return 4
	default:
		return int(float64(wordLen)*0.4 + 0.5)
	}
}

func parseWordToBold(text string, args BoldArgs) string {
	runes := []rune(text)
	wordLength := len(runes)

	var boldCount int
	if args.Percent == 0 {
		boldCount = bionicBoldCount(wordLength)
	} else {
		boldCount = int(float64(wordLength)*(float64(args.Percent)/100.0) + 0.5)
	}

	if boldCount == 0 {
		return text
	}

	boldPortion := string(runes[:boldCount])
	remainder := string(runes[boldCount:])

	switch args.TypeBold {
	case "html":
		return "<b>" + boldPortion + "</b>" + remainder
	case "ansi":
		return "\x1b[1m" + boldPortion + "\x1b[0m" + remainder
	default:
		return "**" + boldPortion + "**" + remainder
	}
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
		case runeSpace, runeSymbol:
			flushWord(string(r))
		}
	}
	return out.String()
}
