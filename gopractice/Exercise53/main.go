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

	fmt.Printf("p1 is: %#v\n", p1)
	fmt.Printf("p1 is: %#v\n", p2)

	for _, str := range p2.iceCF {
		fmt.Println(str)
	}

}
