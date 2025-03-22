package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Сообщение из ch1"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Сообщение из ch2"
	}()

	fmt.Println(<-ch1) // ❌ Блокируется, даже если ch2 уже готов!
	fmt.Println(<-ch2)
}
