package main

import (
	"fmt"
	"runtime"
)

func sendMessage(ch chan<- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("Data %d", i)
	}
	close(ch)
}

func printMessage(ch <-chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {

	runtime.GOMAXPROCS(4)

	messages := make(chan string)

	go sendMessage(messages)

	printMessage(messages)
}
