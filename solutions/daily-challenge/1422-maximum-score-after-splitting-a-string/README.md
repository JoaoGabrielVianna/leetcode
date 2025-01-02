# 1422. Pontuação Máxima Após Dividir uma String

**`Facil`**

Dada uma sequência `s` composta de zeros (`0`) e uns (`1`), o objetivo é retornar a pontuação máxima ao dividir a sequência em duas subsequências **não vazias**: a subsequência esquerda e a subsequência direita.

### Exemplo 1
**Entrada:**  `s = "011101"`

**Saída:**  `5`

**Explicação:**  
Todas as formas possíveis de dividir `s` e suas respectivas pontuações:
- Esquerda = `"0"` e Direita = `"11101"`, Pontuação = `1 + 4 = 5`
- Esquerda = `"01"` e Direita = `"1101"`, Pontuação = `1 + 3 = 4`
- Esquerda = `"011"` e Direita = `"101"`, Pontuação = `1 + 2 = 3`
- Esquerda = `"0111"` e Direita = `"01"`, Pontuação = `1 + 1 = 2`
- Esquerda = `"01110"` e Direita = `"1"`, Pontuação = `2 + 1 = 3`

---

### Exemplo 2
**Entrada:** `s = "00111"`

**Saída:** `5`

**Explicação:**  
Quando Esquerda = `"00"` e Direita = `"111"`, a pontuação máxima é `2 + 3 = 5`.

---

## Exemplo 3
**Entrada:** `s = "1111"`

**Saída:** `3`

### Restrições
- `2 <= s.length <= 500`
- A sequência `s` consiste **apenas** em caracteres `'0'` e `'1'`.

---

## Solução
```go
package main

func maxScore(s string) int {
    var oneCount, maxScore int

    // Contar o total de 1s na string
    for _, char := range s {
        if char == '1' {
            oneCount++
        }
    }

    // Iterar pela string e calcular o score
    currentZeroCount := 0
    currentOneCount := oneCount

    for i := 0; i < len(s)-1; i++ {
        if s[i] == '0' {
            currentZeroCount++
        } else if s[i] == '1' {
            currentOneCount--
        }

        // Calcular o score atual
        currentScore := currentZeroCount + currentOneCount
        if currentScore > maxScore {
            maxScore = currentScore
        }
    }
    return maxScore
}
