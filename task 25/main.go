package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Before")
	sleep(2 * time.Second)
	fmt.Println("After")
}

func sleep(timeToSleep time.Duration) {
	// создаем таймер, который отправит сигнал через заданный промежуток времени, пока сигнала нет работает цикл
	timer := time.NewTimer(timeToSleep)
	for {
		select {
		case <-timer.C:
			return
		default:
			continue
		}
	}
}
