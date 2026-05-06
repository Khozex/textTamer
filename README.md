# Text Tamer

Transform ordinary text into bionic text with this project!

## Introduction

Bionic Reading is a revolutionary approach that enhances the reading
experience by employing artificial fixation points to guide the eyes
across text. With this method, readers primarily concentrate on the
strategically highlighted initial letters, allowing their brain's natural
processing center to effortlessly complete the words. Harness the power
of this Go project to effortlessly transform your text into engaging
bionic text.

## Getting Started

1. Clone the repository.
2. Build it: `cd src && go build -o ../textTamer .`
3. Pipe text into the binary or pass a file path:

```bash
echo "This is my text" | ./textTamer
# **Th**is **i**s **m**y **te**xt
```

## Usage

```
textTamer [flags] [input-file]
```

If no `input-file` is given, the program reads from stdin.

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-t` | `md`    | Bold marker type: `md` (`**word**`), `html` (`<b>word</b>`), or `ansi` (terminal escape sequences). |
| `-f` | `0`     | Bold percent. `0` (default) uses the bionic word-length curve. `1`-`100` forces a flat percent across all words. |
| `-j` | `0`     | Skip this many words between bolded ones. |
| `-o` | (none)  | Output file. If unset, writes to stdout. |

### Default behavior (bionic curve)

By default, the bold prefix length scales with each word's length, matching
the original Bionic Reading method:

| Word length | Letters bolded |
|---|---|
| 1 | 0 |
| 2-3 | 1 |
| 4 | 2 |
| 5-6 | 3 |
| 7-9 | 4 |
| ≥ 10 | ~40% |

To opt back into the old flat-50 behavior:

```bash
echo "The sun rises in the east" | ./textTamer -f 50
```

## License

This project is licensed under the [MIT License](https://www.mit.edu/~amini/LICENSE.md).
