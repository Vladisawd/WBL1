package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("Slice before:%v\n", slice)

	sliceAfter := deleteIndex(slice, 3)
	fmt.Printf("Slice after:%v\n", sliceAfter)
}

// просто принимаем слайс разделяем его на часть до нужного элемента, и после нужного элемента на 1 и возвращаем их соединения
// ... означает что дальше еще есть элементы
func deleteIndex(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
