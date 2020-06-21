package main

import "fmt"

func main() {

	/* chicken := make(map[int]string)

	chicken[1] = "ali"
	chicken[2] = "reza"
	chicken[3] = "ganteng"

	fmt.Println(chicken)

	var bulan = map[string]int{"january": 1, "february": 2, "maret": 3, "april": 4}

	for i, v := range bulan {
		fmt.Println(i, v)
	}

	fmt.Println(bulan)

	delete(bulan, "january")

	fmt.Println("Panjang:", len(bulan))

	fmt.Println(bulan) */

	//detect specific item

	/* var chicken = map[string]int{"january": 1, "february": 2}

	var value, isExist = chicken["january"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("item not exist")
	}
	*/

	var players = []map[string]string{
		{"name": "dybala", "club": "juventus", "country": "argentina"},
		{"name": "messi", "club": "barcelona", "country": "argentina"},
		{"name": "ronaldo", "club": "juventus", "country": "portugal"},
		{"name": "lukaku", "club": "inter", "country": "belgium"},
	}

	for i, v := range players {
		fmt.Println(i, v)
	}
}
