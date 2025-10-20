package main

import (
	"fmt"
	"strings"
)

var justString string

func createHugeString(n int) string {
	return strings.Repeat("0_o ", n)
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString = strings.Clone(v[:100]) // делаем явное копирование данных, избегая утечку памяти
	fmt.Println(justString)
}

func main() {
	someFunc()
}
