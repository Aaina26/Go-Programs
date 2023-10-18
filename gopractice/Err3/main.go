package main

import (
	"fmt"
	"log"
)

type customErr struct {
	err error
}

func (c customErr) Error() string {
	return fmt.Sprintf("Error occured: %v\n", c.err)
}

func main() {
	c := customErr{err: fmt.Errorf("Zero Division Error")} //c := customErr{}->this will work too
	foo(c)
}

func foo(e customErr) {
	log.Fatalln(e)
}
