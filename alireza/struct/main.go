package main

import "fmt"

// Person ...
type Person struct {
	name string
	age  int
}

// Student ...
/* type Student struct {
	grade int
	age   int
	Person
} */

func main() {

	/* var person Student

	person.name = "alireza"

	person.grade = 12

	fmt.Println(person)

	fmt.Println("name:", person.name)
	fmt.Println("grade:", person.grade) */

	/* 	var person1 = Student{name: "john", grade: 2}

	   	var person2 *Student = &person1

	   	fmt.Println("student 1: name:", person1.name)
	   	fmt.Println("student 1: name:", person2.name)

	   	person2.name = "greddy"

	   	fmt.Println("student 1: name:", person1.name)
	   	fmt.Println("student 1: name:", person2.name) */

	// embeded Struct

	/* var murid Student

	murid.name = "alireza"
	murid.age = 23

	murid.Person.age = 24

	murid.grade = 10

	fmt.Println("name : ", murid.name)
	fmt.Println("age : ", murid.age)
	fmt.Println("grade : ", murid.grade)
	fmt.Println(" person age : ", murid.Person.age) */

	/* var murid = struct {
		Person
		grade int
	}{
		Person: Person{"wick", 21},
		grade:  2,
	}

	fmt.Println("age :", murid.Person.age)
	fmt.Println("name :", murid.Person.name)
	fmt.Println("grade :", murid.grade) */

	// Combinasi struct dan slice

	/* var allStudent = []Person{
		{name: "ali reza", age: 12},
		{name: "malih", age: 75},
		{name: "Safri", age: 43},
	}

	for _, student := range allStudent {
		fmt.Println(student.name, "age is:", student.age)
	} */

	// Initialize anonymous struct and slice

	var allStudent = []struct {
		Person
		grade int
	}{
		{Person: Person{"ali", 21}, grade: 12},
		{Person: Person{"ferry", 31}, grade: 22},
		{Person: Person{"Bony", 34}, grade: 54},
	}

	for _, student := range allStudent {
		fmt.Println(student)
	}

}
