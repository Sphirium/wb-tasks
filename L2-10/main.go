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

func main() {
	// Объявляем флаг -r (по умолчанию false)
	reverse := flag.Bool("r", false, "сортировка в обратном порядке")
	numeric := flag.Bool("n", false, "Числовая сортировка")
	key := flag.Int("k", 0, "Сортировка по N-му полю(разделитель - табуляция)")

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
	sort.Slice(lines, func(i, j int) bool {
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

	// Выводим результат
	for _, line := range lines {
		fmt.Println(line)
	}

}
