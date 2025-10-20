package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	// Сигнал рекурсии, что массив (слайс) отсортирован
	if len(arr) <= 1 {
		return arr
	}

	halfSlice := len(arr) / 2

	// pivot - выбор опорного элемента
	pivot := arr[halfSlice]

	// Создаём три среза:
	var left, middle, right []int

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v == pivot {
			middle = append(middle, v)
		} else {
			right = append(right, v)
		}
	}

	// Сортируем рекурсивно
	leftSorter := quickSort(left)
	rightSorter := quickSort(right)

	// Собираем весь результат
	result := append(leftSorter, middle...)
	result = append(result, rightSorter...)

	return result

}

func main() {
	unsorted := []int{45, 30, 60, 15, 90, 75, 105}
	fmt.Println("Дано: ", unsorted)

	sorted := quickSort(unsorted)
	fmt.Println("Результат: ", sorted)
}
