package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {

	dataCh := make(chan int)

	// Проверяем, передано ли количество воркеров
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Не указано число воркеров\n")
		os.Exit(1)
	}

	// Парсим количество воркеров
	numWorkers, err := strconv.Atoi(os.Args[1])
	if err != nil || numWorkers <= 0 {
		fmt.Fprintln(os.Stderr, "Ошибка: укажите положительное число воркеров")
		os.Exit(1)
	}

	// Запуск воркеров
	for i := 1; i <= numWorkers; i++ {
		go worker(i, dataCh)
	}

	// Потоковая запись данных в канал
	counter := 0
	for {
		dataCh <- counter
		counter++
		time.Sleep(1 * time.Second)
	}

}

// Чтение и вывод данных
func worker(id int, dataCh <-chan int) {
	for data := range dataCh {
		fmt.Printf("Воркер %d получил: %d\n", id, data)
	}

}
