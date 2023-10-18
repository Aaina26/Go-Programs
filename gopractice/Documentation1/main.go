// Package Documentation
package main

import "fmt"

// Years function is used to convert dog year to human years
func Years(dy int) int {
	return dy * 7
}
func main() {
	hy := Years(10)
	fmt.Println(hy)
}
