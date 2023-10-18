package main

import "fmt"

func main() {
	ma := map[string][]string{
		"bond_james":       {`shaken, not stirred`, `martinis`, `fast cars`},
		"moneypenny_jenny": {`intelligence`, `literature`, `computer science`},
		"no_dr":            {`cats`, `ice cream`, `sunsets`},
	}
	for k, v := range ma {
		fmt.Printf("%s->", k)
		for i, val := range v {
			fmt.Printf("%d-%s\t", i, val)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()

	//Exercise 50

	ma[`fleming_ian`] = []string{`steaks`, `cigars`, `espionage`} //adding element
	for k, v := range ma {
		fmt.Printf("%s->", k)
		for i, val := range v {
			fmt.Printf("%d-%s\t", i, val)
		}
		fmt.Println()
	}
	fmt.Println()
	fmt.Println()

	//Exercise 51
	delete(ma, "no_dr") //deletion
	for k, v := range ma {
		fmt.Printf("%s->", k)
		for i, val := range v {
			fmt.Printf("%d-%s\t", i, val)
		}
		fmt.Println()
	}
}
