# 1408. String Matching in an Array

**`Fácil`**

Dado um array de strings `words`, o objetivo é retornar todas as strings que são substrings de outra string no mesmo array. A resposta pode ser retornada em qualquer ordem.

Uma substring é uma sequência contígua de caracteres dentro de uma string.

---

### Exemplo 1
**Entrada:** `words = ["mass", "as", "hero", "superhero"]`

**Saída:** `["as", "hero"]`

**Explicação:**  
- `"as"` é uma substring de `"mass"`.  
- `"hero"` é uma substring de `"superhero"`.  
Outra saída válida seria `["hero", "as"]`.

---

### Exemplo 2
**Entrada:** `words = ["leetcode", "et", "code"]`

**Saída:** `["et", "code"]`

**Explicação:**  
- `"et"` e `"code"` são substrings de `"leetcode"`.

---

### Exemplo 3
**Entrada:** `words = ["blue", "green", "bu"]`

**Saída:** `[]`

**Explicação:**  
Nenhuma string é substring de outra string no array.

---

## Restrições
- `1 <= words.length <= 100`
- `1 <= words[i].length <= 30`
- Cada string em `words` contém apenas letras minúsculas do alfabeto inglês.
- Todas as strings em `words` são únicas.

---

## Solução
```go
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
	words := []string{"mass", "as", "hero", "superhero"}
	fmt.Println(stringMatching(words)) // Output: ["as", "hero"]
}
