package main

import (
	"fmt"
	"math"
)

// создадим структуру для самой точки, для более точного определения информации за тип возьмем float, а не int
type Point struct {
	x float64
	y float64
}

// тестим в мейне
func main() {
	p1 := Point{1, 2}
	p2 := Point{3, 4}
	fmt.Printf("Distance: %f", Distance(p1, p2))
}

// конструктор для более удобного использования
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// рассчитываем расстояние
// формула рассчета AB = √((bx - ax)^2 + (by- ay)^2).
func Distance(point1, point2 Point) float64 {
	return math.Sqrt(math.Pow((point2.x-point1.x), 2.0) + math.Pow((point2.y-point1.y), 2.0))
}
