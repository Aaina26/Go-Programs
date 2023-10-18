package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func cal(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func main() {
	x := cal(10, 20, add)
	y := cal(20, 10, subtract)
	fmt.Println(x, y)
}
