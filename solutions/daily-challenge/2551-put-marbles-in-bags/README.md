# 2551. Put Marbles in Bags

**`Hard`**


You have k bags. You are given a 0-indexed integer array weights where weights[i] is the weight of the ith marble. You are also given the integer k.

Divide the marbles into the k bags according to the following rules:

    No bag is empty.
    If the ith marble and jth marble are in a bag, then all marbles with an index between the ith and jth indices should also be in that same bag.
    If a bag consists of all the marbles with an index from i to j inclusively, then the cost of the bag is weights[i] + weights[j].

The score after distributing the marbles is the sum of the costs of all the k bags.

Return the difference between the maximum and minimum scores among marble distributions.

 

### Example 1:

**Input:** weights = [1,3,5,1], k = 2
**Output:** 4
**Explanation:**
The distribution [1],[3,5,1] results in the minimal score of (1+1) + (3+1) = 6. 
The distribution [1,3],[5,1], results in the maximal score of (1+3) + (5+1) = 10. 
Thus, we return their difference 10 - 6 = 4.

### Example 2:

**Input:** weights = [1, 3], k = 2
**Output:** 0
Explanation: The only distribution possible is [1],[3]. 
Since both the maximal and minimal score are the same, we return 0.

 

**Constraints:**

- `1 <= k <= weights.length <= 105`
- `1 <= weights[i] <= 109`

## Solution
```go
func putMarbles(weights []int, k int) (ans int64) {
	n := len(weights)
	arr := make([]int, n-1)
	for i, w := range weights[:n-1] {
		arr[i] = w + weights[i+1]
	}
	sort.Ints(arr)
	for i := 0; i < k-1; i++ {
		ans += int64(arr[n-2-i] - arr[i])
	}
	return
}
```

## Test
```go
func main() {
	testCases := []struct {
		weights []int
		k       int
		output  int64
	}{
		{weights: []int{1, 3, 5, 1}, k: 2, output: 4},
		{weights: []int{1, 3}, k: 1, output: 0},
		{weights: []int{10, 20, 30, 40, 50}, k: 3, output: 60},
		{weights: []int{5, 5, 5, 5}, k: 2, output: 0},
		{weights: []int{1, 2, 3, 4, 5}, k: 4, output: 6},
	}
	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.weights, testCase.k)
		fmt.Printf("OUTPUT: %v\n", putMarbles(testCase.weights, testCase.k))
		fmt.Printf("EXPECTED: %v\n\n", testCase.output)
	}
}
```