package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	go func() {
		q <- 32
	}()

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c, q <-chan int) {
	count := 0
	for count < 5 {
		select {
		case v := <-c:
			fmt.Println("from c", c, v)
			count++
		case v := <-q:
			fmt.Println("from q", q, v)
			count++
		}

	}
}

func gen(q <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}
