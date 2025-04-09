package main

import (
	"fmt"
)

func greet(name string, ch chan string) {
	ch <- "Hello, " + name // Channelga xabar yuborish
}

func main() {
	ch := make(chan string) // Channel yaratish

	// Goroutine yaratish
	go greet("Alice", ch)

	// Channeldan ma'lumot olish
	msg := <-ch // Channeldan olingan qiymatni olish
	fmt.Println(msg)
}
