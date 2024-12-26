package main

import (
	"fmt"
	"time"
)

// Функция count, которая работает с каналом
func count(ch <-chan int) {
	fmt.Println(ch)
	for num := range ch {
		result := num + num
		fmt.Printf("Удвоение числа %d равен %d\n", num, result)
	}
	fmt.Println(ch)
}

// В функции main создаём канал и заставляем его работать
func main() {
	ch := make(chan int)
	ch_2 := make(chan int)
	go count(ch_2)
	go count(ch)
	for i := 1; i <= 20; i++ {
		ch <- i
		ch_2 <- i
	}

	close(ch)
	time.Sleep(1 * time.Second)
}
