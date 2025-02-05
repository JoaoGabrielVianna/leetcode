# 1790. Check if One String Swap Can Make Strings Equal

**`Easy`**

You are given two strings s1 and s2 of equal length. A string swap is an operation where you choose two indices in a string (not necessarily different) and swap the characters at these indices.

Return true if it is possible to make both strings equal by performing at most one string swap on exactly one of the strings. Otherwise, return false.

### Example 1:

**Input:** s1 = "bank", s2 = "kanb"

**Output:** true

**Explanation:** 

For example, swap the first character with the last character of s2 to make "bank".

### xample 2:

**Input:** s1 = "attack", s2 = "defend"

**Output:** false

**Explanation:** 

It is impossible to make them equal with one string swap.

### Example 3:

**Input:** s1 = "kelb", s2 = "kelb"

**Output:** true

**Explanation:** 

The two strings are already equal, so no string swap operation is required.
 

### Constraints:

- `1 <= s1.length, s2.length <= 100`
- `s1.length == s2.length`
- `s1 and s2 consist of only lowercase English letters.`

### Solution

```go
func areAlmostEqual(s1 string, s2 string) bool {
	cnt := 0
	var c1, c2 byte
	for i := range s1 {
		a, b := s1[i], s2[i]
		if a != b {
			cnt++
			if cnt > 2 || (cnt == 2 && (a != c2 || b != c1)) {
				return false
			}
			c1, c2 = a, b
		}
	}
	return cnt != 1
}
```

### Test

```go
func main() {
	testCases := []struct {
		s1 string
		s2 string
	}{
		{s1: "bank", s2: "kanb"},
		{s1: "attack", s2: "defend"},
		{s1: "kelb", s2: "kelb"},
		{s1: "abcd", s2: "dcba"},
	}

	for _, testCase := range testCases {
		fmt.Println("INPUT:", testCase.s1, testCase.s2)
		fmt.Printf("OUTPUT: %v\n\n", areAlmostEqual(testCase.s1, testCase.s2))
	}
}

```