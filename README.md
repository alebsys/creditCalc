# Кредитный калькулятор

## Функции 

* Принимает на вход
  * сумму кредита
  * срок кредита
  * процентную ставку по кредиту
* Отдает 
  * общую сумму
  * сумму переплаты
  * ежемесячный платеж

## Формулы расчета кредита

* Формула аннуитетного платежа

```js
A = K*S, где

A — ежемесячный аннуитетный платеж,

K — коэффициент аннуитета,

S — сумма кредита.
```


* Коэффициент аннуитета

```js
K=i*(1+i)^n/((1+i)^n-1), где 

K — коэффициент аннуитета,

i — месячная процентная ставка по кредиту (= годовая ставка/12 месяцев),

n — количество периодов, в течение которых выплачивается кредит.

```