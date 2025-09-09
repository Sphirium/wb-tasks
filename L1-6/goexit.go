package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("### Запуск имитации работы горутины ###")

	go func() {
		fmt.Println(`
		============Горутина: начало работы=============
		`)
		for v := range 10 {
			time.Sleep(500 * time.Millisecond)
			v++
			fmt.Println("Горутина: вывожу число", v)
			if v == 7 { // какое-то условие
				fmt.Println("Горутина: условие выполнено -  получено число", v)
				runtime.Goexit()
			}
		}

	}()
	time.Sleep(7 * time.Second)

	fmt.Println("Основная программа завершена.")
}
