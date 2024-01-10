package main

import "fmt"

func main() {

	var money, notes int
	var banknotes = [7]int{100, 50, 20, 10, 5, 2, 1}

	fmt.Scan(&money)

	fmt.Printf("%d\n", money)
	for i := 0; i < len(banknotes); i++ {
		notes = int(money / banknotes[i])
		money = money % banknotes[i]
		fmt.Printf("%d nota(s) de R$ %d,00\n", notes, banknotes[i])
	}

}
