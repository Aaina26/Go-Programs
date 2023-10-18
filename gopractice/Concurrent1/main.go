package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func foo() {
	time.Sleep(time.Second * 3)
	fmt.Println("foo")
	wg.Done()
}

func bar() {
	time.Sleep(time.Second * 3)
	fmt.Println("bar")
	wg.Done()
}

func hello() {
	fmt.Println("hello")
}

var wg sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU())
	hello()
	wg.Add(2)
	go bar()
	go foo() //both functions might not print ...if the program exits before the routines are complete...therefore, waitGroup  is important
	//wg.Wait()->number of go routine printed will be 1 as the print statement only runs after the 2 extra go routines have completed
	fmt.Println(runtime.NumGoroutine())
	wg.Wait()

}
