package main

import (
	"fmt"
	"strings"
)

func main() {

	/* 	var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)

	   	var msg = fmt.Sprintf("rata-rata %.2f", avg)

	   	fmt.Println(msg) */

	/* 	var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}

		   	var avg = calculate(numbers...)

		   	var msg = fmt.Sprintf("Rata- rata %.2f", avg)

	         fmt.Println(msg) */

	var hobbies = []string{"makan", "tidur", "coding"}

	yourHobbies("reza", hobbies...)
}

func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")
	fmt.Printf("Hello, my name is: %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

func calculate(numbers ...int) float64 {
	var total int = 0

	for _, val := range numbers {
		total += val
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}
