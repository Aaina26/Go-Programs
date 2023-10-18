package main

import "fmt"

type engine struct {
	electric bool
}

type vehicle struct {
	e     engine //we can just write 'engine' in this statement, name and type of the field will be equal to engine
	make  bool
	model string
	doors int16
	color string
}

func main() {
	v1 := vehicle{
		e:     engine{electric: true},
		make:  true,
		model: "car",
		doors: 4,
		color: "blue",
	}
	v2 := vehicle{
		e:     engine{electric: false},
		make:  true,
		model: "scooter",
		doors: 0,
		color: "red",
	}

	fmt.Println(v1)
	fmt.Println(v1.e.electric)
	fmt.Println(v1.color)
	fmt.Println(v2)
	fmt.Println(v2.e.electric)
	fmt.Println(v2.color)
}
