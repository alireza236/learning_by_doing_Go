package main

import (
	"fmt"
	"runtime"
)

func printMessage(what chan string) {
	fmt.Println(<-what)

}

func main() {

	runtime.GOMAXPROCS(4)

	messages := make(chan string)

	user := [...]string{"wick", "hunt", "bourne", "warno"}

	for _, each := range user {
		go func(who string) {
			var data = fmt.Sprintf(" Hello %s", who)
			messages <- data
		}(each)
	}

	for i := 0; i < 4; i++ {
		printMessage(messages)

	}
	/* messages := make(chan string)

	sayHello := func(who string) {
		data := fmt.Sprintf(" Hello %s", who)
		messages <- data
	}

	go sayHello("John Wick")
	go sayHello("Maria Bellina")
	go sayHello("Dian nitami")

	messages1 := <-messages

	fmt.Println(messages1)

	messages2 := <-messages
	fmt.Println(messages2)

	messages3 := <-messages
	fmt.Println(messages3) */

}
