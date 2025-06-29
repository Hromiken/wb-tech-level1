package main

import (
	"fmt"
	"math/big"
)

/*22. Разработать программу, которая перемножает, делит, складывает, вычитает две
числовых переменных a,b, значение которых > 2^20.
*/

func main() {
	// Через big.Int

	a := new(big.Int)
	a.SetString("3486784401", 10)
	b := new(big.Int)
	b.SetString("2621123515", 10)
	Calculate(a, b)

	// Через int64
	var c, d int64 = 3486784401, 2621123515

	Calc(c, d)

}

func Calculate(a, b *big.Int) {
	res := new(big.Int)
	fmt.Println("Способ big.Int")
	fmt.Printf("a * b = %s\n", res.Mul(a, b))
	fmt.Printf("a / b = %s\n", res.Div(a, b))
	fmt.Printf("a - b = %s\n", res.Sub(a, b))
	fmt.Printf("a + b = %s\n", res.Add(a, b))
}

func Calc(a, b int64) {
	fmt.Println("Способ Int64")
	fmt.Printf("a * b = %d\n", a*b)
	fmt.Printf("a / b = %d\n", a/b)
	fmt.Printf("a - b = %d\n", a-b)
	fmt.Printf("a + b = %d\n", a+b)
}
