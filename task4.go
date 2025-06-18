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

/*4. Реализовать постоянную запись данных в канал (главный поток). Реализовать
набор из N воркеров, которые читают произвольные данные из канала и
выводят в stdout. Необходима возможность выбора количества воркеров при
старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
способ завершения работы всех воркеров.
*/

func main() {
	// Задаем кол-во воркеров
	workers := flag.Int("w", 2, "workers")
	flag.Parse()
	// ВВОД ДАННЫХ
	data := 0
	var wg sync.WaitGroup
	ch := make(chan int, 2)

	// Обработка сигнала
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-sig
		cancel()
		time.Sleep(time.Millisecond * 100)
		close(ch)
	}()

	// Запуск воркеров которые будут читать из канала
	for i := 0; i < *workers; i++ {
		wg.Add(1)
		go work(ch, &wg, ctx, i)
	}

	// Реализация постоянной записи в главном потоке
	for {
		select {
		case <-ctx.Done():
			log.Println("\nReceived interrupt signal")
			return
		case ch <- data:
			data++
			time.Sleep(1 * time.Second)
		}
	}

}

func work(ch chan int, wg *sync.WaitGroup, ctx context.Context, index int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			log.Printf("worker %d terminated", index)
			return
		case data, ok := <-ch:
			if !ok {
				log.Printf("worker %d terminated", index)
				return
			}
			log.Printf("worker %d received data: %d", index, data)
		}
	}
}
