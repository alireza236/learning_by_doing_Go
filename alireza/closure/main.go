package main

import "fmt"

func main() {
	/* var getMinMax = func(n []int) (min int, max int) {
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min:
				min = e
			}
		}
		return min, max
	}

	var numbers = []int{2, 3, 4, 3, 4, 2, 3}

	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin  : %v\nmax  : %v\n", numbers, min, max) */

	//closure with IIFE

	/* var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}

			r = append(r, e)

		}

		return r
	}(3)

	fmt.Println("original number:", numbers)
	fmt.Println("filters number:", newNumbers) */

	// closure sebagai return value

	var max = 3

	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}

	howMany, getNumbers := minMax(numbers, 5)

	var theNumbers = getNumbers()

	fmt.Println("numbers\t", numbers)
	fmt.Printf("find \t: %d\n\n", max)

	fmt.Println("found \t:", howMany)    // 9
	fmt.Println("value \t:", theNumbers) // [2 3 0 3 2 0 2 0 3

}

func minMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _, e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}

	return len(res), func() []int {
		return res
	}

}
