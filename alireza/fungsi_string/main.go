package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		isExists := strings.Contains("john wick", "wick")
		fmt.Println(isExists) */

	// ==  	string Join =====

	data := []string{"banana", "pepaya", "tomatto"}

	str := strings.Join(data, " ")

	fmt.Println(str)
}
