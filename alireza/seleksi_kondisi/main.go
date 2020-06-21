package main

import "fmt"

func main() {
	/* 	var point int = 0

	   	if point == 10 {
	   		fmt.Println("lulus dengan nilai sempurna")
	   	} else if point > 5 {
	   		fmt.Println("lulus")
	   	} else if point == 4 {
	   		fmt.Println("hampir lulus")
	   	} else {
	   		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	   	} */

	/* var point float64 = 884.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s sempurna\n", percent, "%")
	} else if percent >= 100 {
		fmt.Printf("%.1f%s good\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad\n", percent, "%")

	} */

	/* 	var point = 6

	   	switch point {
	   	case 8:
	   		fmt.Println("perfect")
	   	case 7, 6, 5, 4:
	   		fmt.Println("awesome")
	   	default:
	   		fmt.Println("not bad")

	   		} */

	/* var point int = 0

	switch {
	case point == 8:
		fmt.Println("sempurna")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you nedd learn more")
		}

	} */

	/* 	var point int = 6

	   	switch {
	   	case point == 8:
	   		fmt.Println("perfect")
	   	case (point < 8) && (point > 3):
	   		fmt.Println("awesome")
	   		fallthrough
	   	case point < 5:
	   		fmt.Println("you need to learn more")
	   	default:
	   		{
	   			fmt.Println("not bad")
	   			fmt.Println("you need to learn more")

	   		}
	   	} */

	var point int = 4

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect")
		case 5:
			fmt.Println("kurang")
		default:
			fmt.Println("nice")
		}
	} else {
		if point == 5 {
			fmt.Println("not Bad")
		} else if point == 3 || point < 2 {
			fmt.Println("tetap lah mencoba")
		} else {
			fmt.Println("you can do it")
			if point == 0 {
				fmt.Println("coba lebih keras lagi")
			}
		}
	}
}
