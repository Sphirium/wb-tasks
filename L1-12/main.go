package main

import "fmt"

func main() {
	// Ряд строковых значений по заданию
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	// Создаём множество с помощью идиомы "map[string]struct{}" - так реализовано множество в Go
	set := make(map[string]struct{})

	// Добавляем все элементы в множество
	for _, word := range sequence {
		set[word] = struct{}{}
	}

	// Преобразуем множество обратно в срез (если нужно)
	uniqueWords := make([]string, 0, len(set))
	for word := range set {
		uniqueWords = append(uniqueWords, word)
	}

	fmt.Println("Множество уникальных слов:", uniqueWords)
}
