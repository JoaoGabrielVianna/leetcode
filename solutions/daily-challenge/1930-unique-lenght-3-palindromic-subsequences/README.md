# 1930. Subsequências Palindrômicas de Comprimento Único-3

**`Médio`**

Dada uma string s, retorna o número de **palíndromos exclusivos de comprimento três** que são uma subsequência de `s`.

Observe que, mesmo que haja várias maneiras de obter a mesma subsequência, ela ainda será contada apenas **uma vez** .

Um **palíndromo** é uma sequência de caracteres que é lida da mesma forma de trás para frente.

Uma **subsequência** de uma string é uma nova string gerada a partir da string original com alguns caracteres (pode ser nenhum) excluídos sem alterar a ordem relativa dos caracteres restantes.

Por exemplo, `"ace"`é uma subsequência de .`"abcde"`
 
---

### Exemplo 1:

**Entrada:** s = "aabca"

**Saída: 3**

**Explicação:** 

As 3 subsequências palíndrômicas de comprimento 3 são:
- "aba" (subsequência de " a a b c a ")
- "aaa" (subsequência de " aa bc a ")
- "aca" (subsequência de " a ab ca ")

---
### Exemplo 2:

**Entrada:** s = "adc"

**Saída:** 0

**Explicação:** 

Não há subsequências palíndrômicas de comprimento 3 em "adc".

---

### Exemplo 3:

**Entrada:** s = "bbcbaba"

**Saída:** 4

**Explicação:** 

As 4 subsequências palíndromas de comprimento 3 são:
- "bbb" (subsequência de " bb c b aba")
- "bcb" (subsequência de " b b cb aba")
- "bab" (subsequência de " b bcb ab a")
- "aba" (subsequência de "bbcb aba ")
 

Restrições:

- `3 <= s.length <= 105`
- `s` consiste apenas em letras minúsculas do inglês.

## Solução
```go
func countPalindromicSubsequence(s string) int {
	var result int
	palindrome := []string{}

	// Percorre a string, verificando todas as combinações de 3 caracteres
	for i := 0; i < len(s)-2; i++ {
		for j := i + 1; j < len(s)-1; j++ {
			for k := j + 1; k < len(s); k++ {
				first := s[i] // Pega o primeiro caractere
				v := s[j]     // Pega o caractere do meio
				last := s[k]  // Pega o último caractere

				// Verifica se a subsequência é palíndroma
				if first == last {
					subseq := string(first) + string(v) + string(last)

					// Verifica se a subsequência já foi adicionada à lista
					found := false
					for _, p := range palindrome {
						if p == subseq {
							found = true
							break
						}
					}

					// Se não encontrou, adiciona à lista e incrementa o contador
					if !found {
						palindrome = append(palindrome, subseq)
						result++
					}
				}
			}
		}
	}

	// Imprime todos os palíndromos encontrados
	fmt.Println("Palíndromos encontrados:", palindrome)

	return result
}
```


## TESTE

```go
func main() {
	s := "bbcbaba"
	fmt.Println("Número de palíndromos de 3 caracteres:", countPalindromicSubsequence(s))
}

````