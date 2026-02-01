package main

import (
	"errors"
	"fmt"
	"os"
	"unicode"
)

func Unpack2(s string) (string, error) {
	if s == "" {
		return "", nil
	}

	runes := []rune(s)
	var result []rune
	i := 0

	for i < len(runes) {
		var char rune
		isEscaped := false // проверка на экранирование

		// Определяем текущий символ
		if runes[i] == '\\' {
			if i+1 >= len(runes) { // проверка на переполнение
				return "", errors.New("Обратный слеш в конце строки")
			}
			char = runes[i+1] // присваиваем символ вида \\5 => '5'
			isEscaped = true  // устанавливаем флаг, что символ был экранирован
			i += 2            // пропускаем два символа
		} else {
			char = runes[i]
			i++
		}

		// Если символ — цифра и не экранирован, это ошибка
		if unicode.IsDigit(char) && !isEscaped {
			return "", errors.New("некорректный формат: цифра без предшествующего символа")
		}

		// Читаем множитель
		count := 1
		if i < len(runes) && unicode.IsDigit(runes[i]) {
			count = 0
			for i < len(runes) && unicode.IsDigit(runes[i]) {
				count = count*10 + int(runes[i]-'0')
				i++
			}
		}
		// Добавляем символ count раз
		for j := 0; j < count; j++ {
			result = append(result, char)
		}
	}

	return string(result), nil
}

func main() {
	input := "ab2c3\\44"
	output, err := Unpack2(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v", err)
	}
	fmt.Println(output)

}
