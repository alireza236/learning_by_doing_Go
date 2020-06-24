package main

import (
	"fmt"
	"os/exec"
)

func main() {

	output1, err := exec.Command("node", "--version").Output()

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf(" -> ls\n%s\n", string(output1))
}
