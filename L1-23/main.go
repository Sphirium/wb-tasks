package main

import "fmt"

func cutFromArr(sl []int, i int) []int {
	if i < 0 || i >= len(sl) {
		return sl
	}
	copy(sl[i:], sl[i+1:])

	return sl[:len(sl)-1]

}

func main() {
	sl := []int{0, 1, 2, 3, 4}

	sl = cutFromArr(sl, 3)

	fmt.Println(sl)

}
