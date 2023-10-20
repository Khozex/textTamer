package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func makeTextToBold(text string, args BoldArgs) string {
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
		} else if runeIdentify.typeChar == "space" {
			words = append(words, parseWordToBold(word, args))
			words = append(words, runeIdentify.char)
			word = ""
		} else {
			words = append(words, parseWordToBold(word, args)+runeIdentify.char)
			word = ""
		}
	}
	fmt.Println(strings.Join(words, ""))
	return strings.Join(words, "")
}

func main() {
	var stdinWord string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stdinWord += scanner.Text()
	}
	args := BoldArgs{
		typeBold: "md",
		porcent:  50,
	}
	makeTextToBold(stdinWord, args)
}
