package main

import (
	"fmt"
	"sort"
)

/*17. Реализовать бинарный поиск встроенными методами языка.*/

func main() {
	slice := []int{1, 456, 32445, 111, 3245}
	number := 111
	predict, steps := binary(slice, number)
	fmt.Printf("Predict number:%d\nSteps:%d\n", predict, steps)
}

func binary(slice []int, number int) (int, int) {
	sort.Ints(slice)
	low, high := 0, len(slice)-1
	steps := 0

	for low <= high {
		steps++
		mid := low + (high-low)/2
		midValue := slice[mid]

		if midValue == number {
			return midValue, steps
		} else if midValue > number {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1, steps
}
