package main

import (
	"fmt"
	"log"
	"os"
)

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func main() {
	f, r := os.Create("log.txt")
	if r != nil {
		log.Fatalln(r)
	}
	defer f.Close()
	log.SetOutput(f)
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		se := sqrtError{"55", "66", fmt.Errorf("Unfortunate Error..Please correct it.")}
		log.Println(se)
	}
	return 42, nil
}
