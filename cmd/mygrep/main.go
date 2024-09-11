package main

import (
	"os"
)

// Usage: echo <input_text> | your_program.sh -E <pattern>
func main() {
	line, pattern := readLineBoilerplate()
	// line := []byte("log")
	// pattern := "^log"
	ok := doGrep(line, pattern)
	exitOnError(ok)
	os.Exit(0)
}

func doGrep(line []byte, pattern string) bool {
	//Match the entire pattern from the front
	if pattern[0] == '^' {
		pattern = pattern[1:]
		ok, err := matchMultipleCharacterClasses(line, pattern)
		handleGenericError(err)
		return ok
	}
	//Match the entire pattern from the back
	if pattern[len(pattern)-1] == '$' {
		pattern = pattern[:len(pattern)-1]
		ok, err := matchMultipleCharacterClasses(reverseBytes(line), reverseString(pattern))
		handleGenericError(err)
		return ok
	}
	patLength := patternLength(pattern)
	for idx, l := range line {
		if patLength > len(line)-idx {
			return false
		}
		match, err := match([]byte{l}, pattern)
		handleGenericError(err)
		if match {
			ok, err := matchMultipleCharacterClasses(line[idx:], pattern)
			handleGenericError(err)
			if ok {
				return true
			}
		}
	}
	return false
}

func matchMultipleCharacterClasses(line []byte, pattern string) (bool, error) {
	var ok bool
	var err error
	lineCounter := 0
	for i := 0; i < len(pattern); {
		char := []byte{line[lineCounter]}
		pat := ""
		if pattern[i] == '\\' {
			pat = pattern[i : i+2]
			i += 2
		} else if pattern[i] == '[' {
			end := -1
			for j := i; j < len(pattern); j++ {
				if pattern[j] == ']' {
					end = j
					break
				}
			}
			pat = pattern[i:end]
			i = end + 1
		} else {
			pat = string(pattern[i])
			i++
		}
		ok, err = match(char, pat)
		if !ok {
			return false, nil
		}
		handleGenericError(err)
		lineCounter++
	}
	return ok, nil
}
