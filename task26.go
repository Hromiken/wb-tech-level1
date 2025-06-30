package main

import (
	"fmt"
	"strings"
)

/*
Разработать программу, которая проверяет, что все символы в строке
уникальные (true — если уникальные, false etc). Функция проверки должна быть
регистронезависимой.
*/
func main() {
	var1 := "abcd"
	var2 := "aabcd"
	var3 := "abCdefAaf"
	var4 := "lmn"
	fmt.Println(chechUnique(var1))
	fmt.Println(chechUnique(var2))
	fmt.Println(chechUnique(var3))
	fmt.Println(chechUnique(var4))
}

func chechUnique(str string) bool {
	str = strings.ToLower(str)
	m := make(map[rune]int)
	for _, v := range str {
		m[v]++
	}
	for _, v := range m {
		if v > 1 {
			return false
		}
	}
	return true
}
