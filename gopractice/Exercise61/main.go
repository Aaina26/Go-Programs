package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, "\nAnd my age is", p.age)
}

func main() {
	p1 := person{
		first: "Aaina",
		age:   22,
	}
	p1.speak()
}
