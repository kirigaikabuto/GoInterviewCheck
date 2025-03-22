package main

import "fmt"

func main() {
	tempCh := make(chan int)
	go func() { tempCh <- 3 }()
	fmt.Println(<-tempCh)
}
