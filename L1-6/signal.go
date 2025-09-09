package main

import (
	"fmt"
	"time"
)

func worker(exit chan bool) {
	for {
		select {
		case <-exit:
			fmt.Println("Завершение горутины")
			return
		default:
			fmt.Println("Горутина работает...")
			time.Sleep(5 * time.Second)
		}
	}

}

func main() {
	exit := make(chan bool)

	go worker(exit)

	time.Sleep(10 * time.Second)

	fmt.Println("Отправка сигнала о завершении работы.")
	exit <- true

	<-exit
	fmt.Println("Main-горутина завершила работу.")

}
