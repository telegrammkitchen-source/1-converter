package main

import (
	"fmt"
	"strings"
)

// Используем map для хранения курсов валют
var exchangeRates = map[string]float64{
	"USD_EUR": 0.93,
	"USD_RUB": 93.5,
	"EUR_USD": 1 / 0.93,
	"EUR_RUB": 93.5 / 0.93,
	"RUB_USD": 1 / 93.5,
	"RUB_EUR": 0.93 / 93.5,
}

// Функция для ввода и проверки валюты
func inputCurrency(prompt string) string {
	var currency string
	validCurrencies := map[string]bool{
		"USD": true,
		"EUR": true,
		"RUB": true,
	}

	for {
		fmt.Print(prompt)
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)

		// Проверяем через map вместо цикла
		if validCurrencies[currency] {
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

	// Используем switch для определения пары валют
	var rate float64
	currencyPair := fromCurrency + "_" + toCurrency

	// Получаем курс из map
	rate = exchangeRates[currencyPair]

	// Вычисляем результат
	result := amount * rate

	return result
}

func showExchangeRates() {
	fmt.Println("\nТЕКУЩИЕ КУРСЫ ВАЛЮТ:")

	// Используем switch для красивого форматирования вывода
	for pair, rate := range exchangeRates {
		currencies := strings.Split(pair, "_")
		from := currencies[0]
		to := currencies[1]

		switch {
		case from == "USD" && to == "EUR":
			fmt.Printf("1 USD = %.2f EUR\n", rate)
		case from == "USD" && to == "RUB":
			fmt.Printf("1 USD = %.2f RUB\n", rate)
		case from == "EUR" && to == "USD":
			fmt.Printf("1 EUR = %.2f USD\n", rate)
		case from == "EUR" && to == "RUB":
			fmt.Printf("1 EUR = %.2f RUB\n", rate)
		case from == "RUB" && to == "USD":
			fmt.Printf("1 RUB = %.6f USD\n", rate)
		case from == "RUB" && to == "EUR":
			fmt.Printf("1 RUB = %.6f EUR\n", rate)
		}
	}
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

		// Выводим результат с использованием switch для форматирования
		fmt.Printf("\nРЕЗУЛЬТАТ КОНВЕРТАЦИИ:\n")

		switch {
		case fromCurrency == "USD" && toCurrency == "EUR":
			fmt.Printf("%.2f USD = %.2f EUR\n", amount, result)
		case fromCurrency == "USD" && toCurrency == "RUB":
			fmt.Printf("%.2f USD = %.2f RUB\n", amount, result)
		case fromCurrency == "EUR" && toCurrency == "USD":
			fmt.Printf("%.2f EUR = %.2f USD\n", amount, result)
		case fromCurrency == "EUR" && toCurrency == "RUB":
			fmt.Printf("%.2f EUR = %.2f RUB\n", amount, result)
		case fromCurrency == "RUB" && toCurrency == "USD":
			fmt.Printf("%.2f RUB = %.2f USD\n", amount, result)
		case fromCurrency == "RUB" && toCurrency == "EUR":
			fmt.Printf("%.2f RUB = %.2f EUR\n", amount, result)
		default:
			fmt.Printf("%.2f %s = %.2f %s\n", amount, fromCurrency, result, toCurrency)
		}

		// Спрашиваем, хочет ли пользователь продолжить
		var choice string
		fmt.Print("\nХотите выполнить еще одну конвертацию? (y/n): ")
		fmt.Scan(&choice)

		// Используем switch для обработки выбора пользователя
		switch strings.ToLower(choice) {
		case "y", "yes", "да", "д":
			fmt.Println()
			continue
		case "n", "no", "нет", "н":
			fmt.Println("Спасибо за использование конвертера валют!")
			return
		default:
			fmt.Println("Неизвестный выбор. Выход из программы.")
			return
		}
	}
}
