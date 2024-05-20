package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//считываем из консоли нужное количество воркеров
	var numOfWorkers int
	fmt.Scan(&numOfWorkers)

	//создаем не буферизированный канал, так как не знаем сколько будет данных
	channel := make(chan string)

	//контекст для прерывания работы по сигналу системы
	//данный паттерн нужен для того, чтобы завершить работы системы по определнному сценарию который будет написан здесь
	// syscall.SIGINT, syscall.SIGTERM - отвечают за завершение программы
	ctx, closeFunc := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer closeFunc()

	//запускаем воркеры
	for wor := 1; wor <= numOfWorkers; wor++ {
		go worker(wor, channel)
	}

	//запускаем сценарии остановки программы и записи данных в воркер
	for {
		select {
		//сценарий получения сигнала с консоли и завершения работы
		case <-ctx.Done():
			fmt.Println("\nThe program exits. Wait...")

			//Имитируем завершение работы
			time.Sleep(1 * time.Second)

			close(channel)
			return

			//сценарий передачи данных в канал
		default:
			info := puscher()
			channel <- info

			//ожидание, чтобы можно было нормально прочитать данные
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// создаем функцию для чтения инфы воркерами с канала
func worker(wor int, channel chan string) {
	for information := range channel {
		fmt.Printf("Worker: %d received information: %s\n", wor, information)
	}
}

// Функция имитации подачи данных в канал
func puscher() string {
	for {
		return fmt.Sprintf("Aboba %d", rand.Intn(1000))
	}
}
