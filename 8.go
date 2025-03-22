package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			ch <- 3
			defer wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for val := range ch {
		fmt.Println(val)
	}
}
