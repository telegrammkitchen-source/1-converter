package main

import (
	"fmt"
	"strings"
)

const (
	USDtoEUR = 0.93
	USDtoRUB = 93.5
)

// Функция для ввода и проверки валюты
func inputCurrency(prompt string) string {
	var currency string
	validCurrencies := []string{"USD", "EUR", "RUB"}

	for {
		fmt.Print(prompt)
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)

		// Проверяем, что валюта допустима
		isValid := false
		for _, validCurrency := range validCurrencies {
			if currency == validCurrency {
				isValid = true
				break
			}
		}

		if isValid {
			break
		} else {
			fmt.Println("Ошибка: допустимые валюты - USD, EUR, RUB")
		}
	}
	return currency
}

// Функция для ввода и проверки числа
func inputAmount(prompt string) float64 {
	var amount float64
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&amount)
		if err != nil || amount <= 0 {
			fmt.Println("Ошибка: введите положительное число")
			// Очищаем буфер ввода
			var discard string
			fmt.Scanln(&discard)
		} else {
			break
		}
	}
	return amount
}

func getUserInput() (float64, string, string) {
	fmt.Println("=== КОНВЕРТЕР ВАЛЮТ ===")

	// Ввод исходной валюты
	fromCurrency := inputCurrency("Введите исходную валюту (USD, EUR, RUB): ")

	// Ввод суммы
	amount := inputAmount("Введите сумму для конвертации: ")

	// Ввод целевой валюты
	toCurrency := inputCurrency("Введите целевую валюту (USD, EUR, RUB): ")

	return amount, fromCurrency, toCurrency
}

func calculateConversion(amount float64, fromCurrency, toCurrency string) float64 {
	// Если валюты одинаковые, возвращаем ту же сумму
	if fromCurrency == toCurrency {
		return amount
	}

	// Конвертация через USD как базовую валюту
	var amountInUSD float64

	// Конвертируем исходную валюту в USD
	switch fromCurrency {
	case "USD":
		amountInUSD = amount
	case "EUR":
		amountInUSD = amount / USDtoEUR
	case "RUB":
		amountInUSD = amount / USDtoRUB
	}

	// Конвертируем из USD в целевую валюту
	var result float64
	switch toCurrency {
	case "USD":
		result = amountInUSD
	case "EUR":
		result = amountInUSD * USDtoEUR
	case "RUB":
		result = amountInUSD * USDtoRUB
	}

	return result
}

func showExchangeRates() {
	EURtoRUB := (1 / USDtoEUR) * USDtoRUB
	EURtoUSD := 1 / USDtoEUR
	RUBtoUSD := 1 / USDtoRUB
	RUBtoEUR := 1 / EURtoRUB

	fmt.Println("\nТЕКУЩИЕ КУРСЫ ВАЛЮТ:")
	fmt.Printf("1 USD = %.2f EUR\n", USDtoEUR)
	fmt.Printf("1 USD = %.2f RUB\n", USDtoRUB)
	fmt.Printf("1 EUR = %.2f USD\n", EURtoUSD)
	fmt.Printf("1 EUR = %.2f RUB\n", EURtoRUB)
	fmt.Printf("1 RUB = %.6f USD\n", RUBtoUSD)
	fmt.Printf("1 RUB = %.6f EUR\n", RUBtoEUR)
	fmt.Println()
}

func main() {
	// Бесконечный цикл для возможности повторных расчетов
	for {
		// Показываем текущие курсы
		showExchangeRates()

		// Получаем ввод от пользователя
		amount, fromCurrency, toCurrency := getUserInput()

		// Выполняем расчет
		result := calculateConversion(amount, fromCurrency, toCurrency)

		// Выводим результат
		fmt.Printf("\nРЕЗУЛЬТАТ КОНВЕРТАЦИИ:\n")
		fmt.Printf("%.2f %s = %.2f %s\n", amount, fromCurrency, result, toCurrency)

		// Спрашиваем, хочет ли пользователь продолжить
		var choice string
		fmt.Print("\nХотите выполнить еще одну конвертацию? (y/n): ")
		fmt.Scan(&choice)

		if strings.ToLower(choice) != "y" {
			fmt.Println("Спасибо за использование конвертера валют!")
			break
		}
		fmt.Println()
	}
}
