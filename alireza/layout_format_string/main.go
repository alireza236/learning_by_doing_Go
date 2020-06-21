package main

import "fmt"

type student struct {
	name        string
	height      float64
	age         int64
	isGraduated bool
	hobbies     []string
}

var data = student{
	name:        "wick",
	height:      182.5,
	age:         26,
	isGraduated: false,
	hobbies:     []string{"eating", "sleeping"},
}

func main() {
	fmt.Printf("%d\n", data.age)

	fmt.Printf("%e\n", data.height)
	// 1.825000e+02

	fmt.Printf("%E\n", data.height)

}
