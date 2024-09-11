package main

import (
	"fmt"
	"io"
	"os"
)

func handleGenericError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}
}

func exitOnError(ok bool) {
	if !ok {
		os.Exit(1)
	}
}

func readLineBoilerplate() ([]byte, string) {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: mygrep -E <pattern>\n")
		os.Exit(2) // 1 means no lines were selected, >1 means error
	}

	pattern := os.Args[2]

	line, err := io.ReadAll(os.Stdin) // assume we're only dealing with a single line
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read input text: %v\n", err)
		os.Exit(2)
	}

	return line, pattern
}

func reverseBytes(line []byte) []byte {
	left, right := 0, len(line)-1
	for left < right {
		// Swap elements
		line[left], line[right] = line[right], line[left]
		left++
		right--
	}
	return line
}

// Function to reverse a string
func reverseString(pattern string) string {
	// Convert string to a slice of runes to handle multi-byte characters
	runes := []rune(pattern)
	left, right := 0, len(runes)-1
	for left < right {
		// Swap runes
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	return string(runes)
}
