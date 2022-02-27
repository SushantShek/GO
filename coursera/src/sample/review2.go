package main

import (
	"fmt"
	"strings"
)

func main() {
	var userString string

	fmt.Printf("Enter a string:\n")

	fmt.Scan(&userString)

	userString = strings.ToLower(userString)

	var hasIan bool = userString[0:1] == "i" &&
		userString[len(userString)-1:] == "n" &&
		strings.Contains(userString, "a")

	if hasIan {
		fmt.Printf("Found!")
	} else {
		fmt.Printf("Not Found!")
	}
}
