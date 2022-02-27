package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	eat, move, speak string
}

func (a Animal) Eat() {
	fmt.Println(a.eat)
}
func (a Animal) Move() {
	fmt.Println(a.move)
}
func (a Animal) Speak() {
	fmt.Println(a.speak)
}
func (a Animal) GetFunc(name string) func() {
	switch name {
	case "eat":
		return a.Eat
	case "move":
		return a.Move
	case "speak":
		return a.Speak
	default:
		fmt.Println("Cannot perform action.")
		os.Exit(1)
	}
	return foo
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

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
		case "cow":
			cow.GetFunc(command[1])()
		case "bird":
			bird.GetFunc(command[1])()
		case "snake":
			snake.GetFunc(command[1])()
		default:
			fmt.Println("Not an allowed animal")
			os.Exit(1)
		}
		fmt.Print("> ")
	}

}

func foo() {
	fmt.Print("")
}
