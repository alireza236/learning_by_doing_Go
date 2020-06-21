package main

import (
	"fmt"
	"os"
)

func orderSomeFood(menu string) {
	defer fmt.Println("Thank's, Silakan tunggu")

	if menu == "pizza" {
		fmt.Print("Pilihan tepat", " ")
		fmt.Print("Pizza ditempat kami paling enak!", "\n")
		return
	}

	fmt.Println("pesanan anda", menu)

}

func main() {

	orderSomeFood("pizza")
	orderSomeFood("burger")

	number := 3

	os.Exit(1)

	if number == 3 {
		fmt.Println("Halo 1")
		func() {
			defer fmt.Println("Halo 3")
		}()
	}

	fmt.Println("Halo 2")

}
