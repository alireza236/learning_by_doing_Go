package main

import (
	"fmt"
)

func main() {

	/* str := "1234"

	num, err := strconv.Atoi(str)
	if err != nil {

		fmt.Println(err.Error())
	}
	fmt.Println(num) */

	/// ==== parseInt =====

	/* str := "1234"

	num, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(num) */

	// ===== Format int =====

	/* 	angka := int64(24)

	   	str := strconv.FormatInt(angka, 8)

	   	fmt.Println(str) */

	// ==== ParseFloat ===

	/* str := "24.12"

	num, err := strconv.ParseFloat(str, 64)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(num) */

	// ==== FormatFloat ====

	/* num := float64(24.14)

	str := strconv.FormatFloat(num, 'f', 6, 64)
	fmt.Println(str) */

	// ==== ParseBool =====

	/* str := "true"

	num, err := strconv.ParseBool(str)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(num) */

	// ===== konversi data mengguakan teknik type casting =====

	/* a := float64(24)

	fmt.Println(a)

	b := int(32)
	c := &b
	fmt.Println(c) */

	//  ==== string <--> byte ====

	/* text1 := "hallo"

	b := []byte(text1)

	fmt.Printf("%d %d %d %d \n", b[0], b[1], b[2], b[3])

	byte1 := []byte{104, 97, 108, 111}

	s := string(byte1)

	fmt.Printf("%s \n", s) */

	// ======= Type Assertions Pada Interface Kosong ( interface{} ) =====

	data := map[string]interface{}{
		"nama":    "john wick",
		"grade":   2,
		"height":  156.5,
		"isMale":  true,
		"hobbies": []string{"eating", "sleeping"},
	}
	/*
		fmt.Println(data["nama"].(string))
		fmt.Println(data["grade"].(int))
		fmt.Println(data["height"].(float64))
		fmt.Println(data["isMale"].(bool))
		fmt.Println(data["hobbies"].([]string)) */

	for _, each := range data {
		switch each.(type) {
		case string:
			fmt.Println(each.(string))
		case int:
			fmt.Println(each.(int))
		case float64:
			fmt.Println(each.(float64))
		case bool:
			fmt.Println(each.(bool))
		case []string:
			fmt.Println(each.([]string))
		default:
			fmt.Println(each.(int))

		}
	}

}
