package main

import (
	"fmt"
	"math"
)

/*
Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
*/
func main() {
	slice := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(tempature(slice))
}

func tempature(slice []float64) map[int][]float64 {
	m := make(map[int][]float64)
	for _, v := range slice {
		group := int(math.Floor(v/10) * 10)
		m[group] = append(m[group], v)
	}
	return m
}
