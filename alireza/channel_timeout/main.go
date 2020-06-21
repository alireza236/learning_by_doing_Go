package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func sendData(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Int()%10+1) * time.Second)
	}
}

func retrieveData(ch <-chan int) {
loop:
	for {
		select {
		case data := <-ch:
			fmt.Print(`receive data "`, data, `"`, "\n")
		case <-time.After(time.Second * 5):
			fmt.Println("timeout... no activities under 5 second ")
			break loop
		}
	}
}

func main() {
	fmt.Println("Start..")

	rand.Seed(time.Now().Unix())

	runtime.GOMAXPROCS(4)

	messages := make(chan int)

	go sendData(messages)

	retrieveData(messages)

}
