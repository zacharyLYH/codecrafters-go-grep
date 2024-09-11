package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func matchDigit(line []byte, pattern string) (bool, error) {
	if utf8.RuneCountInString(pattern) != 2 {
		return false, fmt.Errorf("unsupported pattern: %q", pattern)
	}

	ok := false

	for _, l := range line {
		if l >= 48 && l <= 57 {
			ok = true
			break
		}
	}

	return ok, nil
}

func matchAlphanumeric(line []byte, pattern string) (bool, error) {
	if utf8.RuneCountInString(pattern) != 2 {
		return false, fmt.Errorf("unsupported pattern: %q", pattern)
	}

	ok := false

	for _, l := range line {
		if (l >= 65 && l <= 90) || (l >= 97 && l <= 122) || l == 95 {
			ok = true
			break
		}
	}

	//if not a alphabet character, might be a numeric character
	if !ok {
		ok, _ = matchDigit(line, pattern)
	}

	return ok, nil
}

func matchLine(line []byte, pattern string) (bool, error) {
	if utf8.RuneCountInString(pattern) != 1 {
		return false, fmt.Errorf("unsupported pattern: %q", pattern)
	}

	ok := bytes.ContainsAny(line, pattern)

	return ok, nil
}
