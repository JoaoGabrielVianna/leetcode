# 2270. Número de maneiras de dividir uma matriz
**`Médio`**

Você recebe uma matriz de inteiros indexados em 0`nums` de comprimento `n`.

`nums` contém uma divisão válida no índice `i` se o seguinte for verdadeiro:

- A soma dos primeiros `i + 1` elementos é maior ou igual à soma dos últimos `n - i - 1` elementos.
- Há pelo menos um elemento à direita de `i. Ou seja, `0 <= i < n - 1`.
Retorna o número de divisões válidas em `nums` .

 

### Exemplo 1:

**Entrada:** nums = [10,4,-8,7]

 **Saída:** 2

**Explicação:**  
Existem três maneiras de dividir nums em duas partes não vazias: 
- Dividir nums no índice 0. Então, a primeira parte é [10], e sua soma é 10. A segunda parte é [4,-8,7], e sua soma é 3. Como 10 >= 3, i = 0 é uma divisão válida. 
- Dividir nums no índice 1. Então, a primeira parte é [10,4], e sua soma é 14. A segunda parte é [-8,7], e sua soma é -1. Como 14 >= -1, i = 1 é uma divisão válida. 
- Dividir nums no índice 2. Então, a primeira parte é [10,4,-8], e sua soma é 6. A segunda parte é [7], e sua soma é 7. Como 6 < 7, i = 2 não é uma divisão válida. 
Portanto, o número de divisões válidas em nums é 2.

---

### Exemplo 2:

**Entrada:** nums = [2,3,1,0]

**Saída:** 2

**Explicação:**  
Existem duas divisões válidas em nums: 
- Dividir nums no índice 1. Então, a primeira parte é [2,3], e sua soma é 5. A segunda parte é [1,0], e sua soma é 1. Como 5 >= 1, i = 1 é uma divisão válida. 
- Dividir nums no índice 2. Então, a primeira parte é [2,3,1], e sua soma é 6. A segunda parte é [0], e sua soma é 0. Como 6 >= 0, i = 2 é uma divisão válida.
 

### Restrições:

- `2 <= nums.length <= 105`
- `-105 <= nums[i] <= 105`

---

## Solução

```go
func waysToSplitArray(nums []int) int {
	totalSum := 0

	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0
	numberOfValidSplit := 0

	// Percorrer o array até o penúltimo elemento (último elemento não pode ser início de uma divisão válida)
	for i := 0; i < len(nums)-1; i++ {
		leftSum += nums[i]

		rightSum := totalSum - leftSum

		if leftSum >= rightSum {
			numberOfValidSplit++
		}
	}

	return numberOfValidSplit
}
```

## Teste

```go
func main() {
	num1 := []int{10, 4, -8, 7} // -> Matriz
	num2 := []int{2, 3, 1, 0}   // -> Matriz

	fmt.Println(waysToSplitArray(num1))
	fmt.Println(waysToSplitArray(num2))

	waysToSplitArray(num1)

}
```