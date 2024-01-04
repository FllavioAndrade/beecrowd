package main

import "fmt"

func main() {

	var distance int
	var fuel_total, distance_fuel float64

	fmt.Scan(&distance, &fuel_total)

	distance_fuel = float64(distance) / fuel_total

	fmt.Printf("%.3f km/l\n", distance_fuel)

}
