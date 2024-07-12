// Write your answer here, and then test your code.
// Your job is to implement the getCartFromJson() method.

package main

import (
	"encoding/json"
	"strings"
	//"golang.org/x/exp/constraints"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
const showExpectedResult = false
const showHints = false

type cartItem struct {
	Name     string
	Price    float64
	Quantity int
}

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []cartItem {
	var cart []cartItem
	// Your code goes here.

	cart = convertFromJson[cartItem](jsonString)

	return cart
}

func convertFromJson[T any](content string) []T {
	items := make([]T, 0, 3)

	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	if err != nil {
		panic(err)
	}
	var item T

	for decoder.More() {
		err := decoder.Decode(&item)
		if err != nil {
			panic(err)
		}

		items = append(items, item)
	}

	return items
}
