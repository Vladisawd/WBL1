package main

import (
	"fmt"
	"strings"
)

func main() {
	normString := "салфетка"
	fmt.Println(stringReversal(normString))
}

/*
Создаем стринг билдер который будет конкотинировать строку
переводим строку в тип рун, чтобы корректно обращаться к элементам строки
запускаем цикл в котором записываем в начало перемернутой строки элементы с конца нормальной строки
*/
func stringReversal(normString string) string {
	buildstr := strings.Builder{}
	stringToRune := []rune(normString)
	lastElem := len(stringToRune) - 1
	for lastElem >= 0 {
		buildstr.WriteRune(stringToRune[lastElem])
		lastElem--
	}
	return buildstr.String()
}
