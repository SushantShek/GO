package main

import (
	"fmt"
	"math"
)

var accel float64
var initVel float64
var initDist float64
var time float64

func GenDisplaceFn(a float64, iv float64, id float64) func(t float64) float64 {
	fn := func(t float64) float64 {
		d := 0.5*a*math.Pow(t, 2) + iv*t + id
		return d
	}
	return fn
}

func main() {
	fmt.Printf("Enter an acceleration value :")
	fmt.Scan(&accel)
	fmt.Printf("Enter an initial velocity :")
	fmt.Scan(&initVel)
	fmt.Printf("Enter an initial distance :")
	fmt.Scan(&initDist)
	fmt.Printf("Enter an elapsed time :")
	fmt.Scan(&time)
	fn := GenDisplaceFn(accel, initVel, initDist)
	fmt.Printf("Displacement travelled : %g", fn(time))
}
