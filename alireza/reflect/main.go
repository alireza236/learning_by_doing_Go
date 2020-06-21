package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println(" Nama :", reflectType.Field(i).Name)
		fmt.Println(" Nama :", reflectType.Field(i).Type)
		fmt.Println(" Nama :", reflectValue.Field(i).Interface())

		fmt.Println(" ")
	}
}

func (s *student) setName(name string) {
	s.Name = name
}

func main() {

	var s1 = &student{Name: "john wick", Grade: 2}

	fmt.Println(" nama :", s1.Name)

	var reflectValue = reflect.ValueOf(s1)

	var method = reflectValue.MethodByName("SetName")

	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println(" nama :", s1.Name)
	/* s1.getPropertyInfo() */

	/* var number = 23.0

	var reflectValue = reflect.ValueOf(number)

	fmt.Println(" type variable :", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println(" Nilai variable :", reflectValue.Int())
	}

	fmt.Println(" type variabel :", reflectValue.Type())

	fmt.Println(" Nilai variable :", reflectValue.Interface()) */
}
