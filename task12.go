package main

import (
	"fmt"
)

/*
12. Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
собственное множество.
*/
func main() {
	slice := []string{"cat", "dog", "dog", "sea", "dog", "man", "tam", "tam"}
	answer := spisok(slice)
	fmt.Println(answer)
}

func spisok(slice []string) []string {
	result := make([]string, 0, len(slice))
	m := make(map[string]int)
	for _, v := range slice {
		m[v] += 1
	}
	for k := range m {
		result = append(result, k)
	}
	return result
}
