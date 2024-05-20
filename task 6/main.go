package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	contextMethod()
}

// первый способ остановить горутину - закрыть канал с которым она взаимодействует
func closeChannel() {
	channel := make(chan string)

	//функция чтения из канала
	go func() {
		for {
			//канал вернет два параметра: значение которое в нем лежит и буллевый параметр который говорит закрыт он или нет.
			//если канал закрыт, горутина прекратит свое выполнение
			value, channelState := <-channel
			if !channelState {
				fmt.Println("Channel close")
				return
			}
			fmt.Println(value)
		}
	}()

	channel <- "Biba"
	channel <- "Boba"

	close(channel)
	time.Sleep(1 * time.Second)
}

// второй способ - так же с помощью канала
func feedIntoTheChannel() {
	channel := make(chan bool)

	go worker(channel)

	//имитируем работу, и побаем параметр через 2 секунды
	time.Sleep(2 * time.Second)
	channel <- true

	//ждем пока горутина закончит работу
	time.Sleep(1 * time.Second)
	fmt.Println("Work is complite")
}

// если в канал подается булевый параметр, он завешает работу горутины
func worker(stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("Channel stop work")
			return
		default:
			fmt.Println("Channel working")
		}
	}
}

// третий способ - с помощью контекста
func contextMethod() {
	//создаем контекст, который автоматически пошлет сигнал через 2 секунды
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	//в конце закрываем
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Channel close")
				return
			default:
				fmt.Println("Channel working")
			}
		}
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Work is complite")
}
