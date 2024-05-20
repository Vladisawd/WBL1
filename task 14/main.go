package main

import (
	"fmt"
	"reflect"
)

// есть несколько способов определить переменную
func main() {
	//создадим интерфейс с переменными разных типов
	data := []interface{}{"Biba", 123, true, make(chan int), make(map[int]int)}

	//первый способ это просто пройтись по элементам и используя пакет рефлект просто вывести тип
	fmt.Println("First way")
	for _, val := range data {
		fmt.Printf("Variable: %v, matches the type:%v\n", val, reflect.TypeOf(val))
	}

	//второй способ, это просто вывести тип через спецификатор
	fmt.Println("Second way")
	for _, val := range data {
		fmt.Printf("Variable: %v, matches the type:%T\n", val, val)
	}

	fmt.Println("Third way")
	fmt.Println(typeOf("a"))
	fmt.Println(typeOf(1))
	fmt.Println(typeOf(1.2))
}

// третий способ так же с использованием пакета рефлект, если нам нужно определять только определенные типы, а остальные отсеивать
// можно написать следующую функцию
func typeOf(data interface{}) string {

	switch reflect.ValueOf(data).Kind() {
	case reflect.String:
		return fmt.Sprintf("Variable: %s, matches the type string", data)
	case reflect.Bool:
		return fmt.Sprintf("Variable: %t, matches the type bool", data)
	case reflect.Int:
		return fmt.Sprintf("Variable: %d, matches the type int", data)
	case reflect.Float64:
		return fmt.Sprintf("Variable: %g, matches the type float64", data)
	}

	return "Unknown type"
}
