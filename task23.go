package main

import "fmt"

/*23. Удалить i-ый элемент из слайса.*/
func main() {
	// Первый способ
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	number := 5
	result := delIndex(number, slice)
	fmt.Printf("Слайс после удаления %d\nСлайл который был до удаления НЕ изменился ! %d\n", result, slice)
	// Исходный slice остался [0,1,2,3,4,5,6,7,8,9]
	// result = [0,1,2,3,4,6,7,8,9]

	// Второй способ
	slice2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := sliceChange(number, slice2)
	fmt.Printf("Слайс после удаления %d\nСлайл который был до удаления изменился ! %d\n", res, slice2)
	// Исходный slice2 изменился!
}

// Новый слайс
func delIndex(i int, slice []int) []int {
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:i]...)
	newSlice = append(newSlice, slice[i+1:]...)
	return newSlice
}

// меняем исходный
func sliceChange(i int, slice []int) []int {
	slice = append(slice[:i], slice[i+1:]...)
	return slice
}
