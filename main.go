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

// Функция для ввода и проверки валюты (принимает указатель на map)
func inputCurrency(prompt string, validCurrencies *map[string]bool) string {
	var currency string

	for {
		fmt.Print(prompt)
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)

		// Проверяем через указатель на map
		if (*validCurrencies)[currency] {
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

// Функция получает ввод пользователя (принимает указатель на map валидных валют)
func getUserInput(validCurrencies *map[string]bool) (float64, string, string) {
	fmt.Println("=== КОНВЕРТЕР ВАЛЮТ ===")

	// Ввод исходной валюты с передачей указателя на map
	fromCurrency := inputCurrency("Введите исходную валюту (USD, EUR, RUB): ", validCurrencies)

	// Ввод суммы
	amount := inputAmount("Введите сумму для конвертации: ")

	// Ввод целевой валюты с передачей указателя на map
	toCurrency := inputCurrency("Введите целевую валюту (USD, EUR, RUB): ", validCurrencies)

	return amount, fromCurrency, toCurrency
}

// Функция расчета конвертации (принимает указатель на map курсов)
func calculateConversion(amount float64, fromCurrency, toCurrency string, rates *map[string]float64) float64 {
	// Если валюты одинаковые, возвращаем ту же сумму
	if fromCurrency == toCurrency {
		return amount
	}

	// Используем switch для определения пары валют
	var rate float64
	currencyPair := fromCurrency + "_" + toCurrency

	// Получаем курс через указатель на map
	rate = (*rates)[currencyPair]

	// Вычисляем результат
	result := amount * rate

	return result
}

// Функция показа курсов валют (принимает указатель на map курсов)
func showExchangeRates(rates *map[string]float64) {
	fmt.Println("\nТЕКУЩИЕ КУРСЫ ВАЛЮТ:")

	// Используем switch для красивого форматирования вывода
	for pair, rate := range *rates {
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

// Функция вывода результата (оптимизирована для избежания дублирования кода)
func printConversionResult(amount float64, fromCurrency, toCurrency string, result float64) {
	fmt.Printf("\nРЕЗУЛЬТАТ КОНВЕРТАЦИИ:\n")

	// Упрощенный вывод без избыточного switch
	fmt.Printf("%.2f %s = %.2f %s\n", amount, fromCurrency, result, toCurrency)
}

func main() {
	// Создаем map валидных валют один раз
	validCurrencies := map[string]bool{
		"USD": true,
		"EUR": true,
		"RUB": true,
	}

	// Бесконечный цикл для возможности повторных расчетов
	for {
		// Показываем текущие курсы (передаем указатель на map)
		showExchangeRates(&exchangeRates)

		// Получаем ввод от пользователя (передаем указатель на map валидных валют)
		amount, fromCurrency, toCurrency := getUserInput(&validCurrencies)

		// Выполняем расчет (передаем указатель на map курсов)
		result := calculateConversion(amount, fromCurrency, toCurrency, &exchangeRates)

		// Выводим результат
		printConversionResult(amount, fromCurrency, toCurrency, result)

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
