package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup, ch chan string) {
	defer wg.Done() // Goroutine tugagach, wait groupga xabar yuborish
	msg := fmt.Sprintf("Task %d completed", id)
	ch <- msg // Kanalga xabar yuborish
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 3) // Buffered channel (3 joyli)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go task(i, &wg, ch)
	}

	wg.Wait() // Barcha goroutinelarning tugashini kutish
	close(ch) // Kanalni yopish

	// Kanaldan ma'lumotlarni olish va chiqarish
	for msg := range ch {
		fmt.Println(msg)
	}
}
