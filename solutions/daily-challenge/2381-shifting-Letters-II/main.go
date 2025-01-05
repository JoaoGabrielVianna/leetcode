package main

import (
	"fmt"
)

func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	shiftAccum := make([]int, n+1) // Array auxiliar para marcação de deslocamentos

	// Preenchendo o shiftAccum
	for _, shift := range shifts {
		start, end, direction := shift[0], shift[1], shift[2]

		if direction == 1 { // Deslocar para frente
			shiftAccum[start] += 1
			if end+1 < n {
				shiftAccum[end+1] -= 1
			}
		} else { // Deslocar para trás
			shiftAccum[start] -= 1
			if end+1 < n {
				shiftAccum[end+1] += 1
			}
		}
	}

	// Calculando a soma acumulativa
	for i := 1; i < n; i++ {
		shiftAccum[i] += shiftAccum[i-1]
	}

	// Aplicando os deslocamentos na string
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		shift := shiftAccum[i] % 26 // Garantir que o deslocamento esteja entre 0-25
		if shift < 0 {
			shift += 26 // Corrigir deslocamentos negativos
		}
		result[i] = byte((int(s[i]-'a')+shift)%26 + 'a')
	}

	return string(result)
}

func main() {
	s := "abc"
	shifts := [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}

	fmt.Println(shiftingLetters(s, shifts)) // Esperado: "ace"
}
