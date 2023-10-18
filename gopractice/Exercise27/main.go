package main

import "fmt"

func main() {
	i := 100
	for i > 0 {
		if i == 75 {
			break
		}
		fmt.Println(i)
		i--
	}
}
