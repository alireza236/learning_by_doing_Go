package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Hallo")

	text := "banana burger soup"

	regex, err := regexp.Compile(`[a-z]+`)

	if err != nil {
		fmt.Println(err.Error())
	}

	res1 := regex.FindAllString(text, 2)

	fmt.Printf("%#v \n", res1)

	res2 := regex.FindAllString(text, -1)
	fmt.Printf("%#v \n", res2)

}
