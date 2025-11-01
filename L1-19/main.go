package main

import (
	"fmt"
)

func main() {
	word := "РЕШИЛА"
	reWord := reverseWord(word)
	fmt.Println(reWord)
}

func reverseWord(word string) string {
	runesWord := []rune(word)
	for x, y := 0, len(runesWord)-1; x < y; x, y = x+1, y-1 {
		runesWord[x], runesWord[y] = runesWord[y], runesWord[x]
	}
	return string(runesWord)

}
