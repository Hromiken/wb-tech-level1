package main

import (
	"sync"
	"testing"
)

func TestSquare(t *testing.T) {
	//Arrange
	inCh := make(chan int)
	outCh := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	//Act
	go square(inCh, outCh, &wg)

	// Отправляем значение
	inCh <- 4
	close(inCh)

	if actual := <-outCh; actual != 16 {
		t.Errorf("got %d, want %d", actual, 16)
	}
	wg.Wait()
}
func BenchmarkSquare(b *testing.B) {
	// 1. Подготовка каналов
	inCh := make(chan int, b.N)
	outCh := make(chan int, b.N)
	wg := &sync.WaitGroup{}
	wg.Add(1)

	// 2. Заполняем входной канал
	go func() {
		for i := 0; i < b.N; i++ {
			inCh <- 5 // Тестируем на числе 5 (5² = 25)
		}
		close(inCh)
	}()

	// 3. Запускаем тестируемую функцию
	b.ResetTimer() // Сбрасываем таймер перед измерением
	square(inCh, outCh, wg)

	// 4. Читаем все результаты
	for i := 0; i < b.N; i++ {
		<-outCh
	}
}
