package main

import "fmt"

func main() {
	//первое множество
	firstMultiplicity := map[string]struct{}{
		"a": {},
		"b": {},
		"c": {},
		"3": {},
		"d": {},
		"e": {},
		"7": {},
		"f": {},
		"g": {},
		"h": {},
		"4": {},
	}

	//второе множество
	secondMultiplicity := map[string]struct{}{
		"1": {},
		"2": {},
		"a": {},
		"3": {},
		"4": {},
		"h": {},
		"5": {},
		"6": {},
		"g": {},
		"7": {},
		"8": {},
	}

	fmt.Println(intersectionMultiplicity(firstMultiplicity, secondMultiplicity))
}

// функция возвращающая множество состоящая из пересечения поступающих множеств
func intersectionMultiplicity(first, second map[string]struct{}) map[string]struct{} {
	iM := make(map[string]struct{})

	//берем ключ первого множества и смотрим имеются ли такие же ключи во втором,
	//если да, то добавляем их в общее множество
	for key := range first {
		if _, ok := second[key]; ok {
			iM[key] = struct{}{}
		}
	}
	return iM
}
