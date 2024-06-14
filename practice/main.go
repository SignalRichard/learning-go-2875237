// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

// calculate() returns the sum of the two parameters
func calculate(value1 string, value2 string) float64 {
	// Your code goes here.

	// Convert the first string to a float64
	f1, err1 := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err1 != nil {
		fmt.Println(err1)
		panic("Value 1 must be a number")
	}
	// Convert the second string to a float64
	f2, err2 := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err2 != nil {
		fmt.Println(err2)
		panic("Value 2 must be a number")
	}
	// Calculate and return the result

	return f1 + f2
}
