package main

import "fmt"

func main() {

	var code1, units1, code2, units2 int
	var price1, price2, total float32

	fmt.Scan(&code1, &units1, &price1)
	fmt.Scan(&code2, &units2, &price2)

	total = float32(units1)*price1 + float32(units2)*price2

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)

}
