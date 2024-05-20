package main

import "fmt"

// есть несколько способов чтобы поменять местами переменные
func main() {
	a := 1
	b := 2
	fmt.Printf("Переменные до изменения a: %d, b: %d\n", a, b)
	//первый способ самый простой
	a, b = b, a
	fmt.Printf("Переменные после изменения a: %d, b: %d\n", a, b)
	//второй способ не разумный но поставленную задачу выполняет
	b = a + b
	a = b - a
	b = b - a
	fmt.Printf("Переменные после изменения, возвращаются назад a: %d, b: %d\n", a, b)
}
