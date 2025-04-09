package main

import (
	"fmt"
	"time"
)

func hello(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	hello("direct")
	go hello("goroutine")
	//go func(msg string) {
	fmt.Println("going")
	//fmt.Println(msg)
	//}("going")
	time.Sleep(time.Second)
	fmt.Println("done")
}
