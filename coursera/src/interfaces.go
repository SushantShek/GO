package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Speak()
	Move()
}

type Cow struct {
	eat   string
	speak string
	move  string
}
type Bird struct {
	eat   string
	speak string
	move  string
}
type Snake struct {
	eat   string
	speak string
	move  string
}

func Execute(v Animal, s string) {
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

func (a Cow) Eat() {
	fmt.Println("grass")
}

func (a Cow) Move() {
	fmt.Println("walk")
}

func (a Cow) Speak() {
	fmt.Println("moo")
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
	m := make(map[string]interface{})

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------------------")

	for {
		fmt.Print("> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.ToLower(text)

		split := strings.Split(text, " ")
		var action = split[0]
		var name = split[1]
		var cmd = split[2]

		if "newanimal" == action {
			switch cmd {
			case "cow":
				c := Cow{eat: "grass", move: "walk", speak: "moo"}
				m[name] = c
				break
			case "bird":
				b := Bird{eat: "worms", move: "fly", speak: "peep"}
				m[name] = b
				break
			case "snake":
				s := Snake{eat: "mice", move: "slither", speak: "hsss"}
				m[name] = s
				break
			}
			fmt.Println("Created it!")
		}
		if "query" == action {

			val := m[name]

			var A Animal

			t1, ok0 := val.(Cow)
			if !ok0 {
				t2, ok1 := val.(Snake)
				if !ok1 {
					t3, ok2 := val.(Bird)
					if !ok2 {
						fmt.Println("bad request try again")
					} else {
						// case of("Bird")
						A = t3
					}
				} else {
					// case of("Snake")
					A = t2
				}
			} else {
				// case of("Cow")
				A = t1
			}
			switch cmd {
			case "eat":
				A.Eat()
				break
			case "move":
				A.Move()
				break
			case "speak":
				A.Speak()
				break
			default:
				fmt.Println(A, "Not sure what type item interface")
			}

		}
	}
}
