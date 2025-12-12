package main

import (
	"strings"
)

func CharmUniqueDetector(s string) bool {
	s = strings.ToLower(s)
	seen := make(map[rune]bool)
	for _, charm := range s {
		if seen[charm] {
			return false
		}
		seen[charm] = true
	}
	return true
}

func main() {
	symbols := "asdS"
	result := CharmUniqueDetector(symbols)
	println(result)
}
