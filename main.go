package main

import "fmt"

func main() {
	const (
		USDtoEUR = 0.93
		USDtoRUB = 93.5
	)

	EURtoRUB := (1 / USDtoEUR) * USDtoRUB

	fmt.Printf("Курсы конвертации:\n")
	fmt.Printf("USD to EUR: %.2f\n", USDtoEUR)
	fmt.Printf("USD to RUB: %.2f\n", USDtoRUB)
	fmt.Printf("EUR to RUB: %.2f\n", EURtoRUB)
}
