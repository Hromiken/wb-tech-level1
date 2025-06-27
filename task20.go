package main

import (
	"fmt"
	"strings"
)

/*20. Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/

func main() {
	str := "Dog Sun Wave Boy"
	fmt.Println("Изначальная строка: ", str)
	x := revStat(str)
	fmt.Println()
	fmt.Println("Готовая строка: ", x)
}

func revStat(str string) string {
	words := strings.Fields(str) // Разбиваем на слова
	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i] // Меняем местами
	}
	return strings.Join(words, " ") // Собираем обратно
}
