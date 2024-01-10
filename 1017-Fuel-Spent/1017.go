package main

import "fmt"

func main() {

	var hour, average_speed int
	var liters float64

	fmt.Scan(&hour, &average_speed)

	liters = (float64(hour) * float64(average_speed)) / 12

	fmt.Printf("%.3f\n", liters)
}
