package main

import (
	"encoding/json"
	"os"
)

type user struct {
	First string //exportable
	age   int    //non exportable
}

func main() {
	sl := []user{
		{First: "Jay", age: 30},
		{First: "John", age: 15},
		{First: "Kay", age: 26},
	}

	json.NewEncoder(os.Stdout).Encode(sl)
}
