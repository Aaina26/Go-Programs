package main

import "fmt"

func main() {
	x := 10
	fmt.Printf("%v\t%v\t%T\n", x, &x, &x)
	//Exercise75
	y := "light"
	z := 30

	v1, v2 := &y, &z
	fmt.Println(v1, v2)
	fmt.Printf("%T\t%T\n", v1, v2)
	fmt.Println(*v1, "\t", *v2)
}
