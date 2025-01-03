package main

import (
	"fmt"
)

func vowelStrings(words []string, queries [][]int) []int {
	// Função para verificar se uma palavra começa e termina com vogal
	isVowel := func(r rune) bool {
		return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u'
	}
	// Pré-processar para contar palavras válidas cumulativamente
	n := len(words)
	vowelCount := make([]int, n+1) // Armazena contagem cumulativa de palavras válidas

	for i, word := range words {
		if isVowel(rune(word[0])) && isVowel(rune(word[len(word)-1])) {
			vowelCount[i+1] = vowelCount[i] + 1
		} else {
			vowelCount[i+1] = vowelCount[i]
		}
	}

	// Responder consultas
	results := make([]int, len(queries))

	for i, query := range queries {
		start, end := query[0], query[1]
		results[i] = vowelCount[end+1] - vowelCount[start]
	}

	return results
}

func main() {
	words1 := []string{"aba", "bcb", "ece", "aa", "e"}
	consults1 := [][]int{{0, 2}, {1, 4}, {1, 1}}

	words2 := []string{"a", "e", "i"}
	consults2 := [][]int{{0, 2}, {0, 1}, {2, 2}}

	fmt.Println("EXEMPLO 1 (Entrada: palavras = ['aba','bcb','ece','aa','e'])") // Saída esperada: [2,3,0]
	fmt.Println("RESULTADO -> ", vowelStrings(words1, consults1))
	fmt.Println("EXEMPLO 2 (Entrada: palavras = ['a','e','i'])") // Saída esperada: [3,2,1]
	fmt.Println("RESULTADO -> ", vowelStrings(words2, consults2))
}
