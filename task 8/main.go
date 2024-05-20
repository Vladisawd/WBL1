package main

import "fmt"

func main() {
	var num int64 = 5 // 101
	fmt.Println(bitSetting(num, 0, 0))
	fmt.Println(bitSetting(num, 1, 0))
}

// функция принимает в себя само число, какой бит нужно заменить и на что его заменять
func bitSetting(num int64, i, val int) int64 {
	if val == 1 {
		//если нужный бит и так равен 1, то мы его таким и оставляем, а если он не равен 1, то заменяем
		return num | (1 << i)
	}
	return num &^ (1 << i)
}
