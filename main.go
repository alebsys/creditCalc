package main

import (
	"fmt"
	"math"
)

var creditSum, creditTerm, creditPercent float64

func main() {

	creditSum, creditTerm, creditPercent = question()
	payMonth := annuityPay(creditSum, creditTerm, creditPercent)
	overPay := totalAmount(payMonth, creditTerm)
	answer(payMonth, overPay)
}

// Функция получает исходные данные у пользователя
func question() (creditSum, creditTerm, creditPercent float64) {
	fmt.Println("Введите сумму кредита, в рублях")
	fmt.Scanf("%f", &creditSum)
	fmt.Println("\n")

	fmt.Println("Введите срок кредита, в месяцах")
	fmt.Scanf("%f", &creditTerm)
	fmt.Println("\n")

	fmt.Println("Введите процент по кредиту, простым числом")
	fmt.Scanf("%f", &creditPercent)
	fmt.Println("\n")

	fmt.Println("-----------------")
	return creditSum, creditTerm, creditPercent
}

// Функция считает аннуитентный платеж
func annuityPay(cs, ct, cp float64) float64 {
	percentMonth := cp / 100 / 12
	x := (math.Pow((1 + percentMonth), ct))
	return cs * (percentMonth + percentMonth/(x-1))
}

// Функция считает общую сумму платежей
func totalAmount(pm, ct float64) float64 {
	return pm * ct
}

// Функция выводит на экран результат
func answer(pm, op float64) {
	fmt.Printf("Ваш эжемесячный платеж равен: %.2f\n", pm)
	fmt.Printf("Общая сумма платежей: %.2f\n", op)
}
