package main

import "fmt"

// contains принимает два слайса интов и возвращает слайс интов с перечениями значений
func contains(nums1, nums2 []int) []int {
	// создаем мапу для хранения первого слайса
	setNums1 := make(map[int]bool)
	for _, v := range nums1 {
		setNums1[v] = true
	}

	var result []int
	// создаем мапу для сравнения элементов двух слайсов
	doublicate := make(map[int]bool)
	for _, v := range nums2 {
		if setNums1[v] && !doublicate[v] {
			result = append(result, v)
			doublicate[v] = true
		}
	}
	return result

}
func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{4, 5, 6, 7, 8}

	result := contains(nums1, nums2)
	fmt.Printf("nums1 = %v\n", nums1)
	fmt.Printf("nums2 = %v\n", nums2)
	fmt.Printf("Пересечение = %v\n", result)
}
