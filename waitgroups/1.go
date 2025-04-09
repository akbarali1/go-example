package main

import (
	"fmt"
	"sync"
)

func tasks(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Vazifa tugagach, WaitGroupni xabardor qiladi
	fmt.Printf("Task %d started\n", id)
	fmt.Printf("Task %d completed\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Bir nechta goroutine yaratish
	for i := 1; i <= 10; i++ {
		wg.Add(1) // WaitGroupga yangi vazifa qo'shish
		go tasks(i, &wg)
	}

	// Barcha goroutinelar tugashini kutish
	wg.Wait()
	fmt.Println("All tasks completed!")
}
