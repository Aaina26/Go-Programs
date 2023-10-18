package main

import "fmt"

func Int42() int {
	return 42
}

func Funcret(f func() int) func() int {
	return f //another possible way is to not take function as an argument and directly return anonymous func
}

func main() {
	f := Funcret(Int42)
	x := f()
	fmt.Println(x)
}
