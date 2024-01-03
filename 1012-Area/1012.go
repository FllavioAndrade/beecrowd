package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c float64
	var triangle, circle, trapezium, square, rectangle float64
	var pi = 3.14159
	fmt.Scan(&a, &b, &c)

	triangle = (a * c) / 2.0
	circle = pi * math.Pow(c, 2)
	trapezium = (a + b) / 2.0 * c
	square = math.Pow(b, 2)
	rectangle = a * b

	fmt.Printf("TRIANGULO: %.3f\nCIRCULO: %.3f\nTRAPEZIO: %.3f\nQUADRADO: %.3f\nRETANGULO: %.3f\n\n", triangle, circle, trapezium, square, rectangle)

}
