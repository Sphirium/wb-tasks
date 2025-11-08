package main

import "fmt"

func reverse(s []rune, left, right int) {
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// reverWords переворачивает порядок слов в строке
func reverWords(str string) string {
	runes := []rune(str)

	// Разворачиваем всю строку
	reverse(runes, 0, len(runes)-1)

	// Разворачиваем каждое слово в строке
	start := 0
	for i := 0; i <= len(runes); i++ {
		if i == len(runes) || runes[i] == ' ' {
			reverse(runes, start, i-1)
			start = i + 1
		}
	}

	return string(runes)
}

func main() {
	words := "dance for life"

	reWords := reverWords(words)

	fmt.Println(reWords)
}
