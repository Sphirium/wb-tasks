package main

import (
	"errors"
	"fmt"
	"os"
	"unicode"
)

// Unpack распаковывает строку с поддержкой escape-последовательностей.
// Примеры:
//
//	"a4bc2d5e"     → "aaaabccddddde"
//	"qwe\\4\\5"    → "qwe45"
//	"qwe\\45"      → "qwe44444"
//	"5a"           → ошибка
//	""             → ""
func Unpack(data string) (string, error) {
	if data == "" {
		return "", nil
	}

	runes := []rune(data)
	var result []rune
	i := 0

	for i < len(runes) {
		var char rune
		isEscaped := false

		// Определяем текущий символ
		if runes[i] == '\\' {
			if i+1 >= len(runes) {
				return "", errors.New("обратный слеш в конце строки")
			}
			char = runes[i+1]
			isEscaped = true
			i += 2
		} else {
			char = runes[i]
			i++
		}

		// Если символ — цифра и не экранирован, это ошибка
		if unicode.IsDigit(char) && !isEscaped {
			return "", errors.New("некорректный формат: цифра без предшествующего символа")
		}

		// Читаем множитель
		count := 1 // присваиваем стандартное число
		if i < len(runes) && unicode.IsDigit(runes[i]) {
			count = 0 // сбрасываем стандартное число на 0, чтобы собрать новое число
			for i < len(runes) && unicode.IsDigit(runes[i]) {
				count = count*10 + int(runes[i]-'0') // стандартный парсинг числа
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
	input := "d2\\5\\4g3"
	output, err := Unpack(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(output)

}
