package main

import (
	"fmt"
	"runtime"
)

func run(pesan string, messages chan int) {
	for {
		i := <-messages
		fmt.Println(pesan, "receive data", i)
	}
}

func main() {

	runtime.GOMAXPROCS(4)

	messages := make(chan int, 2)

	go run("run 1", messages)

	for i := 0; i < 3; i++ {
		fmt.Println("send data", i)
		messages <- i
	}
}
