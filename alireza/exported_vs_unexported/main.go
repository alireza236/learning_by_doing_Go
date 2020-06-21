package main

import (
	"fmt"

	"gitlab.com/alireza/go_novalagung/github.com/alireza/exported_vs_unexported/library"
)

func main() {

	fmt.Printf("name  %s\n", library.Student.Name)

	fmt.Printf("Grade %d\n", library.Student.Grade)

}
