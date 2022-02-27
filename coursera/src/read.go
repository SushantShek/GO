package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var input string

	slice := make([]Name, 0, 3)
	fmt.Println("Enter file name")

	fmt.Scan(&input)

	f, err := os.Open(input)
	check(err)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), " ")
		var namestruct Name
		namestruct.fname, namestruct.lname = s[0], s[1]
		slice = append(slice, namestruct)
	}
	defer f.Close()

	for _, len := range slice {
		fmt.Println(len.fname, len.lname)
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
