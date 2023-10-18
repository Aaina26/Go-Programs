package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 42; i++ {
		fmt.Println("Iteration", i+1)
		x := rand.Intn(5)
		fmt.Printf("Variable is int of val %d\n", x)
	}
}
