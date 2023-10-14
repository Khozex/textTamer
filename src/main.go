package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func breakTextIntoWords(text string) []string {
	var words []string
	var word string
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

func readToFile(path string) string {
	absPath, _ := filepath.Abs(path)
	file, err := os.Open(absPath)
	if err != nil {
		fmt.Println("Error reading file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return strings.Join(text, " ")
}

func parseWordToBoldMd(text string) string {
	wordLength := len(text)
	halfWordLength := wordLength / 2
	wordBold := "**" + text[:halfWordLength] + "**" + text[halfWordLength:]
	return wordBold
}

func parseWordToBoldHtml(text string) string {
	wordLength := len(text)
	halfWordLength := wordLength / 2
	wordBold := "<b>" + text[:halfWordLength] + "</b>" + text[halfWordLength:]
	return wordBold
}

func main() {
	var filePathFlag = flag.String("f", "", "path to file")
	var formatFlag = flag.String("t", "md", "format of output text")
	flag.Parse()
	var parser = parseWordToBoldMd
	if *formatFlag == "html" {
		parser = parseWordToBoldHtml
	}
	if *filePathFlag != "" {
		text := readToFile(*filePathFlag)
		words := breakTextIntoWords(text)
		var bionicWords []string
		for _, word := range words {
			bionicWords = append(bionicWords, parser(word))
		}
		fmt.Print(bionicWords)
		return
	}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	words := breakTextIntoWords(text)
	var bionicWords []string
	for _, word := range words {
		bionicWords = append(bionicWords, parser(word))
	}
	fmt.Print(bionicWords)
}
