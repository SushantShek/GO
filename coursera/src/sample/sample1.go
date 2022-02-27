package main

import "fmt"

// type P struct {
// 	x string
// 	y int
// }

// func main() {
// 	s := make([]int, 0, 3)
// 	s = append(s, 100)
// 	fmt.Println(len(s), cap(s))
// }

// func main() {
// 	b := P{"x", -1}
// 	a := [...]P{P{"a", 10},
// 		P{"b", 2},
// 		P{"c", 3}}
// 	for _, z := range a {
// 		if z.y > b.y {
// 			b = z
// 		}
// 	}
// 	fmt.Println(b.x)
// }

// func fA() func() int {
// 	i := 0
// 	return func() int {
// 		i++
// 		return i
// 	}
// }
// func main() {
// 	fB := fA()
// 	fmt.Print(fB())
// 	fmt.Print(fB())
// }

func main() {

	i := 1

	fmt.Print(i)

	i++

	defer fmt.Print(i + 1)

	fmt.Print(i)

}
