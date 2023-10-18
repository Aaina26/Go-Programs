package main

import (
	"fmt"
	"sync"
)

var count int = 0
var wg sync.WaitGroup

func main() {
	wg.Add(2)
	c1, c2 := make(chan int), make(chan int)
	go send(c1)
	FanInOut(c1, c2)
	go receive(c2)
	wg.Wait()
	fmt.Println("Program execution complete...")
}

func send(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	wg.Done()
	close(c)
}

func FanInOut(c1 <-chan int, c2 chan int) <-chan int {
	for i := 0; i < 10; i++ {
		go func(i1 int) {
			for v := range c1 {
				c2 <- task(v)
			}

		}(i)

	}

	return c2
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
	wg.Done()
}

func task(v int) int {
	count++
	return v + count
}
