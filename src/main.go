package main

import (
	"bufio"
	"flag"
	"os"
	"textTamer/lib"
)

func setFlags() {
	flag.Int("f", 50, "Porcent of word to bold")
	flag.String("t", "md", "Type of bold")
	flag.Int("j", 0, "Jumps of word to bold")
	flag.Parse()
}

func main() {
	setFlags()
	var stdinWord string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stdinWord += scanner.Text()
	}
	args := lib.BoldArgs{
		TypeBold: flag.Lookup("t").Value.(flag.Getter).Get().(string),
		Porcent:  flag.Lookup("f").Value.(flag.Getter).Get().(int),
		Jumps:    flag.Lookup("j").Value.(flag.Getter).Get().(int),
	}
	lib.MakeTextToBold(stdinWord, args)
}
