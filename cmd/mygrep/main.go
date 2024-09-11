package main

import (
	"os"
)

// Usage: echo <input_text> | your_program.sh -E <pattern>
func main() {
	line, pattern := readLineBoilerplate()
	var ok bool
	var err error
	if pattern == "\\d" {
		ok, err = matchDigit(line, pattern)
	} else if pattern == "\\w" {
		ok, err = matchAlphanumeric(line, pattern)
	} else if startsAndEndsWith(pattern, '[', ']') {
		ok, err = matchCharacterGroup(line, pattern[1:len(pattern)-1])
	} else {
		ok, err = matchLine(line, pattern)
	}
	handleGenericError(err)
	exitOnError(ok)
	os.Exit(0)
}
