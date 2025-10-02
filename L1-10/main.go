package main

import (
	"fmt"
	"sort"
)

func main() {
	tempers := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 35.0, 41.2, 47.4, 24.5, -21.0, 32.5, 22.3}

	tempGroups := make(map[int][]float64)

	for _, x := range tempers {

		separator := int(x) / 10 * 10 // приводим float64 к int, чтобы отбросить дробную часть

		tempGroups[separator] = append(tempGroups[separator], x)
	}

	// Сбор и сортировка ключей
	keys := make([]int, 0, len(tempGroups))
	for k := range tempGroups {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	// Вывод результата
	for _, k := range keys {
		fmt.Printf("%d:{", k)
		for i, v := range tempGroups[k] {
			if i > 0 {
				fmt.Print(", ")
			}
			// Сохраняем один десятичный знак для единообразия
			if v == float64(int(v)) {
				fmt.Printf("%.1f", v)
			} else {
				fmt.Print(v)
			}
		}
		fmt.Println("}")
	}

}
