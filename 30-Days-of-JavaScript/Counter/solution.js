// Author: João Gabriel Vianna
// Date: 16/04/2024
// source: https://leetcode.com/problems/counter/?envType=study-plan-v2&envId=30-days-of-javascript


/**
 * This function return the current value of the counter `num`.
 * 
 *  When called again, teh returns value will be incremented by 1 `num + 1`
 * 
 * Each subsequent call to the function will returns the incremented `value`.
 * 
 * `@requires number`
 * 
 * `@returns {number}` The current value of the counter `num`
 */
var createCounter = function(n) {
    return function() {
        return n++;
    };
};


var n = 10;
var counter = createCounter(n);

console.log(counter()); // Returns: num 
console.log(counter()); // Returns: num + 1
console.log(counter()); // Returns: num + 2