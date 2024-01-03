package main

import "fmt"

func main() {

	var number, hour int
	var hour_worked, salary float32

	fmt.Scan(&number, &hour, &hour_worked)
	salary = float32(hour) * hour_worked
	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", number, salary)
}
