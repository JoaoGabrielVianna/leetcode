# 1079. Letter Tile Possibilities

**`Medium`**

You have n  tiles, where each tile has one letter tiles[i] printed on it.

Return the number of possible non-empty sequences of letters you can make using the letters printed on those tiles.

### Example 1:

**Input:** tiles = "AAB"

**Output:** 8

**Explanation:** 

The possible sequences are "A", "B", "AA", "AB", "BA", "AAB", "ABA", "BAA".


### Example 2:

**Input:** tiles = "AAABBC"

**Output:** 188

### Example 3:

**Input:** tiles = "V"

**Output:** 1

 

### Constraints:

- `1 <= tiles.length <= 7`
- `tiles consists of uppercase English letters.`

### Solution


```go
func numTilePossibilities(tiles string) int {
	cnt := [26]int{}
	for _, c := range tiles {
		cnt[c-'A']++
	}
	var dfs func(cnt [26]int) int
	dfs = func(cnt [26]int) (res int) {
		for i, x := range cnt {
			if x > 0 {
				res++
				cnt[i]--
				res += dfs(cnt)
				cnt[i]++
			}
		}
		return
	}
	return dfs(cnt)
}
```


### Test

```go
func main() {
	testCases := []struct {
		tiles string
	}{
		{tiles: "AAB"},
		{tiles: "AAABBC"},
	}
	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.tiles)
		fmt.Printf("OUTPUT: %v\n\n", numTilePossibilities(testCase.tiles))
	}
}

```