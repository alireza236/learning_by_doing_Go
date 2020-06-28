package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var baseurl = "http://localhost:8080"

type student struct {
	ID    string
	Name  string
	Grade int
}

func fetchUsers() ([]student, error) {
	var err error
	var client = &http.Client{}
	var data []student

	request, err := http.NewRequest("POST", baseurl+"/users", nil)

	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		return nil, err
	}

	return data, nil

}

func fetchUser(ID string) (student, error) {
	var err error
	var client = &http.Client{}
	var data student

	var params = url.Values{}

	params.Set("id", ID)

	var payload = bytes.NewBufferString(params.Encode())

	request, err := http.NewRequest("POST", baseurl+"/user", payload)

	if err != nil {
		return data, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)

	if err != nil {
		return data, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		return data, err
	}

	return data, nil
}

func main() {
	fmt.Println("Hallo bro")

	//users, err := fetchUsers()

	user1, err := fetchUser("E001")

	if err != nil {
		fmt.Println("Error", err.Error())
	}

	//for _, each := range users {
	//	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", each.ID, each.Name, each.Grade)
	//}

	fmt.Printf("ID: %s\t Name: %s\t Grade: %d\n", user1.ID, user1.Name, user1.Grade)

}
