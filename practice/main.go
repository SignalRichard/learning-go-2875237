// Write your answer here, and then test your code.

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

// calculate() returns the result of the requested operation.
func calculate(input1 string, input2 string, operation string) float64 {
	// Your code goes here.
	result := 0.0

	f1 := convertInputToValue(input1)
	f2 := convertInputToValue(input2)

	switch operation {
	case "+":
		result = addValues(f1, f2)

	case "-":
		result = subtractValues(f1, f2)

	case "*":
		result = multiplyValues(f1, f2)

	case "/":
		result = divideValues(f1, f2)

	default:
		panic("Unknown operator")
	}

	return result
}

func convertInputToValue(input string) float64 {
	f1, err1 := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err1 != nil {
		fmt.Println(err1)

		panic("input must be a number")
	}

	return f1
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	return value1 / value2
}
