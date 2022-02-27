package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var input1 string

func main() {
	m := make(map[string]string)
	fmt.Println("Enter name")
	fmt.Scan(&input1)

	m["name"] = input1

	fmt.Println("Enter address")
	fmt.Scan(&input1)

	m["address"] = input1
	// m["1"] = JsonObj{Name: input1, Address: input2}
	jsonString, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(string(jsonString))
}
