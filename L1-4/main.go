package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Задаем контекст с явной отменой
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := sync.WaitGroup{}

	// Задаем число воркеров и запускаем
	numWorkers := 5
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(ctx, i, &wg)
	}

	// Инициализуем канал для получения сигналов
	chInputSignal := make(chan os.Signal, 1)
	signal.Notify(chInputSignal, syscall.SIGINT, syscall.SIGTERM)

	<-chInputSignal // вычитываем сигнал Ctrl+C
	fmt.Println("Получен сигнал отключения...")

	cancel() // Явный выход отмены контекста

	// Ждем, пока все воркеры завершат работу
	wg.Wait()
	time.Sleep(2 * time.Second) // Имитация отчетливого завершения
	fmt.Println("Все воркеры остановлены. Сервер отключен.")
}

func worker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Воркер %d получил сигнал отключения...\n", id)
			time.Sleep(250 * time.Millisecond)
			fmt.Printf("Воркер %d остановлен.\n", id)
			return
		default:
			fmt.Printf("Воркер %d работает.\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
