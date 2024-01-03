package main

import "fmt"

func main() {

	var x, y, z float32

	fmt.Scan(&x, &y)
	z = x*3.5 + y*7.5
	fmt.Printf("MEDIA = %.5f\n", z/11)
}
