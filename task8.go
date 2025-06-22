package main

import "fmt"

/*
Дана переменная int64. Разработать программу, которая устанавливает i-й бит в
1 или 0.
*/
func main() {
	var number int64 = 12 //
	var position uint = 3
	var bitValue int = 1 // 1 or 0
	setBit(number, position, bitValue)

}

func setBit(number int64, pos uint, bitValue int) {
	var answer int64
	if bitValue == 1 {
		answer = number | (1 << pos)
	} else {
		answer = number &^ (1 << pos)
	}
	fmt.Println(answer)
}
