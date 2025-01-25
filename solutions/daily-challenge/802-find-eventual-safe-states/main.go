package main

import (
	"fmt"
	"sort"
)

func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	reverseGraph := make([][]int, n)
	inDegree := make([]int, n)

	// Construir o grafo reverso e calcular o grau de entrada
	for node, neighbors := range graph {
		for _, neighbor := range neighbors {
			reverseGraph[neighbor] = append(reverseGraph[neighbor], node)
		}
		inDegree[node] = len(neighbors)
	}

	// Inicializar a fila com nós terminais (grau de entrada 0)
	queue := []int{}
	for i := 0; i < n; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	safeNodes := []int{}

	// Processar a fila
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		safeNodes = append(safeNodes, node)

		// Reduzir o grau de entrada dos vizinhos no grafo reverso
		for _, neighbor := range reverseGraph[node] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// Ordenar os nós seguros em ordem crescente
	sort.Ints(safeNodes)

	return safeNodes
}

func main() {
	testCases := []struct {
		graph [][]int
	}{
		{graph: [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}},
		{graph: [][]int{{1, 2, 3, 4}, {1, 2}, {3, 4}, {0, 4}, {}}},
	}

	for _, testCase := range testCases {
		result := eventualSafeNodes(testCase.graph)
		fmt.Println("INPUT:", testCase.graph)
		fmt.Printf("OUTPUT: %v \n\n", result)
	}
}
