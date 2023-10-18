package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)
	fmt.Printf("x: %v\n", x)
	switch {
	case x < 100 && x > 0:
		println("between 0 and 100")
	case x < 200 && x > 101:
		println("between 101 and 200")
	case x > 201 && x < 250:
		println("between 201 and 250")
	}

	//exercise 20

}
