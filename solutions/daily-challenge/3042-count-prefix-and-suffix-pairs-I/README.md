# 3042. Pares de prefixo e sufixo de contagem I

**`Fácil`**

Você recebe uma matriz de strings **com índice 0** `words` .

Vamos definir uma função **booleana** `isPrefixAndSuffix` que recebe duas strings `str1` e `str2`:

- `isPrefixAndSuffix(str1, str2)` retorna `true` se `str1` for um​**prefixo** e um **sufixo** de str2, e falsede outra forma.

Por exemplo, isPrefixAndSuffix("aba", "ababa")é trueporque "aba"é um prefixo de "ababa"e também um sufixo, mas isPrefixAndSuffix("abc", "abcd")é false.

Retorna um inteiro denotando o número de pares de índices (i, j)tais que i < j, e isPrefixAndSuffix(words[i], words[j])é true.

 

### Exemplo 1:

**Entrada:** palavras = ["a","aba","ababa","aa"]

**Saída:** 4

**Explicação:** 

Neste exemplo, os pares de índices contados são:
i = 0 e j = 1 porque isPrefixAndSuffix("a", "aba") é verdadeiro.
i = 0 e j = 2 porque isPrefixAndSuffix("a", "ababa") é verdadeiro.
i = 0 e j = 3 porque isPrefixAndSuffix("a", "aa") é verdadeiro.
i = 1 e j = 2 porque isPrefixAndSuffix("aba", "ababa") é verdadeiro.
Portanto, a resposta é 4.

---

### Exemplo 2:

**Entrada:** palavras = ["pa","papa","ma","mama"]

**Saída:** 2

**Explicação:** 

Neste exemplo, os pares de índices contados são:
i = 0 e j = 1 porque isPrefixAndSuffix("pa", "papa") é verdadeiro.
i = 2 e j = 3 porque isPrefixAndSuffix("ma", "mama") é verdadeiro.
Portanto, a resposta é 2. 

---

### Exemplo 3:

**Entrada:** words = ["abab","ab"]

**Saída:** 0

**Explicação:** 

Neste exemplo, o único par de índices válido é i = 0 e j = 1, e isPrefixAndSuffix("abab", "ab") é falso.
Portanto, a resposta é 0.
 

### Restrições:

- `1 <= words.length <= 50`
- `1 <= words[i].length <= 10`
- `words[i]` consiste apenas em letras minúsculas do inglês.

### Solução

```go
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
```
### Teste

```go
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
```