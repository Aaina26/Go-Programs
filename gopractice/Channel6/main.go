package main

import "fmt"

func main() {
	c := make(chan int)

	go send(c)
	receive(c) //in this case wait condition is not required because main can't end until execution of receive is completed-> which further gets completed along with send
	//if both function hv go prefix..then receive will hv separate go routine than main and main will be over before both routines execute completely->will result in no outputs
}

func send(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func receive(c <-chan int) {
	for v := range c {
		fmt.Println(v)
	}
}
