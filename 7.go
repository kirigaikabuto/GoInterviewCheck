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
		ch1 <- "Message from ch1"
	}()
	go func() {
		time.Sleep(1 * time.Second)
		ch2 <- "Message from ch2"
	}()
	for i := 0; i < 2; i++ {
		select {
		case val := <-ch1:
			fmt.Println(val)
		case val := <-ch2:
			fmt.Println(val)
		}
	}
}
