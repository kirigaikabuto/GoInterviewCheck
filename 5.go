package main

import "fmt"

func main() {
	ch := make(chan int)
	close(ch)

	fmt.Println(<-ch) // 0, так как канал закрыт
	fmt.Println(<-ch) // 0, но программа продолжает работать

	fmt.Println(<-ch) // Канал пуст и закрыт, но это не дедлок!
}
