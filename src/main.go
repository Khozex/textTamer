package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"textTamer/lib"
)

type Config struct {
	TypeBold   string
	Percent    int
	Jumps      int
	OutputFile string
	InputFile  string
}

func parseFlags() Config {
	var cfg Config
	flag.StringVar(&cfg.TypeBold, "t", "md", "Bold marker type: md | html | ansi")
	flag.IntVar(&cfg.Percent, "f", 0, "Bold percent (0 = bionic curve mode, 1-100 = flat percent)")
	flag.IntVar(&cfg.Jumps, "j", 0, "Number of words to skip between bolded words")
	flag.StringVar(&cfg.OutputFile, "o", "", "Output file (default: stdout)")
	flag.Parse()
	if args := flag.Args(); len(args) > 0 {
		cfg.InputFile = args[0]
	}
	return cfg
}

func validate(cfg Config) error {
	switch cfg.TypeBold {
	case "md", "html", "ansi":
	default:
		return fmt.Errorf("invalid -t %q (must be md, html, or ansi)", cfg.TypeBold)
	}
	if cfg.Percent < 0 || cfg.Percent > 100 {
		return fmt.Errorf("invalid -f %d (must be 0..100)", cfg.Percent)
	}
	if cfg.Jumps < 0 {
		return fmt.Errorf("invalid -j %d (must be >= 0)", cfg.Jumps)
	}
	return nil
}

func readAll(cfg Config) (string, error) {
	var src *os.File
	if cfg.InputFile != "" {
		f, err := os.Open(cfg.InputFile)
		if err != nil {
			return "", err
		}
		defer f.Close()
		src = f
	} else {
		src = os.Stdin
	}

	scanner := bufio.NewScanner(src)
	var content string
	for scanner.Scan() {
		content += scanner.Text() + "\n"
	}
	return content, scanner.Err()
}

func writeAll(cfg Config, content string) error {
	if cfg.OutputFile == "" {
		fmt.Print(content)
		return nil
	}
	f, err := os.Create(cfg.OutputFile)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(content)
	return err
}

func die(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	cfg := parseFlags()
	if err := validate(cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		flag.Usage()
		os.Exit(2)
	}

	input, err := readAll(cfg)
	if err != nil {
		die(err)
	}

	content := lib.MakeTextToBold(input, lib.BoldArgs{
		TypeBold: cfg.TypeBold,
		Percent:  cfg.Percent,
		Jumps:    cfg.Jumps,
	})

	if err := writeAll(cfg, content); err != nil {
		die(err)
	}
}
