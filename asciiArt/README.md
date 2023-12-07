# ASCII Art Generator

ASCII Art Generator is a command-line program written in Go that converts text into ASCII art using specified ASCII art fonts.

## How to Run

First, you need to install [Go](https://golang.org/doc/install) to your system if it is not already installed.

To run the program, open a terminal and navigate to the directory that contains `main.go`. Input text and ASCII art style are passed as command line arguments like this:

```bash
go run main.go "[TEXT]" "[OPTIONAL: ASCII_ART_STYLE]"

Replace [TEXT] with the text you want to convert to ASCII art. The text should not contain any spaces.

You can optionally replace [OPTIONAL: ASCII_ART_STYLE] with the ASCII art style. The program currently supports these styles: standard, shadow, thinkertoy. If no style is specified, the program will use standard as the default style.

Examples
Convert "HELLO" into ASCII art with standard style:
go run main.go "HELLO"

Convert "HELLO" into ASCII art with shadow style:
go run main.go "HELLO" shadow

Convert "HELLO" into ASCII art with thinkertoy style:
go run main.go "HELLO" thinkertoy

The program will print the ASCII art text to the terminal.

Available ASCII Art Styles
The following ASCII art styles are available:

standard
shadow
thinkertoy


Unit Test
You can test the program by running the following command:
go test

This will run a series of unit tests that verify the ASCII Art output for accuracy.