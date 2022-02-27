package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter upto 10 integer values comma seperated")
	var input string
	slice := make([]int, 0)

	fmt.Scan(&input)
	s := strings.Split(input, ",")

	for i := 0; i < len(s); i++ {
		intVar, err := strconv.Atoi(s[i])
		check(err)
		slice = append(slice, intVar)
	}
	BubbleSort(slice)

}

func BubbleSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		for j := 0; j < len(sli)-i-1; j++ {
			Swap(sli, j)
		}
	}
	fmt.Println(sli)
}

func Swap(sli []int, pos int) {

	if sli[pos] > sli[pos+1] {
		sli[pos], sli[pos+1] = sli[pos+1], sli[pos]
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
