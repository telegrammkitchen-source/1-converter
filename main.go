package main

import "fmt"

const (
	USDtoEUR = 0.93
	USDtoRUB = 93.5
)

func getUserInput() (float64, string, string) {
	var amount float64
	var fromCurrency, toCurrency string

	fmt.Print("Введите сумму для конвертации: ")
	fmt.Scan(&amount)

	fmt.Print("Введите исходную валюту (USD, EUR, RUB): ")
	fmt.Scan(&fromCurrency)

	fmt.Print("Введите целевую валюту (USD, EUR, RUB): ")
	fmt.Scan(&toCurrency)

	return amount, fromCurrency, toCurrency
}

func calculateConversion(amount float64, fromCurrency, toCurrency string) float64 {

	return 0
}

func main() {
	// Выводим курсы валют
	EURtoRUB := (1 / USDtoEUR) * USDtoRUB

	fmt.Printf("Курсы конвертации:\n")
	fmt.Printf("USD to EUR: %.2f\n", USDtoEUR)
	fmt.Printf("USD to RUB: %.2f\n", USDtoRUB)
	fmt.Printf("EUR to RUB: %.2f\n", EURtoRUB)
	fmt.Println()

	// Получаем ввод от пользователя
	amount, fromCurrency, toCurrency := getUserInput()

	// Выводим введенные данные
	fmt.Printf("\nВведенные данные:\n")
	fmt.Printf("Сумма: %.2f\n", amount)
	fmt.Printf("Из валюты: %s\n", fromCurrency)
	fmt.Printf("В валюту: %s\n", toCurrency)

	// Вызываем функцию расчета (пока заглушку)
	result := calculateConversion(amount, fromCurrency, toCurrency)
	fmt.Printf("Результат расчета: %.2f\n", result)
}
