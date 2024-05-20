package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	//инициализируем буферизированный канал для чисел
	numsChan := make(chan int, len(nums))

	for _, num := range nums {
		numsChan <- num
	}

	//не забываем закрыть
	close(numsChan)

	//инициализируем буферизированный канал для умноженных чисел
	multiplicationNumsChan := make(chan int, len(nums))

	for num := range numsChan {
		multiplicationNumsChan <- num * 2
	}

	//не забываем закрыть
	close(multiplicationNumsChan)

	//проверяем
	for mNum := range multiplicationNumsChan {
		fmt.Println(mNum)
	}

}
