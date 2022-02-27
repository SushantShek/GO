package main

import "fmt"

func main() {
	var userFloat float32

	fmt.Printf("Enter a floating point number:\n")

	fmt.Scan(&userFloat)

	var userInt int = int(userFloat)

	fmt.Printf("Your number as an integer:\n%d", userInt)
}
