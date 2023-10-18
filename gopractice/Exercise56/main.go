package main

import "fmt"

func main() {
	as := struct {
		first    string
		friends  map[string]int
		favDrink []string
	}{
		first: "James",
		friends: map[string]int{
			"Robin":   30,
			"Natasha": 40,
			"Steve":   50,
		},
		favDrink: []string{"vodka", "lemongrass", "red wine"},
	}
	fmt.Println(as.first)
	fmt.Println("Friends are: ")
	for k, v := range as.friends {
		fmt.Println(k, v)
	}
	fmt.Println("Favourite drinks are:")
	for i, v := range as.favDrink {
		fmt.Println(i, v)
	}
}
