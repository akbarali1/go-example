package main

import (
	"fmt"
	"time"
)

func task(id int) {
	fmt.Printf("Task %d started\n", id)
	time.Sleep(2 * time.Second) // Vaqt oladi
	fmt.Printf("Task %d completed\n", id)
}

func main() {
	// Bir nechta goroutine yaratish
	for i := 1; i <= 3; i++ {
		go task(i) // Har bir vazifa alohida goroutine sifatida ishlaydi
	}

	// Boshqa vazifalarni kutish
	time.Sleep(3 * time.Second) // 3 sekund kutish, barcha goroutine tugaydi
	fmt.Println("All tasks completed!")
}
