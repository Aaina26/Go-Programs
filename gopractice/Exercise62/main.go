package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func main() {
	sq := square{
		length: 10,
		width:  22.5,
	}

	cir := circle{
		radius: 7,
	}
	info(cir)
	info(sq)
}

func (c circle) area() float64 {
	a := math.Pi * math.Pow(c.radius, 2)
	return a
}

func (s square) area() float64 {
	return s.length * s.width
}

type shape interface {
	area() float64
}

func info(sh shape) {
	fmt.Println(sh.area())
}
