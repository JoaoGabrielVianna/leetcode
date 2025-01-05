# 2381. Mudança de Letras II

**`Médio`**

Você recebe uma string `s` composta por letras minúsculas do alfabeto inglês e um array bidimensional de inteiros `shifts`, onde cada elemento é definido como `shifts[i] = [starti, endi, directioni]`. Para cada `i`, desloque os caracteres na string `s` do índice `starti` até o índice `endi` (inclusive).

- Se `directioni = 1`, desloque os caracteres para **frente** (substituindo pela próxima letra do alfabeto).  
- Se `directioni = 0`, desloque os caracteres para **trás** (substituindo pela letra anterior no alfabeto).

Deslocar para frente significa avançar para a próxima letra no alfabeto (com retorno circular, onde `'z'` se transforma em `'a'`).  
De forma semelhante, deslocar para trás significa retroceder para a letra anterior no alfabeto (com retorno circular, onde `'a'` se transforma em `'z'`).

Retorne a string final após aplicar todas as mudanças descritas no array `shifts` à string `s`.

---

### Exemplo 1:

**Entrada:** s = "abc", shifts = [[0,1,0],[1,2,1],[0,2,1]]

**Saída:** "ace"

**Explicação:** 

Primeiramente, desloque os caracteres do índice 0 para o índice 1 para trás. Agora s = "zac".

Em segundo lugar, mova os caracteres do índice 1 para o índice 2 para a frente. Agora s = "zbd".
Por fim, mova os caracteres do índice 0 para o índice 2 para a frente. Agora s = "ace".

---

### Exemplo 2:

**Entrada:** s = "dztz", shifts = [[0,0,0],[1,1,1]]

**Saída:** "catz"

**Explicação:** 

Primeiramente, desloque os caracteres do índice 0 para o índice 0 para trás. Agora s = "cztz".

Por fim, mova os caracteres do índice 1 para o índice 1 para a frente. Agora s = "catz".
 

### Restrições:

- `1 <= s.length, shifts.length <= 5 * 104`
- `shifts[i].length == 3`
- `0 <= starti <= endi < s.length`
- `0 <= directioni <= 1`
- `s` consiste em letras minúsculas em inglês.


## SOLUÇÃO

```go
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
```

## TESTE

```go
func main() {
	s := "abc"
	shifts := [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}

	fmt.Println(shiftingLetters(s, shifts)) // Esperado: "ace"
}

```