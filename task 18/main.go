package main

import (
	"fmt"
	"sync"
	"time"
)

// создаем сам счетчик, чтобы он функционировал в конкурентной среде нужен мьютекс
type ConcurrentCounter struct {
	tick  int
	mutex sync.Mutex
}

// функция инкрементации счетчика
func (c *ConcurrentCounter) increment() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.tick++
}

func main() {
	//инифиализируем счетчик
	count := ConcurrentCounter{}

	timer := time.NewTimer(time.Duration(3) * time.Second)

	//имитируем работу функции
	go func() {
		for {
			select {

			case <-timer.C:
				fmt.Println("\nThe program exits. Wait...")
				fmt.Printf("Counter value: %d\n", count.tick)
				return
			default:
				fmt.Println("Work")

				time.Sleep(100 * time.Millisecond)
				count.increment()
			}
		}
	}()

	time.Sleep(5 * time.Second)
}
