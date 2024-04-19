// Author: João Gabriel Vianna
// Date: 16/04/23
// source: https://leetcode.com/problems/counter-ii


/**
 * This function have 3 methods,
 * 
 * `increment() // n + 1`
 * 
 * `decrement() // n - 1`
 * 
 * `reset() // n`
 * @param {int} init 
 * @returns 
 */
var createCounter = function (init) {
    let count = init;

    function increment() {
        return ++count
    }
    function decrement() {
        return --count
    }
    function reset() {
        return (count = init)
    }

    return { increment, decrement, reset };
}

