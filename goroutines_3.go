package main

import (
	"fmt"
	"time"
)

// Oddiy funksiya
func hi(params string) {
	fmt.Println("Hello from Goroutine!:", params)
}

func main() {
	// Goroutine yaratish
	hi("1")    // Bu yerda sayHello() funksiyasi goroutine sifatida ishlaydi.
	go hi("2") // Bu yerda sayHello() funksiyasi goroutine sifatida ishlaydi.
	hi("3")    // Bu yerda sayHello() funksiyasi goroutine sifatida ishlaydi.

	// Boshqa kodni kutamiz, aks holda dastur darhol tugaydi
	time.Sleep(1 * time.Second) // 1 sekund kutish, shunda goroutine ishlashga vaqt oladi
	fmt.Println("Hello from main function!")
}
