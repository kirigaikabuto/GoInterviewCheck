package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 3
		ch <- 4
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
