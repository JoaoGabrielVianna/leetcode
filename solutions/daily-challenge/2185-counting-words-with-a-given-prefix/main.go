package main

import (
	"fmt"
	"strings"
)

func prefixCount(words []string, pref string) int {
	count := 0
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			fmt.Println(word)
			count++
		}
	}
	return count
}

func main() {
	// Casos de teste organizados como pares de palavras e prefixos
	testCases := []struct {
		words []string
		pref  string
	}{
		{[]string{"pay", "attention", "practice", "attend"}, "at"},
		{[]string{"leetcode", "win", "loops", "success"}, "code"},
	}

	// Itera sobre os casos de teste
	for i, testCase := range testCases {
		fmt.Printf("Caso #%d: Palavras: %v | Prefixo: %q\n", i+1, testCase.words, testCase.pref)
		count := prefixCount(testCase.words, testCase.pref)
		fmt.Printf("Palavras que começam com %q: %d\n\n", testCase.pref, count)
	}
}
