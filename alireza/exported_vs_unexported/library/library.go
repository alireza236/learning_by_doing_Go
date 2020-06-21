package library

import "fmt"

// Student ...
var Student = struct {
	Name  string
	Grade int
}{}

func init() {
	Student.Name = "Alireza"
	Student.Grade = 12

	fmt.Println("--> library/library.go imported ")
}
