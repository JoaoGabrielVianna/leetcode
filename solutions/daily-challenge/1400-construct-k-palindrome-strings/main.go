package main

import "fmt"

func canConstruct(s string, k int) bool {
	// Se k é maior que o tamanho da string, não é possível criar tantas strings.
	if k > len(s) {
		return false
	}

	// Frequência de cada caractere
	freq := make(map[rune]int)
	for _, char := range s {
		freq[char]++
	}

	// Conte quantos caracteres têm frequência ímpar
	oddCount := 0
	for _, count := range freq {
		if count%2 != 0 {
			oddCount++
		}
	}

	// Verifique se é possível formar k palíndromos
	return oddCount <= k
}

func main() {
	testCases := []struct {
		s string
		k int
	}{
		{s: "annabelle", k: 2},
	}

	for _, testCase := range testCases {
		result := canConstruct(testCase.s, testCase.k)
		fmt.Printf("Input s: %v\n", testCase.s)
		fmt.Printf("Input k: %v\n", testCase.k)
		fmt.Printf("Output: %v\n\n", result)
	}

}
