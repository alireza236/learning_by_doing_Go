package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	// ===== TIMER =======

	/* second := time.Second

	fmt.Println("start")

	time.Sleep(second * 4)

	fmt.Println("after 4 second")

	fmt.Println("time Duration", int64(time.Duration(second)*time.Second))
	*/

	/* 	for i := 0; i < 10; i++ {

		fmt.Println("Hello !!", i)
		time.Sleep(1 * time.Second)

		if i >= 9 {
			recover()
		}
	} */

	// ===== NewTimer =====

	/* timer := time.NewTimer(4 * time.Second)

	fmt.Println("start")

	data := <-timer.C

	fmt.Println("StopFinish", data) */

	//===== AfterFunc =====

	/* ch := make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		fmt.Println("expired")

		for i := 0; i < 15; i++ {

			if i >= 10 {
				ch <- false
				return
			}
		}
	})

	exp := <-time.After(4 * time.Second)
	fmt.Println("not expired", exp)

	fmt.Println("start")

	res := <-ch
	fmt.Println("Finish", res) */

	//===== TICKER ======

	/* 	done := make(chan bool)

	   	ticker := time.NewTicker(time.Second)

	   	go func() {
	   		time.Sleep(60 * time.Second) // wait for 10 second

	   		done <- true
	   	}()

	   	for {
	   		select {
	   		case <-done:
	   			ticker.Stop()
	   			return
	   		case t := <-ticker.C:
	   			fmt.Println("hello", t.Second())

	   		}
		   } */

	timeout := 5

	ch := make(chan bool)

	go timer(timeout, ch)
	go watcher(timeout, ch)

	var input string

	fmt.Print("what is 725/25 ?")

	fmt.Scan(&input)

	if input == "29" {
		fmt.Println("the answer is right")
	} else {
		fmt.Println("the answer is wrong")

	}

}

func timer(timeout int, ch chan<- bool) {
	time.AfterFunc(time.Duration(timeout)*time.Second, func() {
		ch <- true
	})
}

func watcher(timeout int, ch <-chan bool) {
	<-ch

	fmt.Println("\nwaktu habis! no answer more than", timeout, "seconds")
	os.Exit(0)

}
