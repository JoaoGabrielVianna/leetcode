package main

import (
	"fmt"
)

// Função para contar os pares válidos
func countPrefixSuffixPairs(words []string) int {
	count := 0
	// Iterar sobre todos os pares (i, j) com i < j
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if isPrefixAndSuffix(words[i], words[j]) {
				count++
				fmt.Printf("Prefixo e sufixo encontrado: %s em %s\n", words[i], words[j])
			}
		}
	}
	return count
}

// Verifica se str1 é prefixo e sufixo de str2
func isPrefixAndSuffix(str1 string, str2 string) bool {
	// Verificar se str1 é prefixo de str2
	isPrefix := len(str1) <= len(str2) && str2[:len(str1)] == str1
	// Verificar se str1 é sufixo de str2
	isSuffix := len(str1) <= len(str2) && str2[len(str2)-len(str1):] == str1
	// Retorna verdadeiro apenas se for ambos
	return isPrefix && isSuffix
}

func main() {
	// Teste 1
	words1 := []string{"a", "aba", "ababa", "aa"}
	fmt.Println("Resultado esperado: 4, Saída:", countPrefixSuffixPairs(words1))

	// Teste 2
	words2 := []string{"pa", "papa", "ma", "mama"}
	fmt.Println("Resultado esperado: 2, Saída:", countPrefixSuffixPairs(words2))

	// Teste 3
	words3 := []string{"abab", "ab"}
	fmt.Println("Resultado esperado: 0, Saída:", countPrefixSuffixPairs(words3))
}
