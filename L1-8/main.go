package main

import "fmt"

// setBit устанавливает i-й бит числа n в значение bit (0 или 1)
func setBit(n int64, i int, bit int) int64 {
	if bit == 1 {
		// Устанавливаем i-й бит в 1: n | (1 << i)
		return n | (1 << i)
	} else {
		// Устанавливаем i-й бит в 0: n &^ (1 << i) — это битовая операция "и-не"
		return n &^ (1 << i)
	}
}

func main() {
	var num int64 = 6
	i := 1
	value := 0

	result := setBit(num, i, value)

	fmt.Printf("Число %d (бинарно: %b)\n", num, num)
	fmt.Printf("Установка бита %d в %d\n", i, value)
	fmt.Printf("Результат: %d (бинарно: %b)\n", result, result)
}
