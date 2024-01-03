package main

import (
	"fmt"
	"math"
)

func main() {
	var radius int
	var pi = 3.14159
	var volume float64

	fmt.Scan(&radius)

	volume = (4.0 / 3) * pi * math.Pow(float64(radius), 3)

	fmt.Printf("VOLUME = %.3f\n", volume)

}
