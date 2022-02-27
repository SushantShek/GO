package main

import (
	"fmt"
	"math"
)

func main() {

	var time, acceleration, init_velocity, init_displacement float64

	fmt.Println("Enter float value for acceleration")
	fmt.Scan(&acceleration)

	fmt.Println("Enter float value for initial velocity")
	fmt.Scan(&init_velocity)

	fmt.Println("Enter float value for initial displacement")
	fmt.Scan(&init_displacement)

	fmt.Println("Enter float value for time")
	fmt.Scan(&time)

	fn := GenDisplaceFn(acceleration, init_velocity, init_displacement)
	fmt.Println("Dispalcement ", fn(time))
}

func GenDisplaceFn(a, v, s float64) func(float64) float64 {

	fn := func(time float64) float64 {
		return (((a * math.Pow(time, 2)) / 2) + (v * time) + s)
	}
	return fn
}
