package main

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		strtChar := s[start]
		endChar := s[end]

		if unicode.IsSpace(rune(strtChar)) || unicode.IsPunct(rune(strtChar)) || unicode.IsSymbol(rune(strtChar)) {
			start++
			continue
		}

		if unicode.IsSpace(rune(endChar)) || unicode.IsPunct(rune(endChar)) || unicode.IsSymbol(rune(endChar)) {
			end--
			continue
		}

		if unicode.ToLower(rune(strtChar)) != unicode.ToLower(rune(endChar)) {
			return false
		}

		start++
		end--
	}

	return true
}

func main() {
	fmt.Printf("%s is Palindrome %v\n", "0P", isPalindrome("0P"))
}
