package main

import (
	"fmt"
	"sync"
)

func IncreaseValue(wg *sync.WaitGroup, values chan int) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		values <- 1
	}
}

func main() {
	newCounter := 0
	wg := sync.WaitGroup{}
	ch := make(chan int, 100)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go IncreaseValue(&wg, ch)
	}
	go func() {
		for val := range ch {
			newCounter += val
		}
	}()
	wg.Wait()
	close(ch)
	fmt.Println(newCounter)
}
