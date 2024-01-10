package main

import "fmt"

func main() {

	var time, secounds, minutes, hours int

	fmt.Scan(&time)

	hours = time / 3600
	time = time % 3600

	minutes = time / 60
	time = time % 60

	secounds = time % 60

	fmt.Printf("%d:%d:%d\n", hours, minutes, secounds)

}
