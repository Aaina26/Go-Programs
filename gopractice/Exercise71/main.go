package main

import "fmt"

func square(n int) int {
	return n * n
}

func printSquare(f func(int) int, n int) string {
	s := fmt.Sprint(f(n))
	return s
}

func main() {
	fmt.Println(printSquare(square, 8))
}
