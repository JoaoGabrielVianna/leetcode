# 916. Subconjuntos de palavras

**`Médio`**

Você recebe dois arrays de strings `words1` e `words2`.

Uma string bé um **subconjunto** de string ase cada letra em bocorre em, aincluindo multiplicidade.

Por exemplo, `"wrr"` é um subconjunto de "warrior" mas não é um subconjunto de `"world"`.
Uma string `a` de `words1` é **universal** se para cada string `b` em `words2`, bfor um subconjunto de a.

Retorna uma matriz de todas as strings **universais** em `words1`. Você pode retornar a resposta em qualquer ordem .

 

### Exemplo 1:

**Entrada:** palavras1 = ["amazon","apple","facebook","google","leetcode"], palavras2 = ["e","o"]

**Saída:** ["facebook","google","leetcode"]

---

### Exemplo 2:

**Entrada:** palavras1 = ["amazon","apple","facebook","google","leetcode"], palavras2 = ["l","e"]

**Saída:** ["apple","google","leetcode"]
 
---

### Restrições:

- `1 <= words1.length, words2.length <= 104`
- `1 <= words1[i].length, words2[i].length <= 10`
- `words1[i]` e `words2[i]` consistem apenas em letras minúsculas do inglês.
- Todas as strings de `words1` são únicas .

---

### Solução
```go
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
```


### Teste
```go
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

```