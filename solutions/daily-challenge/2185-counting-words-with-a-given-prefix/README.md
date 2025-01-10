# 2185. Counting Words With a Given Prefix

**`Easy`**

Você recebe um array de strings wordse uma string pref.

Retorna o número de strings wordsque contêm prefcomo prefixo .

Um prefixo de uma string sé qualquer substring contígua inicial de s.

### Exemplo 1:

**Entrada:** words = ["pay","attention","practice","attend"], pref = "at"

**Saída:** 2

As 2 strings que contêm "at" como prefixo são: " em "attention" e "attend".

---

### Exemplo 2:

**Entrada:** words = ["leetcode","win","loops","success"], pref = "code"

**Saída:** 0

**Explicação:** Não há strings que contenham "code" como prefixo.
 
---

### Restrições:

- `1 <= words.length <= 100`
- `1 <= words[i].length, pref.length <= 100`
- `words[i]` and `pref` consistem em letras minúsculas em inglês.

### Solução
```go
func prefixCount(words []string, pref string) int {
	count := 0
	for _, word := range words {
		if strings.HasPrefix(word, pref) {
			fmt.Println(word)
			count++
		}
	}
	return count
}
```

### Teste
```go
func main() {
	// Casos de teste organizados como pares de palavras e prefixos
	testCases := []struct {
		words []string
		pref  string
	}{
		{[]string{"pay", "attention", "practice", "attend"}, "at"},
		{[]string{"leetcode", "win", "loops", "success"}, "code"},
	}

	// Itera sobre os casos de teste
	for i, testCase := range testCases {
		fmt.Printf("Caso #%d: Palavras: %v | Prefixo: %q\n", i+1, testCase.words, testCase.pref)
		count := prefixCount(testCase.words, testCase.pref)
		fmt.Printf("Palavras que começam com %q: %d\n\n", testCase.pref, count)
	}
}
```