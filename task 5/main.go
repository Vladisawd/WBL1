package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// считываем из консоли сколько времени работать функции
	var timeToStop int
	fmt.Scan(&timeToStop)

	// создаем не буферизированный канал, так как не знаем сколько будет данных
	channel := make(chan string)

	//Инициализируем таймер
	timer := time.NewTimer(time.Duration(timeToStop) * time.Second)

	//запускаем чтение из канала
	go reader(channel)

	//запускаем сценарии остановки программы и записи данных в воркер
	for {
		select {
		//сценарий получения сигнала с консоли и завершения работы
		case <-timer.C:
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

// Функция имитации подачи данных в канал
func puscher() string {
	for {
		return fmt.Sprintf("Aboba %d", rand.Intn(1000))
	}
}

// Функция имитации чтения из канала
func reader(channel chan string) {
	for info := range channel {
		fmt.Printf("Read info: %s\n", info)
	}
}
