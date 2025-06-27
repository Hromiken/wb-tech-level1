package main

import "fmt"

/*
19. Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.
*/
func main() {
	stroka := "Савва"
	result := reverse(stroka)
	resultGpt := reverseGPT(stroka)
	fmt.Println(result, resultGpt)
}

// Мое решение
func reverse(stroka string) string {
	str := []rune(stroka)
	rev := make([]rune, len(str))
	for i := 0; i < len(str); i++ {
		rev[i] = rune(str[len(str)-i-1])
	}
	return string(rev)
}

// GPT
func reverseGPT(stroka string) string {
	str := []rune(stroka)
	for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i] // Меняем символы местами
	}
	return string(str)
}
