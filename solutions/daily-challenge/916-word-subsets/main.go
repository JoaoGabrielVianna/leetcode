package main

import (
	"fmt"
)

func wordSubsets(words1 []string, words2 []string) []string {
	// Mapa para armazenar a frequência máxima de cada letra em words2
	maxFreq := make(map[rune]int)

	// Calcula a frequência máxima de cada letra em todas as palavras de words2
	for _, word := range words2 {
		tempFreq := make(map[rune]int)
		for _, char := range word {
			tempFreq[char]++
			if tempFreq[char] > maxFreq[char] {
				maxFreq[char] = tempFreq[char]
			}
		}
	}

	// Verifica cada palavra em words1
	result := []string{}
	for _, word := range words1 {
		wordFreq := make(map[rune]int)
		for _, char := range word {
			wordFreq[char]++
		}

		isUniversal := true
		for char, count := range maxFreq {
			if wordFreq[char] < count {
				isUniversal = false
				break
			}
		}

		if isUniversal {
			result = append(result, word)
		}
	}

	return result
}

func main() {

	testCases := []struct {
		words1 []string
		words2 []string
	}{
		{words1: []string{"amazon", "apple", "facebook", "google", "leetcode"}, words2: []string{"e", "o"}},
		{words1: []string{"amazon", "apple", "facebook", "google", "leetcode"}, words2: []string{"l", "e"}},
	}

	for _, testCase := range testCases {
		result := wordSubsets(testCase.words1, testCase.words2)
		fmt.Printf("Input words1: %v\n", testCase.words1)
		fmt.Printf("Input words2: %v\n", testCase.words2)
		fmt.Printf("Universal words: %v\n\n", result)
	}
}
