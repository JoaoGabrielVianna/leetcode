package main

import (
	"fmt"
)

func countPalindromicSubsequence(s string) int {
	var result int
	palindrome := []string{}

	// Percorre a string, verificando todas as combinações de 3 caracteres
	for i := 0; i < len(s)-2; i++ {
		for j := i + 1; j < len(s)-1; j++ {
			for k := j + 1; k < len(s); k++ {
				first := s[i] // Pega o primeiro caractere
				v := s[j]     // Pega o caractere do meio
				last := s[k]  // Pega o último caractere

				// Verifica se a subsequência é palíndroma
				if first == last {
					subseq := string(first) + string(v) + string(last)

					// Verifica se a subsequência já foi adicionada à lista
					found := false
					for _, p := range palindrome {
						if p == subseq {
							found = true
							break
						}
					}

					// Se não encontrou, adiciona à lista e incrementa o contador
					if !found {
						palindrome = append(palindrome, subseq)
						result++
					}
				}
			}
		}
	}

	// Imprime todos os palíndromos encontrados
	fmt.Println("Palíndromos encontrados:", palindrome)

	return result
}

func main() {
	s1 := "aabca"
	s2 := "adc"
	s3 := "bbcbaba"
	fmt.Println("Número de palíndromos de 3 caracteres:", countPalindromicSubsequence(s1))
	fmt.Println("Número de palíndromos de 3 caracteres:", countPalindromicSubsequence(s2))
	fmt.Println("Número de palíndromos de 3 caracteres:", countPalindromicSubsequence(s3))
}
