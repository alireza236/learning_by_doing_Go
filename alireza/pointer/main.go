package main

import "fmt"

func main() {
	/* var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA:(value)", numberA)
	fmt.Println("numberA:(addres)", &numberA)

	fmt.Println("numberB:(value)", *numberB)
	fmt.Println("numberB:(addres)", &numberB)

	fmt.Println("===================")

	numberA = 5

	fmt.Println("numberA (value)", numberA)
	fmt.Println("numberA (address)", &numberA)
	fmt.Println("numberB (value)", *numberB)
	fmt.Println("numberB (address)", &numberB) */

	var number = 4

	var nama string = "alireza"

	changename(&nama, "Stifan")

	fmt.Println("name:", nama)

	fmt.Println("before:", number)

	change(&number, 10)
	fmt.Println("after:", number)
}

func changename(name *string, val string) {
	*name = val
}

func change(original *int, value int) {
	*original = value
}
