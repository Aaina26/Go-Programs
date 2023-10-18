package main

import "fmt"

func main() {
	x := [][]string{{"James", "Bond", "Shaken, not stirred"}, {"Miss", "Moneypenny", "I'm 008."}}

	for _, v := range x {
		for _, val := range v {
			fmt.Printf("%s\t", val)
		}
		fmt.Println()
	}
}
