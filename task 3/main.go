package main

import (
	"fmt"
	"sync"
)

func main() {
	array := []int{2, 4, 6, 8, 10}
	squareOfTheNumber(array)
}

// создаем функцию для конкурентного вычесления квадратов
func squareOfTheNumber(array []int) {
	var sqSum int
	//объявляем вейт группу, чтобы мейн подождал выполнения всех горутин
	wg := sync.WaitGroup{}
	//даем понять вейт группе сколько будет горутин, их количество понятное дело равно количеству элементов массива
	wg.Add(len(array))
	for _, value := range array {
		go func(value int) {
			sqValue := value * value
			//придавляем новое значение к имеющимся
			sqSum += sqValue
			//даем понять вейт группе что горутина закончила выполнение
			wg.Done()
		}(value)
	}
	//ждем пока закончат выпонения все горутины
	wg.Wait()
	fmt.Printf("The sum of the squares is: %v", sqSum)
}
