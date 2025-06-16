package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
)

/*2. Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	// ЗАГРУЖАЕМ ПЕРЕМЕННУЮ ОКРУЖЕНИЯ
	workersStr := os.Getenv("WORKERS")

	// Парсим значение
	workers, err := strconv.Atoi(workersStr)
	if err != nil {
		fmt.Printf("Ошибка парсинга WORKERS: %v,\n", err)
	}

	fmt.Printf("Кол-во workers = %d\n", workers)

	// ВВОД ДАННЫХ
	slice := []int{2, 4, 6, 8, 10}
	chInput := make(chan int)
	chOutput := make(chan int)

	// Записываем данные в канал
	go taskProducer(slice, chInput)
	// Из одного канала передаем в другой конкуретно
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go square(chInput, chOutput, &wg)
	}
	// Конкуретно закрываю и wg.wait
	go func() {
		wg.Wait()
		close(chOutput)
	}()

	for v := range chOutput {
		fmt.Println(v)
	}
}

func taskProducer(slice []int, ch chan int) {
	for _, x := range slice {
		ch <- x
	}
	close(ch)
}

func square(chInput chan int, OutputCh chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range chInput {
		OutputCh <- v * v
	}
}
