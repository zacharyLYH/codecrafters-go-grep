package main

import (
	"bytes"
)

func match(line []byte, pattern string) (bool, error) {
	if pattern[0] == '\\' {
		if pattern[:2] == "\\d" {
			return matchDigit(line)
		}
		if pattern[:2] == "\\w" {
			return matchAlphanumeric(line)
		}
	}
	if pattern[0] == '[' {
		return matchCharacterGroup(line, pattern[1:])
	}
	return matchLine(line, pattern)
}

func matchDigit(line []byte) (bool, error) {
	ok := false

	for _, l := range line {
		if l >= 48 && l <= 57 {
			ok = true
			break
		}
	}

	return ok, nil
}

func matchAlphanumeric(line []byte) (bool, error) {
	ok := false

	for _, l := range line {
		if (l >= 65 && l <= 90) || (l >= 97 && l <= 122) || l == 95 {
			ok = true
			break
		}
	}

	//if not a alphabet character, might be a numeric character
	if !ok {
		ok, _ = matchDigit(line)
	}

	return ok, nil
}

func matchLine(line []byte, pattern string) (bool, error) {
	ok := bytes.ContainsAny(line, pattern)
	return ok, nil
}

// Returns true when the first item matches. Otherwise immediately throws exit(1) status
func matchCharacterGroup(line []byte, group string) (bool, error) {
	charGrp := make(map[rune]bool)
	isNegativeGrp := false
	for _, c := range group {
		charGrp[c] = true
		if c == '^' {
			isNegativeGrp = true
		}
	}
	for _, l := range line {
		if charGrp[rune(l)] && !isNegativeGrp {
			return true, nil
		} else if isNegativeGrp && charGrp[rune(l)] {
			exitOnError(false)
		}
	}
	if isNegativeGrp {
		return true, nil
	}
	return false, nil
}
