package main

import "log"

/*
14. Разработать программу, которая в рантайме способна определить тип
переменной: int, string, bool, channel из переменной типа interface{}.
*/
func main() {
	cnt := 10
	stroka := "Yo"
	ch := make(chan int, cnt)
	multie(cnt)
	multie(ch)
	multie(stroka)
}

func multie(inter interface{}) {
	log.Printf("Type: %T", inter)
}
