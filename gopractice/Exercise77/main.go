package main

import "fmt"

type person struct {
	first string
}

func changeVal(p person) string {
	p.first = "James"
	return p.first
}

func changeRef(p *person) {
	p.first = "Tom"
}

func main() {
	p1 := person{
		first: "Jay",
	}
	fmt.Println(p1)
	p1.first = changeVal(p1)
	fmt.Println(p1)
	changeRef(&p1)
	fmt.Println(p1)
}
