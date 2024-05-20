package main

import (
	"fmt"
	"strings"
)

//данный фрагмент кода создает большую строку и записывает переменную
//первая проблема, это отсутствие информации о функции createHugeString,
//но допустим в функции просто повторяется какой-то набор символов и создается строка по получаемым параметрам
//вторая проблема, заключается в определении одной и той же строки два раза, вметсо того, чтобы просто ссылаться на нее
/*
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}

func main() {
  someFunc()
}
*/

func main() {
	v := createHugeString(1 << 10)
	//строку мы не инициализируем, а ссылаемся на нее чтобы не выделять дополнительно памяти
	justString := strings.Clone(v[:100])
	fmt.Println(justString)
}

// как бы могла выглядеть функция createHugeString
func createHugeString(len int) string {
	result := ""
	for len > 0 {
		result += "A"
		len--
	}
	return result
}
