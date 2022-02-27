package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const numberOfArguments = 3
const commandQuery = "query"
const commandCreateAnimal = "newanimal"

var cowData = [...]string{"grass", "walk", "moo"}
var birdData = [...]string{"worms", "fly", "peep"}
var snakeData = [...]string{"mice", "slither", "hsss"}

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (a Cow) Eat() {
	fmt.Println(a.food)
}

func (a Cow) Move() {
	fmt.Println(a.locomotion)
}

func (a Cow) Speak() {
	fmt.Println(a.noise)
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (a Bird) Eat() {
	fmt.Println(a.food)
}

func (a Bird) Move() {
	fmt.Println(a.locomotion)
}

func (a Bird) Speak() {
	fmt.Println(a.noise)
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (a Snake) Eat() {
	fmt.Println(a.food)
}

func (a Snake) Move() {
	fmt.Println(a.locomotion)
}

func (a Snake) Speak() {
	fmt.Println(a.noise)
}

func main() {
	animalMap := make(map[string]Animal)
	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			request := scanner.Text()
			splittedReq := strings.Split(strings.ToLower(request), " ")
			if len(splittedReq) == numberOfArguments {
				query := strings.TrimSpace(splittedReq[0])
				arg1 := strings.TrimSpace(splittedReq[1])
				arg2 := strings.TrimSpace(splittedReq[2])

				switch query {
				case commandQuery:
					handleQuery(animalMap, arg1, arg2)
				case commandCreateAnimal:
					animalMap = handleCreateAnimal(animalMap, arg1, arg2)
				default:
					fmt.Println("wrong option!")
				}
			}
		}
	}
}

func handleQuery(myMap map[string]Animal, animalName string, animalInfo string) {
	PrintAnimalInfo(myMap[animalName], animalInfo)
}

func handleCreateAnimal(myMap map[string]Animal, animalName string, animalType string) map[string]Animal {
	a := CreateAnimal(animalType)
	if a != nil {
		myMap[animalName] = a
		fmt.Println("“Created it!")
	} else {
		fmt.Println("Not added !")
	}
	return myMap
}

func PrintAnimalInfo(a Animal, animalInfo string) {
	if a != nil {
		switch animalInfo {
		case "eat":
			a.Eat()
		case "move":
			a.Move()
		case "speak":
			a.Speak()
		default:
			fmt.Println("Wrong command !")
		}
	} else {
		fmt.Println("Wrong name !")
	}
}

func CreateAnimal(animalType string) Animal {
	switch animalType {
	case "cow":
		return Cow{cowData[0], cowData[1], cowData[2]}
	case "bird":
		return Bird{birdData[0], birdData[1], birdData[2]}
	case "snake":
		return Snake{snakeData[0], snakeData[1], snakeData[2]}
	}
	return nil
}
