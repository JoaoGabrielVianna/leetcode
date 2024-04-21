# 2635. Apply Transform Over Each Element in Array

Given an integer array `arr` and a mapping function `fn`, return a new array with a transformation applied to each element.

- `Easy`

The returned array should be created such that `returnedArray[i] = fn(arr[i], i)`.

Please solve it without the built-in `Array.map` method.

 

### Example 1:

> **Input:** arr = [1,2,3], fn = function plusone(n) { return n + 1; }<br/>
> **Output:** [2,3,4]<br/>
> **Explanation:**<br/>
> const newArray = map(arr, plusone); // [2,3,4]<br/>
> The function increases each value in the array by one. <br/>

### Example 2:

> **Input:** arr = [1,2,3], fn = function plusI(n, i) { return n + i; }<br/>
> **Output:** [1,3,5]<br/>
> **Explanation:** The function increases each value by the index it resides in.<br/>

### Example 3:

> **Input:** arr = [10,20,30], fn = function constant() { return 42; }<br/>
> **Output:** [42,42,42]<br/>
> **Explanation:** The function always returns 42.<br/>

 

Constraints:

- `0 <= arr.length <= 1000`
- `-10^9 <= arr[i] <= 10^9`
- `fn` returns a number

#

# My Solution:
```js
// Author: João Gabriel Vianna
// Date: 16/04/2024
// source: https://leetcode.com/problems/

/**
 * Applies a given function to each element of an array.
 * 
 * @param {number[]} arr - The array of numbers to be processed.
 * @param {Function} fn - The function to apply to each element of the array.
 */

function map(arr, fn) {
    let resultArray = [];

    // Iterating over each element of the input array
    arr.forEach((item, i) => {
        // Applying the given function to the current element and adding the result to the result array
        resultArray.push(fn(item, i));
    });

    return resultArray;
}

/**
 * Increments a number by 1.
 * 
 * @param {number} n - The number to be incremented.
 * @returns {number} - The incremented number.
 */
function plusone(n) { return n + 1; }

/**
 * Increments a number by its index.
 * 
 * @param {number} n - The number to be incremented.
 * @param {number} i - The index of the number in the array.
 * @returns {number} - The incremented number.
 */
function plusl(n, i) { return n + i; }

/**
 * Returns a constant value of 42.
 * 
 * @returns {number} - The constant value 42.
 */
function constant() { return 42; }
```