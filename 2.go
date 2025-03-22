package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func increaseCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go increaseCounter(&wg)
	}
	wg.Wait()
	fmt.Println(counter)
}
