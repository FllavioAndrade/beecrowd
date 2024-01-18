package main

import (
	"fmt"
	"math"
)

func operation(n1 int, n2 int, d1 int, d2 int, ope string) (int, int, int, int) {

	var numerator, denominator int

	if ope == "+" {
		numerator = n1*d2 + n2*d1
		denominator = d1 * d2
	} else if ope == "-" {
		numerator = n1*d2 - n2*d1
		denominator = d1 * d2
	} else if ope == "*" {
		numerator = n1 * n2
		denominator = d1 * d2
	} else {
		numerator = n1 * d2
		denominator = n2 * d1
	}
	numerator1 := numerator
	denominator1 := denominator
	maior := math.Max(float64(denominator), float64(numerator))

	for j := 2; j <= int(maior); j++ {
		for {
			if numerator1%j != 0 || denominator1%j != 0 {
				break
			}
			numerator1 = numerator1 / j
			denominator1 = denominator1 / j
		}

	}
	return numerator, numerator1, denominator, denominator1

}

func main() {

	var n1, n2, d1, d2, qtd int
	var div, ope string

	fmt.Scan(&qtd)
	for i := 0; i < qtd; i++ {
		fmt.Scan(&n1, &div, &d1, &ope, &n2, &div, &d2)

		n1, n2, d1, d2 = operation(n1, n2, d1, d2, ope)

		fmt.Printf("%d/%d = %d/%d\n", n1, d1, n2, d2)

	}

}
