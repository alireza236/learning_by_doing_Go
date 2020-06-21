package main

import "fmt"

func main() {
	var positiveNumber = 100
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif: %d\n", positiveNumber)
	fmt.Printf("bilangan negatif: %d\n", negativeNumber)

	var decimalNumber float64 = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	var exist = true
	fmt.Printf("exist? %t \n", exist)

	var message = `Nama saya "John Wick".
Salam kenal.
Mari belajar "Golang".`

	fmt.Println(message)
}
