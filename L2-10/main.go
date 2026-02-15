package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// extractField извлекает N-е поле из строки, разделённой табуляцией.
// Нумерация полей начинается с 1 (как в GNU sort).
// Если N <= 0 или поле не существует — возвращается вся строка.
func exctractField(line string, n int) string {
	// Если номер поля <= 0 — используем всю строку (поведение по умолчанию)
	if n <= 0 {
		return line
	}

	// Разделяем строку по табуляции
	fields := strings.Split(line, "\t")

	// Проверяем, существует ли запрошенное поле
	if n > len(fields) {
		return ""
	}

	return fields[n-1]
}

// removeDuplicates удаляет соседние дубликаты из ОТСОРТИРОВАННОГО среза.
// Сравнивает ПОЛНЫЕ строки (как в оригинальном sort -u).
func removeDuplicates(lines []string) []string {
	if len(lines) == 0 {
		return lines
	}

	// Результат начинается с первой строки
	result := []string{lines[0]}

	for i := 1; i < len(lines); i++ {
		if lines[i] != lines[i-1] {
			// Если текущая строка НЕ равна предыдущей — добавляем в результат
			result = append(result, lines[i])
		}
		// Если равна — пропускаем (дубликат)
	}
	return result

}

func main() {
	// Объявляем флаг -r (по умолчанию false)
	reverse := flag.Bool("r", false, "сортировка в обратном порядке")
	numeric := flag.Bool("n", false, "Числовая сортировка")
	key := flag.Int("k", 0, "Сортировка по N-му полю(разделитель - табуляция)")
	unique := flag.Bool("u", false, "выводить только уникальные строки")

	// Парсим флаги ДО обращения к os.Args[1]
	flag.Parse()

	// Определяем источник ввода: файл или stdin
	input := os.Stdin
	if flag.NArg() > 0 {
		// Если передан аргумент — это имя файла
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			log.Fatalf("Ошибка чтения данных: %v", err)
		}
		defer file.Close()
		input = file
	}

	// Читаем все строки в срез
	scanner := bufio.NewScanner(input)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Ошибка сканирования данных: %v", err)
	}

	// Сортируем с учётом флагов -r, -n, -k
	sort.SliceStable(lines, func(i, j int) bool {
		a := exctractField(lines[i], *key)
		b := exctractField(lines[j], *key)

		if *numeric {
			numA, errA := strconv.ParseFloat(strings.TrimSpace(a), 64)
			numB, errB := strconv.ParseFloat(strings.TrimSpace(b), 64)

			if errA == nil && errB == nil {
				if *reverse {
					return numA > numB
				}
				return numA < numB
			}
		}

		if *reverse {
			return a > b
		}
		return a < b
	})

	if *unique {
		lines = removeDuplicates(lines)
	}

	// Выводим результат
	for _, line := range lines {
		fmt.Println(line)
	}

}
