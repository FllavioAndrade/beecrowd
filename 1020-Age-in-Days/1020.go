package main

import "fmt"

func main() {

	var days, period int

	fmt.Scan(&days)
	period = days / 365
	fmt.Printf("%d ano(s)\n", period)
	days = days % 365
	period = days / 30
	fmt.Printf("%d mes(es)\n", period)
	days = days % 30
	fmt.Printf("%d dia(s)\n", days)

}
