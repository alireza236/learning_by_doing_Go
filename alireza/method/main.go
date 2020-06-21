package main

import (
	"fmt"
	"strings"
)

// Student ...
/* type Student struct {
	name  string
	grade int
}

func (s Student) sayHello() {
	fmt.Println("hallo", s.name)
}

func (s Student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
} */

// Student ...
/* type Student struct {
	name  string
	grade int
}

func (s Student) changeName1(name string) {
	fmt.Println("---> on changeName1, chnage name to : ", name)

	s.name = name
}

func (s *Student) changeName2(name string) {
	fmt.Println("---> on changeName2, change name to", name)

	s.name = name
} */

/*
func main() {

	var murid = Student{"john wick", 21}

	   	murid.sayHello()

	   	var name = murid.getNameAt(2)
		   fmt.Println("nama panggilan", name)

		   var murid = Student{"john wick", 22}

		   fmt.Println("murid before, ", murid.name)

		   murid.changeName1("Alireza")

		   fmt.Println("murid after chanrgeName", murid.name)

		   murid.changeName2("adam soleh")

		   fmt.Println("murid after changeName2", murid.name)

		}
*/
type student struct {
	name  string
	grade int
}

func (s student) sayHello() {
	fmt.Println("Halo", s.name)
}

func (s student) getNameAt(i int) string {
	return strings.Split(s.name, " ")[i-1]
}

func main() {
	var s1 = student{"john wick", 21}
	s1.sayHello()

	var name = s1.getNameAt(2)

	fmt.Println("nama panggilan:", name)
}
