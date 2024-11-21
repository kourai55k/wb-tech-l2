package main

import (
	"fmt"
	"time"
)

// Функция orFunc объединяет несколько done-каналов в один.
// Возвращает канал orChan, который закроется, как только закроется любой из переданных каналов.
func orFunc(channels ...<-chan interface{}) <-chan interface{} {
	result := make(chan interface{}) // Создаем выходной канал

	// Для каждого входного канала запускаем горутину
	for i := range channels {
		go func(i int) { // Индекс канала передается в горутину
			result <- (<-channels[i]) // Читаем из канала, как только он закроется
		}(i)
	}

	return result // Возвращаем канал, который закроется первым
}

func main() {
	// Объявляем переменную or и связываем ее с функцией orFunc
	var or func(channels ...<-chan interface{}) <-chan interface{} = orFunc

	// Функция sig создает канал, который закрывается через указанное время
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)    // Закрываем канал после задержки
			time.Sleep(after) // Задержка в работе канала
		}()
		return c
	}

	start := time.Now() // Засекаем время начала работы

	// Используем функцию or для ожидания завершения любого из каналов
	<-or(
		sig(2*time.Hour),
		sig(1*time.Second),
		sig(5*time.Minute),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	// Выводим, сколько времени прошло до закрытия первого канала
	fmt.Printf("done after %v\n", time.Since(start))
}
