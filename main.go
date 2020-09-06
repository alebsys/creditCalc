package main

import (
	"fmt"
	"math"
	"os"
	"text/tabwriter"
)

var creditSum, creditTerm, creditPercent float64

func main() {

	creditSum, creditTerm, creditPercent = question()
	payMonth := annuityPay(creditSum, creditTerm, creditPercent)
	overPay := totalAmount(creditSum, payMonth, creditTerm)
	answer(payMonth, overPay, creditSum, creditTerm, creditPercent)

}

// Функция получает исходные данные у пользователя
func question() (creditSum, creditTerm, creditPercent float64) {
	fmt.Println("Введите сумму кредита, в рублях:")
	fmt.Scanf("%f", &creditSum)
	fmt.Println("\n")

	fmt.Println("Введите срок кредита, в месяцах:")
	fmt.Scanf("%f", &creditTerm)
	fmt.Println("\n")

	fmt.Println("Введите процент по кредиту, простым числом:")
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

// Функция высчитывает сумму переплаты
func totalAmount(cs, pm, ct float64) float64 {
	return pm*ct - cs
}

// Функция выводит на экран результат
func answer(pm, op, cs, ct, cp float64) {

	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s", "Месяцы", "Ежемесячный платеж", "Погашение процентов", "Погашение тела кредита", "Долг на конец месяца")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s", "--------", "--------", "--------", "--------", "--------")

	for i := 1; i <= int(ct); i++ {
		percAnnuity := creditSum * creditPercent / 100 / 12
		payTelo := pm - percAnnuity
		creditSum = creditSum - payTelo

		fmt.Fprintf(w, "\n %d\t%.0f\t%.0f\t%.0f\t%.0f\t", i, pm, percAnnuity, payTelo, creditSum)
	}
}
