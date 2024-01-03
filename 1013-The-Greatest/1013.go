package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, greatest float64

	fmt.Scan(&a, &b, &c)

	greatest = (a + b + math.Abs(a-b)) / 2
	greatest = (greatest + c + math.Abs(greatest-c)) / 2

	fmt.Printf("%d eh o maior\n", int(greatest))

}
