package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)
	values := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			select {
			case values <- i:
				time.Sleep(1 * time.Millisecond)
			case <-done:
				return
			}
		}
	}()

	go func() {
		time.Sleep(1 * time.Second)
		close(done)
		close(values)
	}()
	for val := range values {
		fmt.Println(val)
	}
}
