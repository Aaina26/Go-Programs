package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"A", "C", "B", "b"}
	i := []int{23, 8, 9, 10}

	sort.Ints(i)
	sort.Strings(s)

	fmt.Println(s)
	fmt.Println(i)
}
