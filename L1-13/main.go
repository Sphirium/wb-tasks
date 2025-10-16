package main

import "fmt"

func main() {
	a := 100500
	b := 1234567890
	fmt.Printf("Изначально заданные значения: a = %d, b = %d\n", a, b)

	a = a ^ b
	b = b ^ a
	a = a ^ b
	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)

}

// func main() {
// 	a := 10
// 	b := 5
// 	fmt.Printf("Изначально заданные значения: a = %d, b = %d\n", a, b)

// 	a = a + b
// 	b = a - b
// 	a = a - b
// 	fmt.Printf("После обмена: a = %d, b = %d\n", a, b)

// }
