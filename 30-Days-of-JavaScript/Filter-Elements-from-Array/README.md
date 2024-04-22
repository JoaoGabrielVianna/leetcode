# 2634. Filter Elements from Array


Given an integer array `arr` and a filtering function `fn`, return a filtered array `filteredArr`.
- `Easy`


The `fn` function takes one or two arguments:

- `arr[i]`- number from the `arr`
- `i` - index of `arr[i]`

`filteredArr` should only contain the elements from the `arr` for which the expression `fn(arr[i], i)` evaluates to a truthy value. A truthy value is a value where `Boolean(value)` returns `true`.

Please solve it without the built-in `Array.filter` method.

 

### Example 1:

> **Input:** arr = [0,10,20,30], fn = function greaterThan10(n) { return n > 10; }<br/>
> **Output:** [20,30]<br/>
> **Explanation:**<br/>
> const newArray = filter(arr, fn); // [20, 30]<br/>
> The function filters out values that are not greater than 10<br/>

### Example 2:

> Input: arr = [1,2,3], fn = function firstIndex(n, i) { return i === 0; }
> Output: [1]
> Explanation:
> fn can also accept the index of each element
> In this case, the function removes elements not at index 0

### Example 3:

> **Input:** arr = [-2,-1,0,1,2], fn = function plusOne(n) { return n + 1 }<br/>
> **Output:** [-2,0,1,2]<br/>
> **Explanation:**<br/>
> Falsey values such as 0 should be filtered out<br/>

 

### Constraints:

- `0 <= arr.length <= 1000`
- `-109 <= arr[i] <= 109`

# My Solution:
```js
/**
 * Filters the elements of an array based on a given function.
 * 
 * @param {Number[]} arr - The input array to be filtered.
 * @param {Function} fn - The filtering function. It should accept two arguments: the current element and its index, and return a boolean indicating whether to keep the element or not.
 * @returns {Number[]} - The filtered array containing elements that satisfy the filtering function.
 */
function filter(arr, fn) {
    let newArr = [];
    arr.forEach((item, i) => {
        if (fn(item, i)) {
            newArr.push(item);
        }
    });
    return newArr;
}

/**
 * Checks if a number is greater than 10.
 * 
 * @param {Number} n - The number to be checked.
 * @returns {Boolean} - True if the number is greater than 10, otherwise false.
 */
function greaterThan10(n) {
    return n > 10;
}

/**
 * Checks if the index is the first index in the array.
 * 
 * @param {Number} n - The current element.
 * @param {Number} i - The index of the current element.
 * @returns {Boolean} - True if the index is 0, otherwise false.
 */
function firstIndex(n, i) {
    return i === 0;
}

/**
 * Adds one to a number.
 * 
 * @param {Number} n - The number to which one will be added.
 * @returns {Number} - The input number plus one.
 */
function plusOne(n) {
    return n + 1;
}

// Test Output
const arr1 = [0, 10, 20, 30];
const arr2 = [1, 2, 3];
const arr3 = [-2, -1, 0, 1, 2];
const newArray1 = filter(arr1, greaterThan10);
const newArray2 = filter(arr2, firstIndex);
const newArray3 = filter(arr3, plusOne);

```