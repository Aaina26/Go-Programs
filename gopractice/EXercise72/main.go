package main

import "fmt"

func decrement() func() int {
	i := 100
	c := 0
	return func() int {
		i = i - c
		c++
		return i
	}
}

func main() {
	f := decrement()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
