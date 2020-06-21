package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	/* var names = []string{"john", "wick"}
	printMessage("Hallo", names) */

	/* rand.Seed(time.Now().Unix())

	var randomValue int

	randomValue = randomWithRange(2, 10)
	fmt.Println("random value", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("random value", randomValue)

	randomValue = randomWithRange(2, 10)
	fmt.Println("random value", randomValue) */

	divideNumber(10, 0)
	divideNumber(4, 2)
	divideNumber(8, -4)

}

func divideNumber(m, n int) {
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d \n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d\n", m, n, res)

}

func randomWithRange(min, max int) int {
	var value = rand.Int()%(max-min) + min
	return value
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}
