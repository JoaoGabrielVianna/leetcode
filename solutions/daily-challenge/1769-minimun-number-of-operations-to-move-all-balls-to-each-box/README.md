# 1769. Minimum Number of Operations to Move All Balls to Each Box

**`Médio`**

Você tem ncaixas. Você recebe uma string binária `boxes` de comprimento `n`, onde `boxes[i]` é `'0'` se a caixa estiver vazia e se ela contiver uma bola.`ith` `'1'`

Em uma operação, você pode mover uma bola de uma caixa para uma caixa adjacente. Caixa `i` é adjacente à caixa `j` se `abs(i - j) == 1`. Note que depois de fazer isso, pode haver mais de uma bola em algumas caixas.

Retorna uma matriz `answer` de tamanho n, onde `answer[i]` é o número mínimo de operações necessárias para mover todas as bolas para a caixa.ith

Cada um `answer[i]` é calculado considerando o estado inicial das caixas.

---

### Exemplo 1:

**Entrada:** boxes = "110"

**Saída:** [1,1,3]

**Explicação:** 

A resposta para cada caixa é a seguinte: 
1) Primeira caixa: você terá que mover uma bola da segunda caixa para a primeira caixa em uma operação. 
2) Segunda caixa: você terá que mover uma bola da primeira caixa para a segunda caixa em uma operação. 
3) Terceira caixa: você terá que mover uma bola da primeira caixa para a terceira caixa em duas operações e mover uma bola da segunda caixa para a terceira caixa em uma operação.


---

### Exemplo 2:

**Entrada:** caixas = "001011"

**Saída:** [11,8,5,4,3,4]


### Restrições:

- `n == boxes.length`
- `1 <= n <= 2000`
- `boxes[i]`é `'0'`ou `'1'`.

## Solução

```go
func minOperations(boxes string) []int {
	n := len(boxes)
	answer := make([]int, n)

	for i := 0; i < n; i++ {
		totalOperatiopns := 0
		for j := 0; j < n; j++ {
			if boxes[j] == '1' {
				totalOperatiopns += int(math.Abs(float64(i - j)))
			}
		}
		answer[i] = totalOperatiopns
	}

	return answer
}
```


## Teste

```go
func main() {
	boxes := "110"
	result := minOperations(boxes)
	fmt.Println(result) // Output: [1, 1, 3]

	boxes2 := "001011"
	result2 := minOperations(boxes2)
	fmt.Println(result2) // Output: [11, 8, 5, 4, 3, 4]
}
```

