# 2116. Verifique se uma sequência de parênteses pode ser válida

**`Médio`**

Uma string de parênteses é uma string não vazia que consiste somente de `'('` e `')'`. Ela é válida se qualquer uma das seguintes condições for verdadeira :

- Isso é `()`.
- Pode ser escrito como `AB`( `A` concatenado com `B`), onde `A` e `B` são strings de parênteses válidas.
- Pode ser escrito como (A), onde Aé uma string de parênteses válida.
Você recebe uma string de parênteses `s` e uma string `locked`, ambas de comprimento `n`. `locked` é uma string binária que consiste apenas em `'0'`s e `'1'`s. Para cada índice `i` de `locked`,

- Se` locked[i]` for `'1'`, você não pode mudar `s[i]`.
- Mas se locked[i]for '0', você pode mudar s[i]para `'('` ou `')'`.
Retorna `true` se você puder fazer suma string de parênteses válida . Caso contrário, retorna `false`.

 

### Exemplo 1:


**Entrada:** s = "))()))", bloqueado = "010100"

**Saída:** verdadeiro

**Explicação:** 

bloqueado[1] == '1' e bloqueado[3] == '1', então não podemos alterar s[1] ou s[3]. 

Alteramos s[0] e s[4] para '(' enquanto deixamos s[2] e s[5] inalterados para tornar s válido.

---

### Exemplo 2:

**Entrada:** s = "()()", bloqueado = "0000"

**Saída:** verdadeiro

**Explicação:** 

Não precisamos fazer nenhuma alteração porque s já é válido.

---

### Exemplo 3:

**Entrada:** s = ")", bloqueado = "0"

**Saída:** falso

**Explicação:** 

bloqueado nos permite alterar s[0]. 

Alterar s[0] para '(' ou ')' não tornará s válido.

---

### Restrições:

- `n == s.length == locked.length`
- `1 <= n <= 105`
- `s[i]é '('ou ')'.`
- `locked[i]é '0'ou '1'.`

### Solução

```go
func canBeValid(s string, locked string) bool {
	if len(s)%2 != 0 {
		return false
	}
	openCount, unlockCount := 0, 0
	for i := 0; i < len(s); i++ {
		if locked[i] == '0' {
			unlockCount++
		} else if s[i] == '(' {
			openCount++
		} else {
			if openCount > 0 {
				openCount--
			} else if unlockCount > 0 {
				unlockCount--
			} else {
				return false
			}
		}
	}
	openCount, unlockCount = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if locked[i] == '0' {
			unlockCount++
		} else if s[i] == ')' {
			openCount++
		} else {
			if openCount > 0 {
				openCount--
			} else if unlockCount > 0 {
				unlockCount--
			} else {
				return false
			}
		}
	}
	return true
}
```

### Teste

```go
func main() {
	testCases := []struct {
		s      string
		locked string
	}{
		{s: "))()))", locked: "010100"},
		{s: "()()", locked: "0000"},
		{s: ")", locked: "0"},
	}

	for _, testCase := range testCases {
		fmt.Println(canBeValid(testCase.s, testCase.locked))
	}
}
```