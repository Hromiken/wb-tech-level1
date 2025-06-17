package main

import (
	"flag"
	"fmt"
	"sync"
)

/*
3. Дана последовательность чисел: 2,4,6,8,10. Найти сумму их
квадратов с использованием конкурентных вычислений.
*/
func main() {
	// ВВОД ДАННЫХ
	workers := flag.Int("w", 2, "Number of workers")
	flag.Parse()
	slice := []int{2, 4, 6, 8, 10, 12}
	chIn := make(chan int)
	chOut := make(chan int)

	//Снимаем данные со слайса в канал
	go Producer(chIn, slice)

	//Находим конкуретно квадрат
	var wg sync.WaitGroup
	for i := 1; i <= *workers; i++ {
		wg.Add(1)
		go sumSquare(chIn, chOut, &wg, i)
	}
	// Конкуретно закрываем
	totalSum := 0
	go func() {
		for sum := range chOut {
			totalSum += sum
		}
	}()

	// Ждем завершения
	wg.Wait()
	close(chOut)

	// Выводим общую сумму
	fmt.Printf("Общая сумма квадратов = %d\n", totalSum)
}

func Producer(ch chan int, slice []int) {
	for _, v := range slice {
		ch <- v * v
	}
	close(ch)
}

func sumSquare(chIn chan int, chOut chan int, group *sync.WaitGroup, index int) {
	var sum int
	defer group.Done()
	for v := range chIn {
		sum += v
	}
	chOut <- sum

}
