package main

import (
	"fmt"
	"time"
)

func main() {
	// Закрытие горутины по таймеру (принудительно и небезопасно)
	chanInt := make(chan int)

	const timer = 5 * time.Second

	go func() {
		for v := 0; v <= 100; v++ {
			time.Sleep(300 * time.Millisecond)
			chanInt <- v
		}
	}()

	go func() {
		for data := range chanInt {
			fmt.Printf("Получено число %d\n", data)
		}
		defer close(chanInt)
	}()
	time.Sleep(timer)
	fmt.Println("Работа завершена по таймауту.")
}
