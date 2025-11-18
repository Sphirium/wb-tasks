package main

import (
	"fmt"
	"math/big"
)

func main() {
	// Задаем числа, превышающие 2^20 (1048576)
	// Для демонстрации используем строки, чтобы быть уверенными в точности
	aStr := "1234567890123456789012345678901234567890"
	bStr := "9876543210987654321098765432109876543210"

	// Инициализируем big.Int из строк
	a := new(big.Int)
	a.SetString(aStr, 10)

	b := new(big.Int)
	b.SetString(bStr, 10)

	// Создаем переменные для результатов операций
	sum := new(big.Int)
	difference := new(big.Int)
	product := new(big.Int)
	quotient := new(big.Int)
	remainder := new(big.Int)

	// Сложение
	sum.Add(a, b)
	fmt.Printf("Сложение: %s + %s = %s\n", a, b, sum)

	// Вычитание
	difference.Sub(a, b)
	fmt.Printf("Вычитание: %s - %s = %s\n", a, b, difference)

	// Умножение
	product.Mul(a, b)
	fmt.Printf("Умножение: %s * %s = %s\n", a, b, product)

	// Деление (частное)
	quotient.Div(a, b)
	remainder.Mod(a, b) // остаток от деления
	fmt.Printf("Деление: %s / %s = %s (остаток: %s)\n", a, b, quotient, remainder)

}
