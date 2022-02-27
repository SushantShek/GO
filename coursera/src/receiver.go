package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noice      string
}

func Action(v *Animal, s string) {
	switch s {
	case "eat":
		v.Eat()
		break
	case "speak":
		v.Speak()
		break
	case "move":
		v.Move()
		break
	}
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noice)
}

func main() {

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------------")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.ToLower(text)

		split := strings.Split(text, " ")
		var input = split[0]
		var action = split[1]

		switch input {
		case "cow":
			Action(&cow, action)
		case "bird":
			Action(&bird, action)
		case "snake":
			Action(&snake, action)
		}

	}
}
