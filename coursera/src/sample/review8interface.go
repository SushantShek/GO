package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	Name string
}

func (a Cow) Eat() {
	fmt.Println("grass")
}
func (a Cow) Move() {
	fmt.Println("walk")
}
func (a Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {
	Name string
}

func (a Bird) Eat() {
	fmt.Println("worms")
}
func (a Bird) Move() {
	fmt.Println("fly")
}
func (a Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {
	Name string
}

func (a Snake) Eat() {
	fmt.Println("mice")
}
func (a Snake) Move() {
	fmt.Println("slither")
}
func (a Snake) Speak() {
	fmt.Println("hsss")
}

func main() {

	animalMap := make(map[string]Animal)

	fmt.Print("> ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var line string
		var command []string
		line = scanner.Text()
		// fmt.Println("Got line:", line)
		for _, x := range strings.Fields(line) {
			// fmt.Println("next field: ", x)
			command = append(command, x)
		}

		switch command[0] {
		case "newanimal":
			switch command[2] {
			case "cow":
				animalMap[command[1]] = Cow{command[1]}
			case "bird":
				animalMap[command[1]] = Bird{command[1]}
			case "snake":
				animalMap[command[1]] = Snake{command[1]}
			default:
				fmt.Println("Not an allowed animal")
				os.Exit(1)
			}
			fmt.Println("Created it!")
		case "query":
			switch command[2] {
			case "eat":
				animalMap[command[1]].Eat()
			case "move":
				animalMap[command[1]].Move()
			case "speak":
				animalMap[command[1]].Speak()
			default:
				fmt.Println("Cannot perform action.")
				os.Exit(1)
			}
		default:
			fmt.Println("Not an allowed command")
			os.Exit(1)
		}
		fmt.Print("> ")
	}

}

func foo() {
	fmt.Print("")
}
