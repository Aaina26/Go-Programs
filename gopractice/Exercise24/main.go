package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		println(i)
	}
	j := 0
	for j < 100 {
		x := rand.Intn(10)
		y := rand.Intn(10)
		fmt.Println(x)
		fmt.Println(y)
		if x < 4 && y < 4 {
			fmt.Println("both less than 4")
		} else if x > 6 && y > 6 {
			fmt.Println("Both are greater than 6")
		} else if x >= 4 && x <= 6 {
			fmt.Println("x between 4 and 6")
		} else if y != 5 {
			fmt.Println("y not equal to 5")
		} else {
			fmt.Println("not matching criteria")
		}
		j++
	}
}
