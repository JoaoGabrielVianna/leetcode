package main

import (
	"fmt"
	"strings"
)

func removeOccurrences(s string, part string) string {
	for strings.Contains(s, part) {
		s = strings.Replace(s, part, "", 1)
	}
	return s
}

func main() {
	testCases := []struct {
		s    string
		part string
	}{
		{s: "daabcbaabcbc", part: "abc"},
		{s: "axxxxyyyyb", part: "xy"},
		{s: "abcdef", part: "abc"},
	}

	for _, testCase := range testCases {
		result := removeOccurrences(testCase.s, testCase.part)
		fmt.Printf("Input: %q, %q\n", testCase.s, testCase.part)
		fmt.Printf("Output: %q\n\n", result)
	}
}
