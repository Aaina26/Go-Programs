package main

import (
	"fmt"
	"sort"
)

type user struct {
	First string //exportable
	age   int    //non exportable
}

type u []user

func (a u) Len() int           { return len(a) }
func (a u) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a u) Less(i, j int) bool { return a[i].age < a[j].age }

func main() {
	sl := []user{
		{First: "Jay", age: 30},
		{First: "John", age: 15},
		{First: "Kay", age: 26},
	}
	sort.Sort(u(sl))
	fmt.Println(sl)

}
