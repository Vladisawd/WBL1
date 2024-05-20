package main

import "fmt"

func main() {
	initialData := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	sortingDegrees(initialData)
}

func sortingDegrees(initialData []float32) {
	degreeGroups := make(map[int][]float32)

	for _, val := range initialData {
		degreeGroups[(int(val)/10)*10] = append(degreeGroups[(int(val)/10)*10], val)
	}

	fmt.Printf("Вывод просто мапы: %v \n", degreeGroups)
	for key, val := range degreeGroups {
		fmt.Printf("Группа: %d, значения: %v\n", key, val)
	}
}
