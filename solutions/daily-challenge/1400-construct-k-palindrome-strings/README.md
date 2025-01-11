# 1400. Construir K Cordas Palíndromas
**`Médio`**

Dada uma string `s` e um inteiro `k`, retorne `true` se você pode usar todos os caracteres `s` para construir `k` strings de palíndromo ou `false` de outra forma .

### Exemplo 1:

**Entrada:** s = "annabelle", k = 2
**Saída:** true
**Explicação:** Você pode construir dois palíndromos usando todos os caracteres em s. 
Algumas construções possíveis "anna" + "elble", "anbna" + "elle", "anellena" + "b"
### Exemplo 2:

**Entrada:** s = "leetcode", k = 3
 *Saída:** falso
**Explicação:** É impossível construir 3 palíndromos usando todos os caracteres de s.
Exemplo 3:

**Entrada:** s = "true", k = 4
**Saída:** true
**Explicação:** A única solução possível é colocar cada caractere em uma string separada.
 

### Restrições:

- `1 <= s.length <= 105`
- `s` consiste em letras minúsculas em inglês.
- `1 <= k <= 105`

### Solução
```go
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
```
### Teste
```go
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

```