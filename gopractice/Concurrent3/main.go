package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	counter := 0

	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			v++
			runtime.Gosched()
			counter = v
			wg.Done()
		}()
	}
	fmt.Println(counter)
	wg.Wait()
}
