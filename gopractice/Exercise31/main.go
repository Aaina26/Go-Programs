package main

import "fmt"

func main() {
	m := map[string]int{
		"Aaina": 42,
		"Kia":   43,
		"Leah":  44,
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	//exercise 32
	//when element is in the map
	if x, ok := m["Aaina"]; ok {
		fmt.Println("Value is:", x)
	}
	//when element not in map
	if x, ok := m["bla"]; !ok {
		fmt.Println("Value is:", x)
	}
}
