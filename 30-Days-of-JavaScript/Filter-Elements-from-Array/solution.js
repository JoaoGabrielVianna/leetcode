// Author: João Gabriel Vianna
// Date: 16/04/23
// source: https://leetcode.com/problems/filter-elements-from-array

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
