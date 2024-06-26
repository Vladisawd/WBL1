package main

import (
	"fmt"
)

func main() {
	array := []int{-2, -1, 100, 0, -34, 1000, 3, 4, 5, 10, 12, -3, -1000, 0, 0}
	fmt.Println(sortArr(array))
}

/*
алгоритм быстрой сортировки реализуется через рекурсию, то есть функция вызывает внутри себя саму себя
для начала мы рассматриваем самый положительный сценарий, это когда наш массив состоит из одного элемента, а значит уже отсортирован
к этому сценарию мы и будем стремиться
далее мы разделяем массив на опорный элемент, массив чисел которые больше опорного и меньше опорного
далее записываем числа в два подмассива меньше и больше опорного и потом вызываем эту же фекцию сортировки для них же
таким образом когда рекурсия достигнет одного элемента в массиве, она начнет обратно сварачиваться и получится отсортированный массив
*/
func sortArr(array []int) []int {
	if len(array) < 2 {
		return array
	}
	var left []int
	var right []int

	for _, val := range array[1:] {
		if val <= array[0] {
			left = append(left, val)
		} else {
			right = append(right, val)
		}
	}
	leftup := append(sortArr(left), array[0])
	sortArray := append(leftup, sortArr(right)...)
	return sortArray
}
