package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(CharacterUnique("abcd"))
	fmt.Println(CharacterUnique("abCdefAaf"))
	fmt.Println(CharacterUnique("aabcd"))
}

// функция проверки
func CharacterUnique(stg string) bool {
	//создаем мапу, от нее нам нужны будут только ключи,
	//по этому чтобы занималось меньше памя в качестве объекта указываем структуру
	stgChar := make(map[string]struct{})

	// переводим строку в нижний регистр, чтобы сделать вализацию независящую от регистра
	for _, char := range strings.ToLower(stg) {
		_, ok := stgChar[string(char)]
		//если ключ уже был добавлен, значит строка не уникальна
		if ok {
			return false
		}
		stgChar[string(char)] = struct{}{}
	}
	//если все символы были добавлены и не сработало условие на возвращение falce, значит строка уникальна
	return true
}
