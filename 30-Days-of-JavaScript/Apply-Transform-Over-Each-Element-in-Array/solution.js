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

