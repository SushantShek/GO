package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	var person map[string]string
	person = make(map[string]string)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter name")
	scanner.Scan()
	person["name"] = scanner.Text()

	fmt.Println("Enter address")
	scanner.Scan()
	person["address"] = scanner.Text()

	personJSON, err := json.Marshal(person)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("JSON object:")
	fmt.Println(string(personJSON))
}
