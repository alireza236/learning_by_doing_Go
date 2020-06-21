package main

import "fmt"

func main() {

	var firstName string = "ali"

	var lastName string
	lastName = "reza"

	name := new(string)

	fmt.Printf("%v\n", *name)

	fmt.Printf("halo %s %s!\n", firstName, lastName)
}
