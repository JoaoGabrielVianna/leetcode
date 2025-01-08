package main

import (
	"fmt"
	"strings"
)

func stringMatching(words []string) []string {
	var result []string
	for i1, word1 := range words {
		for i2, word2 := range words {
			if i1 != i2 && strings.Contains(word2, word1) {
				// Adiciona apenas se ainda não estiver na lista de resultados
				if !contains(result, word1) {
					result = append(result, word1)
				}
			}
		}
	}
	return result
}

// Função auxiliar para verificar se o resultado já contém a palavra
func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}

func main() {
	words1 := []string{"mass", "as", "hero", "superhero"}
	fmt.Println(stringMatching(words1)) // Output: ["as", "hero"]

	words2 := []string{"leetcode", "et", "code"}
	fmt.Println(stringMatching(words2)) // Output: ["et","code"]

	words3 := []string{"blue", "green", "bu"}
	fmt.Println(stringMatching(words3)) // Output: []

}
