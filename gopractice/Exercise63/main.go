package main

import "fmt"

// Add func adds two integer values.
func Add(a int, b int) int {
	return a + 2*b
}

func main() {
	x := Add(10, 10)
	fmt.Println(x)
}
