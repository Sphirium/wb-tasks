package main

import "fmt"

func main() {

	numsArr := []int{2, 4, 6, 8, 10}

	ch := make(chan string, len(numsArr))

	for _, number := range numsArr {

		go func(n int) {
			ch <- fmt.Sprintf("Дано: %d, результат: %d", n, n*n)

		}(number)

	}

	for i := 0; i < len(numsArr); i++ {
		fmt.Println(<-ch)
	}
}
