package main

import (
	"fmt"
	"net/url"
)

func main() {

	urlString := "http://localhost:8080/hello?name=john wick&age=27"

	u, e := url.Parse(urlString)

	if e != nil {
		fmt.Println(e.Error())
	}

	fmt.Printf("url: %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme) //http
	fmt.Printf("host: %s\n", u.Host)       //
	fmt.Printf("path: %s\n", u.Path)

	name := u.Query()["name"][0]
	age := u.Query()["age"][0]

	fmt.Printf("name : %s, age: %s\n", name, age)

}
