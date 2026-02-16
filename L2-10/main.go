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
	"unicode"
)

// extractField извлекает N-е поле из строки, разделённой табуляцией.
// Нумерация полей начинается с 1 (как в GNU sort).
// Если N <= 0 или поле не существует — возвращается вся строка.
func extractField(line string, n int) string {
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

// isSorted проверяет, отсортирован ли срез строк согласно флагам.
// Возвращает true, если отсортирован.
func isSorted(lines []string, reverse, numeric, blank bool, key int) bool {
	// Проходим по всем соседним парам строк
	for i := 1; i < len(lines); i++ {
		a := extractField(lines[i-1], key)
		b := extractField(lines[i], key)

		// Если числовая сортировка и обе строки — числа
		if numeric {
			numA, errA := strconv.ParseFloat(strings.TrimSpace(a), 64)
			numB, errB := strconv.ParseFloat(strings.TrimSpace(b), 64)
			if errA == nil && errB == nil {
				// В обратном порядке: предыдущее должно быть >= текущего
				// В прямом порядке: предыдущее должно быть <= текущего
				if reverse {
					if numA < numB {
						return false
					}
				} else {
					if numA > numB {
						return false
					}
				}
				// Числа в порядке — переходим к следующей паре
				continue
				// Если не числа — сравниваем как строки (как в sort -n)
			}
		}

		// Обрезаем все пробелы
		if blank {
			a = strings.TrimFunc(a, unicode.IsSpace)
			b = strings.TrimFunc(b, unicode.IsSpace)
		}

		if reverse {
			if a < b {
				return false
			}
		} else {
			if a > b {
				return false
			}
		}

	}
	return true
}

func main() {
	// Объявляем флаг -r (по умолчанию false)
	reverse := flag.Bool("r", false, "сортировка в обратном порядке")
	numeric := flag.Bool("n", false, "Числовая сортировка")
	key := flag.Int("k", 0, "Сортировка по N-му полю(разделитель - табуляция)")
	unique := flag.Bool("u", false, "выводить только уникальные строки")
	check := flag.Bool("c", false, "проверить, отсортированы ли данные")
	blank := flag.Bool("b", false, "игнорировать начальные и конечные пробелы при сравнении")
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

	fmt.Fprintf(os.Stderr, "Прочитано %d строк:\n", len(lines))
	for i, line := range lines {
		fmt.Fprintf(os.Stderr, "  [%d]: %q\n", i, line)
	}

	// Проверям на факт сортировки
	if *check {
		if !isSorted(lines, *reverse, *numeric, *blank, *key) {
			fmt.Fprintln(os.Stderr, "sort: вход не отсортирован")
			os.Exit(1) // Код выхода 1 = ошибка (как в оригинальном sort)
		}
		return
	}

	// Сортируем с учётом флагов -r, -n, -k
	sort.SliceStable(lines, func(i, j int) bool {
		a := extractField(lines[i], *key)
		b := extractField(lines[j], *key)

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

		// Убираем все пробелы
		if *blank {
			a = strings.TrimFunc(a, unicode.IsSpace)
			b = strings.TrimFunc(b, unicode.IsSpace)
		}

		if *reverse {
			return a > b
		}
		return a < b
	})

	// Если задан -u — удаляем дубликаты
	if *unique {
		lines = removeDuplicates(lines)
	}

	// Выводим результат
	for _, line := range lines {
		fmt.Println(line)
	}

}
