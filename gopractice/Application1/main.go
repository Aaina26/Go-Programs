package main

//test
import (
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	First string //exportable
	age   int    //non exportable
}

type Users struct {
	first string
	age   int
}

func main() {
	sl := []user{
		{First: "Jay", age: 30},
		{First: "John", age: 15},
		{First: "Kay", age: 26},
	}

	sl1 := []Users{
		{first: "Jay", age: 30},
		{first: "John", age: 15},
		{first: "Kay", age: 26},
	}
	outpt, err := json.Marshal(sl)
	outpt1, err1 := json.Marshal(sl1)
	if err != nil {
		log.Fatal("Error:", err)
	}
	if err1 != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println(string(outpt))
	fmt.Println(string(outpt1))
}
