package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		nNumbers  int
		tmpNumber int
	)

	fmt.Println("How many numbers are you inputting?")
	fmt.Scan(&nNumbers)

	sli := make([]int, 0)

	for i := 0; i < nNumbers; i++ {
		fmt.Printf("Please enter number %d\n", i+1)
		fmt.Scan(&tmpNumber)
		sli = append(sli, tmpNumber)

		sort.Slice(sli, func(x, y int) bool {
			return sli[x] < sli[y]
		})

		fmt.Println(sli)
	}
}
