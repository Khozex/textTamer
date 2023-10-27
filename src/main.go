package main

import (
	"bufio"
	"flag"
	"os"
	"textTamer/lib"
)

func setFlags() {
	flag.Int("f", 50, "Percent of word to bold")
	flag.String("t", "md", "Type of bold")
	flag.Int("j", 0, "Jumps of word to bold")
	flag.String("o", "", "Output file")
	flag.Parse()
}

func saveInFile(name string, content string) {
	file, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString(content)
}

func readInFile(name string) string {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var content string
	for scanner.Scan() {
		content += scanner.Text()
	}
	return content
}

func getOutputByStdinOrFile() string {
	var content string
	nonFlagArgs := flag.Args()

	if len(nonFlagArgs) == 0 {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			content += scanner.Text()
		}
		return content
	} else {
		content = readInFile(nonFlagArgs[0])
		return content
	}
}

func main() {
	setFlags()
	stdinWord := getOutputByStdinOrFile()
	args := lib.BoldArgs{
		TypeBold: flag.Lookup("t").Value.(flag.Getter).Get().(string),
		Percent:  flag.Lookup("f").Value.(flag.Getter).Get().(int),
		Jumps:    flag.Lookup("j").Value.(flag.Getter).Get().(int),
	}
	content := lib.MakeTextToBold(stdinWord, args)
	if flag.Lookup("o").Value.(flag.Getter).Get().(string) != "" {
		saveInFile(flag.Lookup("o").Value.(flag.Getter).Get().(string), content)
	} else {
		print(content)
	}
}
