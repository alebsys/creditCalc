package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strings"
	"text/tabwriter"
)

var dosro4kaSum, payMonth float64
var dosro4kaMonth int
var answerDosr string
var dosro4ka = make(map[int]float64)

func main() {

	creditSumFlag := flag.Float64("s", 0, "Сумма кредита, в рублях")
	creditTimeFlag := flag.Float64("t", 0, "Срок кредита, в месяцах")
	creditPercentFlag := flag.Float64("p", 0, "Годовой кредитный процент")
	flag.Parse()

	dosro4kaInfo()

	answer(*creditSumFlag, *creditTimeFlag, *creditPercentFlag)
}

// Функция собирает данные о досрочных платежах
func dosro4kaInfo() map[int]float64 {

	fmt.Print("Будут ли досрочные платежи?(да/нет): ")
	fmt.Scanf("%s", &answerDosr)

	if strings.ToLower(answerDosr) == "да" {

		for strings.ToLower(answerDosr) != "нет" {

			fmt.Print("Введите число месяца: ")
			fmt.Scanf("%d", &dosro4kaMonth)
			fmt.Print("Введите сумму досрочного погашения: ")
			fmt.Scanf("%f", &dosro4kaSum)

			dosro4ka[dosro4kaMonth] = dosro4kaSum

			fmt.Print("Будут ли еще досрочные платежи?(да/нет): ")

			fmt.Scanf("%v", &answerDosr)

		}
	}
	return dosro4ka
}

// Функция выводит на экран результат расчетов
func answer(cs, ct, cp float64) {

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 8, 8, 0, '\t', 0)

	defer w.Flush()

	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s", "Месяцы", "Ежемесячный платеж", "Погашение процентов", "Погашение тела кредита", "Долг на конец месяца")
	fmt.Fprintf(w, "\n %s\t%s\t%s\t%s\t%s", "--------", "--------", "--------", "--------", "--------")

	payMonth := annuityPay(cs, ct, cp)

	s := "Итоговые расчеты:"
	d := 110 // shell width

	fmt.Printf(fmt.Sprintf("%%-%ds", d/2), fmt.Sprintf(fmt.Sprintf("%%%ds\n\n", d/2), s))

	fmt.Printf("Общая сумма переплаты: %.2f\n\n", totalAmount(cs, payMonth, ct))

	for i := 1; cs >= 0; i++ {

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

// Функция считает аннуитентный платеж
func annuityPay(cs, ct, cp float64) float64 {
	percentMonth := cp / 100 / 12
	x := (math.Pow((1 + percentMonth), ct))
	return cs * (percentMonth + percentMonth/(x-1))
}

// Функция считает сумму переплаты
func totalAmount(cs, pm, ct float64) float64 {
	return pm*ct - cs
}
