package main

import (
	"fmt"
)

func main() {
	var name string
	var salary, saller, total float64

	fmt.Scan(&name, &salary, &saller)

	total = salary + saller*0.15

	fmt.Printf("TOTAL = R$ %.2f", total)
}
