package main

import (
	"os"
)

// Usage: echo <input_text> | your_program.sh -E <pattern>
func main() {
	line, pattern := readLineBoilerplate()
	switch pattern {
	case "\\d":
		ok, err := matchDigit(line, pattern)
		handleGenericError(err)
		exitOnError(ok)
	default:
		ok, err := matchLine(line, pattern)
		handleGenericError(err)
		exitOnError(ok)
	}
	os.Exit(0)
}
