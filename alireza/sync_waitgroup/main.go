package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(4)

	var wg sync.WaitGroup

	// Wait for goroutines to finish

	for i := 0; i < 10; i++ {
		var data = fmt.Sprintf("data : %d", i)

		wg.Add(2)
		go doOperation(&wg, data)

		go func() {
			for {
				fmt.Println("Break", &wg, data)
			}
		}()
	}
	wg.Wait()

}

func doOperation(wg *sync.WaitGroup, item string) {
	defer wg.Done()
	// do operation on item
	fmt.Println(item)
}
