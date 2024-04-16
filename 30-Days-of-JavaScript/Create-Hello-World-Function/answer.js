// Author: João Gabriel Vianna
// Date: 16/04/23
// source: https://leetcode.com/problems/create-hello-world-function/


/*** This Function allways return "Hello World"
 * @returns `"Hello World"`
 */
var createHelloWorld = function () {
    return function () {
        return "Hello World"
    }
}

const answer = function () {
    console.log(createHelloWorld()());
}

answer(); // "Hello World"