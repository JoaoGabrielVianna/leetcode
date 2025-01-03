# 2559. Count Vowel Strings in Ranges

**`Medium`**

Você recebe uma matriz de strings com índice 0`words` e uma matriz 2D de inteiros `queries`.

Cada consulta nos pede para encontrar o número de strings presentes no intervalo ( ambas inclusive ) que começam e terminam com uma vogal.`queries[i] = [lᵢ, rᵢ]` `lᵢ` `rᵢ` `words`

Retorna uma matriz `ans` de tamanho `queries.length`, onde `ans[i]` é a resposta para `i`ᵃ consulta .

**Observe** que as letras vogais são `'a'`, `'e'`, `'i'`, `'o'`, e `'u'`.

### Exemplo 1:

**Entrada:** palavras = ["aba","bcb","ece","aa","e"], consultas = [[0,2],[1,4],[1,1]]

**Saída:** [2,3,0]

**Explicação:** 

As sequências que começam e terminam com uma vogal são "aba", "ece", "aa" e "e".
A resposta para a consulta [0,2] é 2 (strings "aba" e "ece").
para consultar [1,4] é 3 (strings "ece", "aa", "e").
para consultar [1,1] é 0.
Retornamos [2,3,0].

---

### Exemplo 2:

**Entrada:** palavras = ["a","e","i"], consultas = [[0,2],[0,1],[2,2]]

**Saída:** [3,2,1]

**Explicação:** Cada string satisfaz as condições, então retornamos [3,2,1].
 
---

### Restrições:

- `1 <= words.length <= 10⁵`
- ` 1 <= words[i].length <= 40`
- `words[i] `consiste apenas em letras minúsculas do inglês.
- `sum(words[i].length) <= 3 * 10⁵`
- `1 <= queries.length <= 10⁵`
- `0 <= lᵢ <= rᵢ < words.length`

---

## Solução
```go
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

```

## Teste

```go
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
```