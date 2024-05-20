package main

import (
	"errors"
	"fmt"
)

func main() {
	num := 111
	array := []int{-2, -1, 0, 3, 4, 10, 35, 60, 70, 98, 100, 111, 1000, 123456}
	numOfNumber, err := bynorySort(num, array)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(numOfNumber)
}

/*
бинарный поиск предпологает поиск элемента по отсортированному массиву
мы берем средний элемент массива и смотрим больше или меньше он нужного нам, если он больше то ищем в левой стороне, если меньше то в правой
далее когда берем левую или правую сторону мы так же берем ее центр и таким образом через n количество шагов центром станет нужный нам элемент
и мы узнаем его местолопожение в массиве
*/
func bynorySort(num int, array []int) (int, error) {
	min := 0
	max := len(array) - 1
	for min <= max {
		mid := int((min + max) / 2)
		switch {
		case array[mid] < num:
			min = mid + 1
		case array[mid] > num:
			max = mid - 1
		default:
			return mid, nil
		}
	}
	return 0, errors.New("the element is not in the array")
}
