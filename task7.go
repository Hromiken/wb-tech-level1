package main

import (
	"log"
	"sync"
)

/*Реализовать конкуретную запись в map'у*/
func main() {
	m := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup
	ch := make(chan int)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go mapWriter(m, &wg, ch, &mu)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for v := range ch {
		log.Println(v)
	}
}

func mapWriter(m map[int]int, wg *sync.WaitGroup, ch chan int, mu *sync.Mutex) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		mu.Lock()
		m[i] = i * i
		v := m[i]
		mu.Unlock()
		ch <- v
	}
	//	Записывать данные надо под мьютексом в map'у
}
