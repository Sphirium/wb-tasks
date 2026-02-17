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

var monthOrder = map[string]int{
	"jan": 1,
	"feb": 2,
	"mar": 3,
	"apr": 4,
	"may": 5,
	"jun": 6,
	"jul": 7,
	"aug": 8,
	"sep": 9,
	"oct": 10,
	"nov": 11,
	"dec": 12,
}

// parseHumanReadable преобразует "2K", "1.5M", "500" в число байт.
func parseHumanReadable(s string) (float64, error) {
	// Удаляем пробелы по краям
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("пустая строка")
	}

	// Пробуем найти суффикс: ищем первую букву СПРАВА
	// "2K" → 'K', "1.5M" → 'M', "500" → нет буквы
	suffixPos := -1
	var suffix rune
	for i := len(s) - 1; i >= 0; i-- {
		c := rune(s[i])
		if (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') {
			suffixPos = i
			suffix = c
			break
		}
	}

	multiplier := 1.0
	numPart := s

	if suffixPos != -1 {
		// Определяем множитель по букве (регистронезависимо)
		// "2K", "2k", "2KB" → все = 1024
		switch strings.ToLower(string(suffix)) {
		case "k":
			multiplier = 1 << 10
		case "m":
			multiplier = 1 << 20
		case "g":
			multiplier = 1 << 30
		case "t":
			multiplier = 1 << 40
		case "p":
			multiplier = 1 << 50
		case "e":
			multiplier = 1 << 60
		default:
			return 0, fmt.Errorf("неизвестный суффикс: %c", suffix)
		}
		numPart = s[:suffixPos] // всё до буквы — число
	}

	// Парсим число
	num, err := strconv.ParseFloat(numPart, 64)
	if err != nil {
		return 0, err
	}

	return num * multiplier, nil
}

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
func isSorted(lines []string, reverse, numeric, blank, month, human bool, key int) bool {
	// Проходим по всем соседним парам строк
	for i := 1; i < len(lines); i++ {
		a := extractField(lines[i-1], key)
		b := extractField(lines[i], key)

		// Приоритет 1: -h (человекочитаемые размеры)
		if human {
			aClean, bClean := a, b
			if blank {
				aClean = strings.TrimFunc(aClean, unicode.IsSpace)
				bClean = strings.TrimFunc(bClean, unicode.IsSpace)
			}
			numA, errA := parseHumanReadable(aClean)
			numB, errB := parseHumanReadable(bClean)
			if errA == nil && errB == nil {
				if reverse {
					if numA < numB { // должно быть убывание
						return false
					}
				} else {
					if numA > numB { // должно быть возрастание
						return false
					}
				}
				continue
			}
			// Если не оба распарсились — переходим к следующему режиму
		}

		// Приоритет 2: -M (месяцы)
		if month {
			aClean := a
			bClean := b
			if blank {
				aClean = strings.TrimFunc(aClean, unicode.IsSpace)
				bClean = strings.TrimFunc(bClean, unicode.IsSpace)
			}

			// Приводим к нижнему регистру для регистронезависимого сравнения
			aKey := strings.ToLower(aClean)
			bKey := strings.ToLower(bClean)
			// Ищем в карте monthOrder: есть ли такие месяцы?
			ordA, okA := monthOrder[aKey]
			ordB, okB := monthOrder[bKey]

			if okA && okB {
				// Оба — валидные месяцы: проверяем порядок
				if reverse {
					if ordA < ordB {
						return false
					}
				} else {
					if ordA > ordB {
						return false
					}
				}
				continue
			}

			// Приоритет 3: -n (числовая)
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

			// Приоритет 4: строковое сравнение (с учётом -b)
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
	}
	return true
}

func main() {
	// Объявляем флаги
	reverse := flag.Bool("r", false, "сортировка в обратном порядке")
	numeric := flag.Bool("n", false, "Числовая сортировка")
	key := flag.Int("k", 0, "Сортировка по N-му полю(разделитель - табуляция)")
	unique := flag.Bool("u", false, "выводить только уникальные строки")
	check := flag.Bool("c", false, "проверить, отсортированы ли данные")
	blank := flag.Bool("b", false, "игнорировать начальные и конечные пробелы при сравнении")
	month := flag.Bool("M", false, "сортировать по названию месяца (Jan, Feb, ..., Dec)")
	human := flag.Bool("h", false, "сортировать по человекочитаемым размерам (2K, 1M, 500G)")

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
		if !isSorted(lines, *reverse, *numeric, *blank, *month, *human, *key) {
			fmt.Fprintln(os.Stderr, "sort: вход не отсортирован")
			os.Exit(1)
		}
		return
	}

	// Сортируем с учётом всех флагов
	sort.SliceStable(lines, func(i, j int) bool {
		a := extractField(lines[i], *key)
		b := extractField(lines[j], *key)

		if *human {
			aClean, bClean := a, b // копируем данные, чтобы не затереть исходники
			if *blank {
				aClean = strings.TrimFunc(aClean, unicode.IsSpace)
				bClean = strings.TrimFunc(bClean, unicode.IsSpace)
			}
			numA, errA := parseHumanReadable(aClean)
			numB, errB := parseHumanReadable(bClean)
			if errA == nil && errB == nil {
				if *reverse {
					return numA > numB
				}
				return numA < numB
			}
		}

		if *month {
			aClean, bClean := a, b // копируем данные, чтобы не затереть исходники
			if *blank {
				aClean = strings.TrimFunc(aClean, unicode.IsSpace)
				bClean = strings.TrimFunc(bClean, unicode.IsSpace)
			}
			aKey := strings.ToLower(aClean)
			bKey := strings.ToLower(bClean)
			ordA, okA := monthOrder[aKey]
			ordB, okB := monthOrder[bKey]

			if okA && okB {
				if *reverse {
					return ordA > ordB
				}
				return ordA < ordB
			}
			// Если не оба месяца — сравниваем как числа (ниже)
		}

		if *numeric {
			numA, errA := strconv.ParseFloat(strings.TrimSpace(a), 64)
			numB, errB := strconv.ParseFloat(strings.TrimSpace(b), 64)

			if errA == nil && errB == nil {
				if *reverse {
					return numA > numB
				}
				return numA < numB
			}
			// Если не оба числа — сравниваем как строки (ниже)
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
