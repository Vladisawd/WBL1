package main

import "fmt"

func main() {
	array := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(createMultiplicity(array))
}

// функция создания множества
func createMultiplicity(array []string) map[string]struct{} {
	multiplicity := make(map[string]struct{})

	//проходимся по массиву строк и добавляем их в качестве ключа в мапу
	//удобнее всего использовать именно структуру в качестве значения так как она практически ничего не весит
	for _, str := range array {
		multiplicity[str] = struct{}{}
	}
	return multiplicity
}
