package main

import (
	"bufio"
	"os"
	"textTamer/lib"
)

func main() {
	var stdinWord string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		stdinWord += scanner.Text()
	}
	args := lib.BoldArgs{
		TypeBold: "md",
		Porcent:  50,
		Jumps:    0,
	}
	lib.MakeTextToBold(stdinWord, args)
}
