package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"
)

func main() {
	// Задаем кол-во воркеров и таймер
	workers := flag.Int("w", 2, "Number of workers")
	timeX := flag.Int("x", 3, "Time to X")
	flag.Parse()

	// ВВОД ДАННЫХ
	var wg sync.WaitGroup
	chOut := make(chan int)
	chIn := make(chan int)

	// Обработка сигнала
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*timeX-1)*time.Second)

	go generate(chOut, ctx)
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			makeSquare(chOut, chIn)
		}()
	}
	// Запускаем горутину для закрытия chIn после завершения всех воркеров
	go func() {
		wg.Wait()
		close(chIn)
	}()

	// Обработка прерывания
	go func() {
		select {
		case <-sig:
			cancel()
			wg.Wait()
			log.Println("Received interrupt signal")
		case <-ctx.Done():
			log.Println("Time is over")
		}
	}()

	result(chIn)
}

func generate(ch chan<- int, ctx context.Context) {
	data := 0
	defer close(ch)
	for {
		select {
		case ch <- data:
			data++
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			return
		}
	}
}

func makeSquare(chOut chan int, chIn chan int) {
	for v := range chOut {
		chIn <- v * v
	}
}

func result(ch <-chan int) {
	for v := range ch {
		log.Printf("Value %d\n", v)
	}
}
