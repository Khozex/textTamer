# Text Tamer

Transform ordinary text into bionic text with this project!

## Introduction


Bionic Reading is a revolutionary approach that enhances the reading experience by employing artificial fixation points to guide the eyes across text. With this method, readers primarily concentrate on the strategically highlighted initial letters, allowing their brain's natural processing center to effortlessly complete the words. Harness the power of this Go project to effortlessly transform your text into engaging bionic text.


## Features

- Read text from stdin: ✅
- Return into stdout Bold: ✅
- Specify an output archive for transformed text : ❌
- Receive an archive containing transformed text.: ❌
- Set word jumps for transformation: ❌
- Specify the output archive format (HTML, Markdown, and MD) : ❌


## Getting Started

To get started with cacheCraft, follow these steps:

1. Clone the repository to your local machine.

2. Build this project with command: ``go build``

3. Write a stind output and send to binary:
```bash
echo "This is my text" | ./textTamer 
//output
[**th**is **i**s **m**y **te**xt]
```
## License

This project is licensed under the [MIT License](https://www.mit.edu/~amini/LICENSE.md).