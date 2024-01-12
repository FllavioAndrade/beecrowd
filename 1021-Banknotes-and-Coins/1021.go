package main

import "fmt"

func main() {
	var banknotes = [6]int{100, 50, 20, 10, 5, 2}
	var bankcoins = [6]float64{100.00, 50.00, 25.00, 10.00, 5.00, 1.00}

	var money float64
	var notas int
	fmt.Scan(&money)

	fmt.Print("NOTAS:\n")
	for i := 0; i < len(banknotes); i++ {

		notas = int(money) / banknotes[i]
		fmt.Printf("%d nota(s) de R$ %d.00\n", notas, (banknotes[i]))
		if money >= float64(banknotes[i]) {
			money = money - float64(banknotes[i]*notas)
		}
	}
	money *= 100.01
	fmt.Print("MOEDAS:\n")

	for i := 0; i < len(banknotes); i++ {

		notas = int(money) / int(bankcoins[i])
		fmt.Printf("%d moeda(s) de R$ %.2f\n", notas, (bankcoins[i] / 100))
		if money >= float64(bankcoins[i]) {
			money = money - (bankcoins[i] * float64(notas))
		}
	}

}
