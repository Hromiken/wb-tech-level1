package main

import (
	"context"
	"time"
)

/*Реализовать собственную функцию sleep.*/
func main() {
	mySleepAfter(10 * time.Second)
	mySleepTimer(10 * time.Second)
	mySleepContext(10 * time.Second)
}

func mySleepAfter(duration time.Duration) {
	<-time.After(duration) // Ждём, пока придёт значение в канал
}

/*
Плюсы:
Интеграция с select и другими каналами
Можно отменить через context

Минусы:
Создаёт новый канал и таймер при каждом вызове

Когда использовать:
В конкурентных сценариях с возможностью отмены*/

func mySleepTimer(duration time.Duration) {
	timer := time.NewTimer(duration)
	<-timer.C
	timer.Stop() // Важно для сборки мусора
}

/*Плюсы:
Лучше для частых вызовов (можно переиспользовать таймер)
Точнее управление ресурсами

Минусы:
Требует ручного вызова Stop()

Когда использовать:
В горячих циклах (high-load)*/

func mySleepContext(duration time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()
	<-ctx.Done()
}

/*Плюсы:
Интеграция с системой отмены через context
Позволяет обработать прерывание

Минусы:
Сложнее, чем базовые варианты

Когда использовать:
В сетевых запросах или долгих операциях*/
