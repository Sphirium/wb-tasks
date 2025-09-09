package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Закрытие горутины по таймауту
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	// Добавил второй контекст, чтобы задать разное время работы горутинам
	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel2()

	defer cancel()

	chanInt := make(chan int)

	go func() {
		for v := 0; v <= 100; v++ {
			select {
			case chanInt <- v:
				time.Sleep(500 * time.Millisecond)
			case <-ctx.Done():
				return
			}

		}
	}()

	go func() {
		for {
			select {
			case data, ok := <-chanInt:
				if !ok {
					return
				}
				fmt.Printf("Получено число %d\n", data)
			case <-ctx2.Done():
				return
			}
		}
	}()
	<-ctx2.Done()

	fmt.Println("Работа завершена по таймауту.")
}
