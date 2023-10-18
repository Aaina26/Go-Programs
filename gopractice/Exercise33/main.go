package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("Iteration:%d-->", i+1)
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("variable is 3")
		}
		fmt.Printf("\n")
	}
}
