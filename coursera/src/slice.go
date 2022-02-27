package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var userInt string

	var intSlice = make([]int, 3)

	for x := 0; ; x++ {
		fmt.Printf("Enter an integer :\n")
		fmt.Scan(&userInt)

		if userInt == "X" || userInt == "x" {
			break
		}
		i, err := strconv.Atoi(userInt)
		if err != nil {
			fmt.Println(err)
			os.Exit(2)
		}
		if x > 2 {
			intSlice = append(intSlice, int(i))
		} else {
			intSlice[x] = i
		}
	}

	sort.Ints(intSlice)
	fmt.Println("Sorted Slice :", intSlice)
}
