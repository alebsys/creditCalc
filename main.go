package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"text/tabwriter"
)

var creditSum, creditTerm, creditPercent, dosro4kaSum float64
var dosro4kaMonth int
var answerDosr string
var dosro4ka = make(map[int]float64)

func main() {

	creditSumPtr := flag.Float64("sum", 0, "Credit sum")
	creditTermPtr := flag.Float64("period", 0, "Credit period in month")
	creditPercentPtr := flag.Float64("percent", 0, "Annual credit interest")

	flag.Parse()

	fmt.Println("Будет ли досрочка?(да/нет)")
	fmt.Scanf("%v", &answerDosr)

	if answerDosr == "да" {

		for answerDosr != "нет" {

			fmt.Println("Введите число месяца")
			fmt.Scanf("%d", &dosro4kaMonth)
			fmt.Println("и сумму досрочного погашения")
			fmt.Scanf("%f", &dosro4kaSum)

			dosro4ka[dosro4kaMonth] = dosro4kaSum

			fmt.Println("Будут ли еще досрочные платежи?(да/нет)")

			fmt.Scanf("%v", &answerDosr)

		}
	}

	answer(*creditSumPtr, *creditTermPtr, *creditPercentPtr)

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
func answer(cs, ct, cp float64) {

	w := new(tabwriter.Writer)

	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s", "Месяцы", "Ежемесячный платеж", "Погашение процентов", "Погашение тела кредита", "Долг на конец месяца")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s", "--------", "--------", "--------", "--------", "--------")

	// dosro4ka[4] = 10000

	payMonth := annuityPay(cs, ct, cp)

	for i := 1; cs >= 0; i++ {

		// overPay := totalAmount(creditSum, payMonth, creditTerm)

		percAnnuity := cs * cp / 100 / 12
		payTelo := payMonth - percAnnuity

		if dosro4ka[i] > 0 {
			cs = cs - payTelo - dosro4ka[i]
			payMonth = annuityPay(cs, ct-float64(i), cp)
		} else {
			cs = cs - payTelo
		}

		fmt.Fprintf(w, "\n %d\t%.2f\t%.2f\t%.2f\t%.2f\t", i, payMonth, percAnnuity, payTelo, cs)
	}
}
