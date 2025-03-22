package main

import (
	"fmt"
	"sync"
)

func NewValueIncrease(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		ch <- 1
	}
}

func main() {
	counter1 := 0
	wg := sync.WaitGroup{}
	ch := make(chan int)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go NewValueIncrease(ch, &wg)
	}
	done := make(chan bool)
	go func() {
		for val := range ch {
			counter1 += val
		}
		done <- true
	}()
	wg.Wait()
	close(ch)
	<-done
	fmt.Println(counter1)
}
