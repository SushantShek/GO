package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Printf("Sorted slice:\n%v", arr)
}

func SequenceInput() []int {
	fmt.Println("Input a sequence of numbers (separated by whitespace):")
	reader := bufio.NewReader(os.Stdin)
	val, _, err := reader.ReadLine()
	if err != nil {
		log.Fatalf("Cannot read line: %s\n", err)
	}
	numberListTmp := strings.Split(string(val), " ")
	var numberList []int
	for _, v := range numberListTmp {
		num, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Cannot convert '%s' to number: %s\n", v, err)
		}
		numberList = append(numberList, num)
	}
	return numberList
}
func main() {
	inputList := SequenceInput()
	BubbleSort(inputList)
}
