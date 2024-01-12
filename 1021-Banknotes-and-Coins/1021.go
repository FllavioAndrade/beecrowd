package main

import "fmt"

func main() {
	var banknotes = [6]float64{100.00, 50, 20, 10, 5, 2}
	var money float64
	var notas int
	money = 548.97
	fmt.Scan(&money)

	fmt.Print("NOTAS\n")
	for i := 0; i < len(banknotes); i++ {

		notas = int(money) / banknotes[i]
		fmt.Printf("%d nota(s) de %d.00\n", notas, banknotes[i])
		if money >= float64(banknotes[i]) {
			money = float64(banknotes[i]*notas) - money
		}
	}
}
