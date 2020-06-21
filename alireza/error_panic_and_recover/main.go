package main

import (
	"errors"
	"fmt"
	"strings"
)

func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}

	return true, nil
}

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured")

	} else {
		fmt.Println("Application running perfectly")
	}
}

func main() {

	/* 	defer catch()

	   	var name string

	   	fmt.Print("Type some number :")

	   	fmt.Scanln(&name)

	   	if ok, err := validate(name); ok {
	   		fmt.Println("halo", name)
	   	} else {
	   		fmt.Println("End")
	   		panic(err.Error())
	   	} */

	data := []string{" superman", "aquaman", "wonder woman"}

	for _, each := range data {
		func() {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic ocuured on looping", each, "| messages:", r)
				} else {
					fmt.Println("Application running perfectly")
				}
			}()

			panic("Some error happen")
		}()
	}

}
