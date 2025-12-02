package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {
	fmt.Println("Начало работы")
	sleep(5 * time.Second)
	fmt.Println("Прошло 5 секунд")
}
