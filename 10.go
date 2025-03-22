package main

import (
	"fmt"
	"time"
)

func worker(heartBreak chan bool) {
	time.Sleep(500 * time.Millisecond)
	select {
	case heartBreak <- true:
	default:

	}
}

func main() {
	heartBreak := make(chan bool)
	go worker(heartBreak)
	for {
		select {
		case <-heartBreak:
			fmt.Println("Active")
		case <-time.After(1 * time.Second):
			fmt.Println("not work")
			return
		}
	}
}
