package main

import (
	"fmt"
)

func main() {

	/* var secret interface{}

	secret = "Mariam"

	fmt.Println(secret)

	secret = []string{}

	secret = "ayam"

	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)
	*/

	/* var secret interface{}

	secret = 2

	var number = secret.(int) * 10

	fmt.Println(secret, "Multiplied by 10 is :", number)

	secret = []string{"apple", "banana", "manggo"}

	var fruits = strings.Join(secret.([]string), ",")
	fmt.Println(fruits, "is my favorite fruits") */

	/* type person struct {
		name string
		age  int
	}

	var secret interface{} = &person{name: "wick", age: 27}

	var name = secret.(*person).name

	fmt.Println(name) */

	/// KOMBINASI SLICE, MAP DAN INTERFACE

	/* 	var person = []map[string]interface{}{
	   		{"name": "reza", "age": 24},
	   		{"name": "malik", "age": 26},
	   		{"name": "ferry", "age": 24},
	   	}

	   	for _, each := range person {
	   		fmt.Println(each["name"], "age is", each["age"])
	   	} */

	var fruits = []interface{}{
		map[string]interface{}{"name": "starwberry", "total": 10},
		[]string{"manggo", "pennaple", "papaya"},
		"orange",
	}

	for _, each := range fruits {
		fmt.Println(each)
	}

}
