package main

import (
	"flag"
	"log"
	"sync"
	"time"
)

/*
Разработать программу, которая будет последовательно отправлять значения в
канал, а с другой стороны канала — читать. По истечению N секунд программа
должна завершаться
*/
func main() {
	//ВВОД ДАННЫХ
	data := 0
	var wg sync.WaitGroup
	ch := make(chan int)
	x := flag.Int("x", 10, "timeX")
	flag.Parse()

	// Рабочие
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go workX(ch, i, &wg)
	}

	// Ограничиваем время жизни программы
	timer := time.NewTimer(time.Duration(*x) * time.Second)
	defer timer.Stop()

	// Запись в канал с таймером
	for {
		select {
		case <-timer.C:
			log.Printf("timeout")
			time.Sleep(100 * time.Millisecond)
			close(ch)
			wg.Wait()
			log.Printf("done")
			return
		default:
			ch <- data
			data++
			time.Sleep(time.Second * 1)
		}
	}
}

func workX(ch chan int, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range ch {
		log.Printf("Рабочий#%d получил:%d", index, v)
	}
}
