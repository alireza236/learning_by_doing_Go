package main

import "fmt"

func main() {
	//slice

	/* var fruits = []string{"apple", "melon", "orange", "grepe"}

	var bFruits = fruits[0:2]
	fmt.Println(cap(bFruits))
	fmt.Println(len(bFruits))

	fmt.Println(fruits)
	fmt.Println(bFruits)

	var cFruits = append(bFruits, "pepaya")

	fmt.Println(fruits)
	fmt.Println(bFruits)
	fmt.Println(cFruits) */

	/* dst := make([]string, 3)
	src := []string{"watermelon", "pinneaple", "apple", "orange"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(n) */

	/* dst := []string{"kerupuk", "kerupuk", "kerupuk"}

	src := []string{"bawang", "cabe"}

	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n) */

	var fruits = []string{"apple", "grape", "banana"}

	var aFruits = fruits[0:2]
	var bFruits = fruits[0:2:2]

	fmt.Println(fruits)
	fmt.Println(len(fruits))
	fmt.Println(cap(fruits))

	fmt.Println(aFruits)
	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))

	fmt.Println(bFruits)
	fmt.Println(len(bFruits))
	fmt.Println(cap(bFruits))

}
