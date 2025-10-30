package main

import (
	"fmt"
	"sort"
)

func BinarySearch(sortedSlice []int, target int) int {
	i := sort.SearchInts(sortedSlice, target)
	if i < len(sortedSlice) && sortedSlice[i] == target {
		return i
	}
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	result := BinarySearch(arr, 9)
	fmt.Println(result)

	result2 := BinarySearch(arr, 15)
	fmt.Println(result2)
}
