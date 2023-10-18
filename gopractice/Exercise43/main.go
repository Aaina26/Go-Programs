package main

import "fmt"

func main() {
	sl := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for i, v := range sl {
		fmt.Printf("%d-%T-%d\n", i, v, v)
	}

	//exercise 44
	sl1 := sl[0:5]
	sl2 := sl[5:]
	sl3 := sl[2:7]
	sl4 := sl[1:6]

	fmt.Println(sl1)
	fmt.Println(sl2)
	fmt.Println(sl3)
	fmt.Println(sl4)

	//Exercise 45
	sl = append(sl, 52)
	fmt.Println(sl)
	sl = append(sl, 53, 54, 55)
	fmt.Println(sl)
	y := []int{56, 57, 58, 59, 60}
	sl = append(sl, y...) //important
	fmt.Println(sl)
}
