package main

import (
	"math/big"
	"testing"
)

var (
	bigA       = big.NewInt(3486784401)
	bigB       = big.NewInt(2621123515)
	intA int64 = 3486784401
	intB int64 = 2621123515
)

func BenchmarkCalc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = intA * intB
		_ = intA / intB
		_ = intA - intB
		_ = intA + intB
	}
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := new(big.Int)
		res.Add(bigA, bigB)
		res.Div(bigA, bigB)
		res.Sub(bigA, bigB)
		res.Mul(bigA, bigB)
	}
}

/*Результаты тестирования:
(int64)
Скорость: 0.1744 ns/op (наносекунд на операцию)
Аллокации: 0 B/op, 0 allocs/op (нет выделения памяти)

(big.Int)
Скорость: 46.82 ns/op (медленнее в ~270 раз)
Аллокации: 56 B/op, 2 allocs/op (дополнительные аллокации)

Выводы:
int64 в разы быстрее, но ограничен по размеру числа (~19 цифр).

big.Int медленнее, но поддерживает очень большие числа (тысячи цифр).

Если ваши числа гарантированно помещаются в int64, лучше использовать его.

Если нужна арифметика с огромными числами, то big.Int — единственный выбор.*/
