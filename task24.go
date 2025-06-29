package main

import (
	"fmt"
	"math"
)

/*
24. Разработать программу нахождения расстояния между двумя точками, которые
представлены в виде структуры Point с инкапсулированными параметрами x,y и
конструктором.
*/
type Point struct {
	x float64
	y float64
}

func main() {
	pos1 := NewPoint(1, 2)
	pos2 := NewPoint(4, 6)
	res := pos2.Distance(pos1)
	fmt.Println(res)
}

// Конструктор
func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

// Вычисление расстояния между точками  √(dx² + dy²)
func (p Point) Distance(other Point) float64 {
	dx := other.x - p.x
	dy := other.y - p.y
	return math.Sqrt(math.Pow(dx, 2) + math.Pow(dy, 2))
}
