package main

import "fmt"

func main() {
	x := func(a int, b int) int { return a + b }(4, 5)
	fmt.Println(x)

	//Exercise 69
	yfunc := func(a int, b int) int {
		return a - b
	}
	fmt.Println(yfunc(10, 4))
}
