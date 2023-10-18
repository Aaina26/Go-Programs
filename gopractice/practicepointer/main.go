package main

import "fmt"

func byVal(a int) {
	a = 10
}

func byRef(b *int) {
	*b = 15
}

func main() {
	x := 20
	fmt.Println(x)
	byVal(x)
	fmt.Println(x)
	byRef(&x)
	fmt.Println(x)
}
