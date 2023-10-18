package main

import "fmt"

func Sum(a ...int) (s int) {
	for _, v := range a {
		s = s + v
	}
	return s
}

func main() {
	a := []int{-1, 5, 6, 7, 9, -10}
	sum := Sum(a...)
	fmt.Println(sum)
}
