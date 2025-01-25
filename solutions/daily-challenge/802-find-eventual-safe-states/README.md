# 802. Find Eventual Safe States

**`Medium`**

There is a directed graph of n nodes with each node labeled from 0 to n - 1. The graph is represented by a 0-indexed 2D integer array graph where graph[i] is an integer array of nodes adjacent to node i, meaning there is an edge from node i to each node in graph[i].

A node is a terminal node if there are no outgoing edges. A node is a safe node if every possible path starting from that node leads to a terminal node (or another safe node).

Return an array containing all the safe nodes of the graph. The answer should be sorted in ascending order.

 

### Example 1:
![Example 1](/assets/picture1.png)

Illustration of graph
**Input:** graph = [[1,2],[2,3],[5],[0],[5],[],[]]

**Output:** [2,4,5,6]

**Explanation:** 

The given graph is shown above.
Nodes 5 and 6 are terminal nodes as there are no outgoing edges from either of them.
Every path starting at nodes 2, 4, 5, and 6 all lead to either node 5 or 6.

### Example 2:

**Input:** graph = [[1,2,3,4],[1,2],[3,4],[0,4],[]]

**Output:** [4]

**Explanation:**

Only node 4 is a terminal node, and every path starting at node 4 leads to node 4.
 

### Constraints:

- `n == graph.length`
- `1 <= n <= 104`
- `0 <= graph[i].length <= n`
- `0 <= graph[i][j] <= n - 1`
- `graph[i] is sorted in a strictly increasing order.`
- `The graph may contain self-loops.`
- `The number of edges in the graph will be in the range [1, 4 * 104].`

### Solutions
```go
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
```

### Test

```go
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

```