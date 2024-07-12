package main

import "fmt"

func main() {
	val := 18.43
	var ptr = &val
	fmt.Printf("%.2f\n", *ptr)
}
