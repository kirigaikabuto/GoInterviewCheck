package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 5
	}()
	fmt.Println(<-ch)
}
