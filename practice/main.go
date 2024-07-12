package main

import (
	"fmt"
	"sort"
)

func main() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = colors[1:]
	fmt.Println(colors)

	colors = colors[:len(colors)-1]
	fmt.Println(colors)

	numbers := make([]int, 5) //make([]int, 5, 5)
	numbers[0] = 134
	numbers[1] = 72
	numbers[2] = 32
	numbers[3] = 12
	numbers[4] = 156
	fmt.Println(numbers)

	numbers = append(numbers, 265)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)
}
