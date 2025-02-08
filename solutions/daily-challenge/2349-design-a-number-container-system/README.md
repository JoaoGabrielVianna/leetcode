# 2349. Design a Number Container System

**`Medium`**

Design a number container system that can do the following:

- Insert or Replace a number at the given index in the system.
- Return the smallest index for the given number in the system.
Implement the NumberContainers class:

- NumberContainers() Initializes the number container system.
-void change(int index, int number) Fills the container at index with the number. If there is already a number at that index, replace it.
- int find(int number) Returns the smallest index for the given number, or -1 if there is no index that is filled by number in the system.
 

### Example 1:

**Input:**
["NumberContainers", "find", "change", "change", "change", "change", "find", "change", "find"]
[[], [10], [2, 10], [1, 10], [3, 10], [5, 10], [10], [1, 20], [10]]
**Output:**
[null, -1, null, null, null, null, 1, null, 2]

**Explanation:**

NumberContainers nc = new NumberContainers();
nc.find(10); // There is no index that is filled with number 10. Therefore, we return -1.
nc.change(2, 10); // Your container at index 2 will be filled with number 10.
nc.change(1, 10); // Your container at index 1 will be filled with number 10.
nc.change(3, 10); // Your container at index 3 will be filled with number 10.
nc.change(5, 10); // Your container at index 5 will be filled with number 10.
nc.find(10); // Number 10 is at the indices 1, 2, 3, and 5. Since the smallest index that is filled with 10 is 1, we return 1.
nc.change(1, 20); // Your container at index 1 will be filled with number 20. Note that index 1 was filled with 10 and then replaced with 20. 
nc.find(10); // Number 10 is at the indices 2, 3, and 5. The smallest index that is filled with 10 is 2. Therefore, we return 2.
 

### Constraints:

- `1 <= index, number <= 109`
- `At most 105 calls will be made in total to change and find.`

### Solution

```go
type NumberContainers struct {
	d map[int]int
	g map[int]*redblacktree.Tree
}

func Constructor() NumberContainers {
	return NumberContainers{map[int]int{}, map[int]*redblacktree.Tree{}}
}

func (this *NumberContainers) Change(index int, number int) {
	if oldNumber, ok := this.d[index]; ok {
		this.g[oldNumber].Remove(index)
	}
	this.d[index] = number
	if _, ok := this.g[number]; !ok {
		this.g[number] = redblacktree.NewWithIntComparator()
	}
	this.g[number].Put(index, nil)
}

func (this *NumberContainers) Find(number int) int {
	if ids, ok := this.g[number]; ok && ids.Size() > 0 {
		return ids.Left().Key.(int)
	}
	return -1
}

/**
 * Your NumberContainers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Change(index,number);
 * param_2 := obj.Find(number);
 */
```

### Test

```go
func main() {
	testCases := []struct {
		operations []string
		inputs     [][]int
		expected   []interface{}
	}{
		{
			operations: []string{"NumberContainers", "find", "change", "change", "change", "change", "find", "change", "find"},
			inputs:     [][]int{{}, {10}, {2, 10}, {1, 10}, {3, 10}, {5, 10}, {10}, {1, 20}, {10}},
			expected:   []interface{}{nil, -1, nil, nil, nil, nil, 1, nil, 2},
		},
	}

	for i, tc := range testCases {
		fmt.Printf("Test Case %d:\n", i+1)
		var obj NumberContainers

		for j, op := range tc.operations {
			switch op {
			case "NumberContainers":
				obj = Constructor()
				fmt.Println("Constructor() -> nil")
			case "change":
				index, number := tc.inputs[j][0], tc.inputs[j][1]
				obj.Change(index, number)
				fmt.Printf("Change(%d, %d) -> nil\n", index, number)
			case "find":
				number := tc.inputs[j][0]
				result := obj.Find(number)
				fmt.Printf("Find(%d) -> %d\n", number, result)
			}
		}
		fmt.Println()
	}
}

```