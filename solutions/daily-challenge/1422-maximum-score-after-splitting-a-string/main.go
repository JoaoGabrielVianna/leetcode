package main

import "fmt"

func maxScore(s string) int {
	var oneCount, maxScore int

	// Contar o total de 1s na string
	for _, char := range s {
		if char == '1' {
			oneCount++
		}
	}

	// Iterar pela string e calcular o score
	currentZeroCount := 0
	currentOneCount := oneCount

	// Não iterar até o último índice porque as partes não podem ser vazias
	for i := 0; i < len(s)-1; i++ {
		if s[i] == '0' {
			currentZeroCount++
		} else if s[i] == '1' {
			currentOneCount--
		}

		// Calcular o score atual
		currentScore := currentZeroCount + currentOneCount
		if currentScore > maxScore {
			maxScore = currentScore
		}
	}
	return maxScore
}

func main() {
	// Testes com entradas e saídas esperadas
	fmt.Println("Exemplo 1 (Entrada: '011101') -> Resultado:", maxScore("011101")) // Saída esperada: 5
	fmt.Println("Exemplo 2 (Entrada: '00111') -> Resultado:", maxScore("00111"))   // Saída esperada: 5
	fmt.Println("Exemplo 3 (Entrada: '1111') -> Resultado:", maxScore("1111"))     // Saída esperada: 3
}
