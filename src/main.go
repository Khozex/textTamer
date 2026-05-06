package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
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

func die(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func openInput(cfg Config) (io.ReadCloser, error) {
	if cfg.InputFile == "" {
		return io.NopCloser(os.Stdin), nil
	}
	return os.Open(cfg.InputFile)
}

func openOutput(cfg Config) (io.WriteCloser, error) {
	if cfg.OutputFile == "" {
		return nopWriteCloser{os.Stdout}, nil
	}
	return os.Create(cfg.OutputFile)
}

type nopWriteCloser struct{ io.Writer }

func (nopWriteCloser) Close() error { return nil }

func processStream(r io.Reader, w io.Writer, args lib.BoldArgs) error {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 1024*1024), 10*1024*1024)
	bw := bufio.NewWriter(w)
	for scanner.Scan() {
		line := lib.MakeTextToBold(scanner.Text(), args)
		if _, err := fmt.Fprintln(bw, line); err != nil {
			return err
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return bw.Flush()
}

func main() {
	cfg := parseFlags()
	if err := validate(cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		flag.Usage()
		os.Exit(2)
	}

	in, err := openInput(cfg)
	if err != nil {
		die(err)
	}
	defer in.Close()

	out, err := openOutput(cfg)
	if err != nil {
		die(err)
	}

	args := lib.BoldArgs{
		TypeBold: cfg.TypeBold,
		Percent:  cfg.Percent,
		Jumps:    cfg.Jumps,
	}
	if err := processStream(in, out, args); err != nil {
		out.Close()
		die(err)
	}
	if err := out.Close(); err != nil {
		die(err)
	}
}
