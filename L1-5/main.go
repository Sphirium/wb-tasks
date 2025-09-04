package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	const duration = 5 * time.Second

	endTime := time.Now().Add(duration) // Общее время работы

	go func() {
		i := 1
		for time.Now().Before(endTime) { // Временной диапозон работы писателя
			ch <- i
			i++
			time.Sleep(300 * time.Millisecond) // Видимость работы
		}
		close(ch)
	}()

	go func() {
		for data := range ch {
			fmt.Printf("Получено число: %d\n", data)
		}
	}()

	time.Sleep(duration) // Ожидаем завершения таймаута
	fmt.Println("Программа завершена по таймауту.")
}
