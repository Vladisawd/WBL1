package main

import (
	"fmt"
	"strings"
)

func main() {
	normString := "snow dog sun"
	fmt.Println(wordReversal(normString))
}

/*
Сначала используем функции пакета strings и делим поступающую строку по пробелам на слова
далее собираем эти слова с конца и добавляем пробел, но если элемент последний то пробел не добавляем
*/
func wordReversal(normString string) string {
	buildstr := strings.Builder{}

	words := strings.Split(normString, " ")

	lastElem := len(words) - 1

	for lastElem >= 0 {
		if lastElem == 0 {
			buildstr.WriteString(words[lastElem])
		} else {
			buildstr.WriteString(words[lastElem] + " ")
		}
		lastElem--
	}
	return buildstr.String()
}
