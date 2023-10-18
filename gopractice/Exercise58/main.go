package main

import "fmt"

func foo() int {
	return 10
}

func bar() (s string, i int) {
	s = "Aaina"
	i = 22
	return
}

func main() {
	defer fmt.Println(foo()) //Defer for Exercise 60
	s, i := bar()
	fmt.Println(s, i)
}
