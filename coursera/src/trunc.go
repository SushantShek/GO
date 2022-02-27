package main

import "fmt"

var input float64

func main() {
	fmt.Println("Enter floating point number")

	fmt.Scan(&input)
	floatToInt()

	fmt.Println("Enter floating number again")

	fmt.Scan(&input)
	floatToInt()
}

func floatToInt() {

	var output = int(input)
	fmt.Println("Int ", output)
}
