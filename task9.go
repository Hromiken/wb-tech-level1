package main

import (
	"log"
	"sync"
	"time"
)

/*
Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
массива, во второй — результат операции x*2, после чего данные из второго
канала должны выводиться в stdout.
*/
func main() {
	var wg sync.WaitGroup
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Записываем в №1 канал
	go tasker(ch1, &wg)

	//Умножаем числа на 2
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go after(ch1, ch2, &wg)
	}
	// Закрываем каналы
	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
	}()
	// читаем из 2 канала и выводим в stdout
	for v := range ch2 {
		log.Println(v)
	}
}

func tasker(ch chan int, wg *sync.WaitGroup) {
	data := 0
	for {
		ch <- data
		data++
		time.Sleep(time.Second * 1)
	}
}

func after(ch1 chan int, ch2 chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch1 {
		ch2 <- v * 2
	}
}
