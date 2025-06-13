package main

import (
	"fmt"
	"log"
	"sync"
)

/*2. Написать программу, которая конкурентно рассчитает значение квадратов чисел
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func main() {
	// Вводные данные
	slice := make([]int, 0, 100)
	for i := 0; i < 10_000_000; i++ {
		slice = append(slice, i*i)
	}

	// Запускаем ф-ию
	//square(slice)
	fmt.Println(slice)
}

func square(slice []int) {
	// ВВОД ДАННЫХ
	var wg sync.WaitGroup
	numWorkers := 2                     // Ограничим кол-во гоуртин кол-вом ядер, чтобы показать Workerpool
	job := make(chan int, numWorkers*2) // Локальный канал для работы

	// Запускаем отправителя
	go func() {
		for _, num := range slice {
			job <- num
		}
		close(job)
	}()

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			for v := range job {
				res := v * v
				log.Printf("Рабочий#%d Квадрат числа (%d) = %v", index, v, res)
			}
		}(i)
	}
	wg.Wait()
}
