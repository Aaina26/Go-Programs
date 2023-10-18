package main

import "fmt"

type dog struct {
	name string
}

func (d dog) walk() {
	fmt.Println("I am ", d.name, "I am going to take a walk.")
}

func (d *dog) run() {
	d.name = "Joy"
	fmt.Println("I am", d.name, "I am going to run!")
}

type refVal interface {
	walk()
	run()
}

func do(r refVal) {
	r.walk()
}

func main() {
	d1 := dog{
		name: "Tommy",
	}
	d2 := &dog{
		name: "Jerry",
	}

	d1.walk()
	d1.run()
	d2.walk()
	d2.run()

	do(&d1) //do(d1) will give error
	do(d2)

}
