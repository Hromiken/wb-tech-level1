package main

import "fmt"

/*
Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры
Human (аналог наследования).
*/
type Action struct {
	Human
}

type Human struct {
	Name string
	Age  int
}

func main() {
	example := Human{"Kate", 12}
	action := Action{
		Human: example,
	}
	action.Walk()
	action.Introduce()

}

func (a Action) Walk() {
	fmt.Printf("%s is walking on\n", a.Name)
}

func (h Human) Introduce() {
	fmt.Printf("I'm %d!\n", h.Age)
}
