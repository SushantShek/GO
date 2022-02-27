package main

import (
	"fmt"
	"strings"
)

var input string

func main() {
	fmt.Println("Enter String value")

	fmt.Scan(&input)
	var x = strings.ToLower(input)
	if strings.HasPrefix(x, "i") && strings.HasSuffix(x, "n") && strings.Contains(x, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
