package main

import (
	"sync"
	"testing"
)

func TestSumSquares(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		want  int
	}{
		{"Два числа", []int{2, 4}, 20},
		{"Три числа", []int{1, 3, 5}, 35},
		{"Одно число", []int{10}, 100},
		{"Пустой ввод", []int{}, 0},
		{"Отрицательные числа", []int{-2, -3}, 13},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// ВВОД ДАННЫХ
			var index int
			chIn := make(chan int)
			chOut := make(chan int)
			var wg sync.WaitGroup
			wg.Add(1)
			// ACTING
			go sumSquare(chIn, chOut, &wg, index)
			// Мы проходимся по слайсам в структуре test
			go func() {
				for _, n := range tt.input {
					chIn <- n * n
				}
				close(chIn)
			}()
			// Проверка результатов
			if actual := <-chOut; actual != tt.want {
				t.Errorf("sumSquares() got = %v, want %v", actual, tt.want)
			}
			wg.Wait()
			close(chOut)
		})
	}
}
