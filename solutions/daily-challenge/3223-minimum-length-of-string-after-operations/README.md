# 3223. Comprimento mínimo da sequência após as operações

**`Médio`**

Você recebe uma string `s`.

Você pode executar o seguinte processo `s` **quantas** vezes quiser:

- Escolha um índice `i` na string de modo que haja pelo menos um caractere à esquerda do índice `i` que seja igual a `s[i]`, e pelo menos um caractere à direita que também seja igual a `s[i]`.
- Exclua o caractere mais próximo à esquerda do índice ique seja igual a `s[i]`.
- Exclua o caractere mais próximo à direita do índice ique seja igual a `s[i]`.
Retorne o comprimento mínimo da string final sque você pode atingir.

 

### Exemplo 1:

**Entrada:** s = "abaacbcbb"

**Saída:** 5

**Explicação:**

Fazemos as seguintes operações:

- Escolha o índice 2 e remova os caracteres nos índices 0 e 3. A string resultante é s = "bacbcbb".
- Escolha o índice 3 e remova os caracteres nos índices 0 e 5. A string resultante é s = "acbcb".

---

### Exemplo 2:

**Entrada:** s = "aa"

**Saída:** 2

**Explicação:**

Não podemos realizar nenhuma operação, então retornamos o comprimento da string original.

---

### Restrições:

- `1 <= s.length <= 2 * 105`
- `s` consiste apenas em letras minúsculas do inglês.
  
### Solução
```go
func minimumLength(s string) int {
	counts := make([]int, 26)
	for i := 0; i < len(s); i++ {
		counts[s[i]-'a']++
	}
	res := 0

	for _, val := range counts {
		if val <= 2 {
			res += val
		} else {
			if val%2 == 0 {
				res += 2
			} else {
				res++
			}
		}
	}
	return res
}

```
### Teste
```go
func main() {
	testCases := []string{
		"abaacbcbb",
	}

	for _, s := range testCases {
		result := minimumLength(s)
		fmt.Printf("Input s: %s\n", s)
		fmt.Printf("Output: %v\n", result)
		fmt.Println("-----")
	}
}
```