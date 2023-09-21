package main

import (
	"bufio"
	"fmt"
	"os"
)

func breakTextIntoWords(text string) []string {
	var words []string
	var word string
	//split using regex
	for _, char := range text {
		if string(char) == " " {
			words = append(words, word)
			word = ""
		} else {
			word += string(char)
		}
	}
	words = append(words, word)
	return words
}

func parseWordToBold(text string) string {
	wordLength := len(text)
	halfWordLength := wordLength / 2
	wordBold := "**" + text[:halfWordLength] + "**" + text[halfWordLength:]
	return wordBold
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	words := breakTextIntoWords(text)
	var bionicWords []string
	for _, word := range words {
		bionicWords = append(bionicWords, parseWordToBold(word))
	}
	fmt.Print(bionicWords)
}
