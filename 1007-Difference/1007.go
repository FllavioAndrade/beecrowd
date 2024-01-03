package main

import "fmt"

func main() {
	var a, b, c, d, dif int

	fmt.Scan(&a, &b, &c, &d)

	dif = a*b - c*d

	fmt.Printf("DIFERENCA = %d\n", dif)

}
