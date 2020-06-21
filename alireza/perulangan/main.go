package main

import "fmt"

func main() {
	/* for i := 0; i < 5; i++ {
		fmt.Println("angka", i)
	} */

	/* var i int = 0

	for i < 5 {
		fmt.Println("angka", i)
		i++
	} */

	/* var i int = 0

	for {
		fmt.Println("angka mulai", i)
		i++

		if i == 5 {
			continue
		}

		if i >= 10 {
			break
		}
		fmt.Println("angka selesai", i)
	} */

	// for i := 0; i < 20; i++ {
	// 	/* if i%2 == 1 {
	// 		continue
	// 	} */

	// 	if i%2 == 1 {
	// 		fmt.Println("hasil", i)
	// 	}
	// 	if i > 12 {
	// 		break
	// 	}
	// 	fmt.Println("========")
	// 	fmt.Println("angka", i)
	// }

	/* for i := 0; i < 5; i++ {
		for j := i; j < 10; j++ {
			for k := j; k < 10; k++ {
				fmt.Print(k, " ")
			}
		}

		fmt.Println()
	} */

outerLoop:

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}

}
