package main

import "fmt"

func main() {

	var r, A float64

	fmt.Scan(&r)

	A = r * r * 3.14159

	fmt.Printf("A=%.4f\n", A)
}
