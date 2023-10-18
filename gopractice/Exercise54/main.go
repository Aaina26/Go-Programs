package main

import "fmt"

type person struct {
	first string
	last  string
	iceCF []string
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		iceCF: []string{"chocolate", "vanilla"},
	}
	p2 := person{
		first: "Tia",
		last:  "L",
		iceCF: []string{"blueberry", "blackforest"},
	}

	mp := map[string]person{
		"Bond": p1,
		"L":    p2,
	}

	for k, v := range mp {
		fmt.Println(k, v.first, v.last)
		for i, val := range v.iceCF {
			fmt.Println(i, val)
		}
	}

}
