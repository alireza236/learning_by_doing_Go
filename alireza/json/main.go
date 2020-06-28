package main

import (
	"encoding/json"
	"fmt"
)

// ==== Decode JSON Ke Variabel Objek Struct ===

// User ...
type User struct {
	Fullname string `json:"name"`
	Age      int    `json:"age"`
}

func main() {

	/* jsonString := `{ "name":"John wick", "age": 31 }`

	jsonData := []byte(jsonString)

	var data User

	err := json.Unmarshal(jsonData, &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Name: ", data.Fullname)
	fmt.Println("Age: ", data.Age) */

	// ======   Decode JSON Ke map[string]interface{} & interface{} =====

	/* var data1 map[string]interface{}

	jsonString := `{ "name":"John wick", "age": 31 }`

	jsonData := []byte(jsonString)

	json.Unmarshal(jsonData, &data1)

	var data2 interface{}

	json.Unmarshal(jsonData, &data2)

	var decodeData = data2.(map[string]interface{})

	fmt.Println("user2", decodeData["name"])
	fmt.Println("age2", decodeData["age"])

	fmt.Println("user:", data1["Name"])
	fmt.Println("age:", data1["Age"]) */

	// ==== Decode Array JSON Ke Array Objek ====

	/* var jsonString = `[
		{"Name" : "John wick", "Age": 28},
		{"Name" : "Kirk Hammet", "Age": 33}
	]`

	var data []User

	var err = json.Unmarshal([]byte(jsonString), &data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user: 1", data[0].Fullname)
	fmt.Println("user: 2", data[1].Fullname) */

	object := []User{{"John wick", 27}, {"Ethan hunt", 32}}

	jsonData, err := json.Marshal(object)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var jsonString = string(jsonData)

	fmt.Println(jsonString)

}
